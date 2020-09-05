// Package etcdv3 ...
package etcdv3

import (
	"context"
	"encoding/json"
	"errors"
	"path"
	"sync"
	"time"

	etcdcli "github.com/coreos/etcd/clientv3"

	"github.com/alex-my/ghelper/registry"
)

const (
	// prefix 避免多项目共用一个 etcd 的时候混淆
	prefix = "registry-Aa1I75"
)

type etcdv3 struct {
	// config 一些配置
	config *registry.Config

	// client 与 etcd 连接的客户端
	client *etcdcli.Client

	closeOnce sync.Once
}

// NewEtcdv3 ..
func NewEtcdv3(opts ...registry.Option) (registry.Registry, error) {
	config := registry.DefaultConfig()

	for _, opt := range opts {
		opt(config)
	}

	client, err := etcdcli.New(etcdcli.Config{
		Endpoints:   config.Addrs,
		DialTimeout: config.DialTimeout,
		TLS:         config.TLSConfig,
		Username:    config.Username,
		Password:    config.Password,
	})
	if err != nil {
		return nil, err
	}

	return &etcdv3{
		config: config,
		client: client,
	}, nil
}

// Register 注册一个服务
func (etcdv3 *etcdv3) Register(service *registry.Service) error {
	// 服务必须带有节点信息
	if len(service.Nodes) == 0 {
		return errors.New("at least one node")
	}

	ctx, cancel := context.WithTimeout(context.Background(), etcdv3.config.WriteTimeout)
	defer cancel()

	var err error
	if etcdv3.config.TTL > 0 {
		err = etcdv3.putWithTTL(ctx, service, etcdv3.config.TTL)
	} else {
		err = etcdv3.put(ctx, service)
	}

	if err != nil {
		return err
	}

	return nil
}

// Deregister 注销一个服务
func (etcdv3 *etcdv3) Deregister(service *registry.Service) error {
	// 服务必须带有节点信息
	if len(service.Nodes) == 0 {
		return errors.New("at least one node")
	}

	ctx, cancel := context.WithTimeout(context.Background(), etcdv3.config.WriteTimeout)
	defer cancel()

	for _, node := range service.Nodes {
		key := nodePath(service.Name, node)
		if _, err := etcdv3.client.Delete(ctx, key); err != nil {
			return err
		}
	}

	return nil
}

// Service 获取指定名称的服务信息
func (etcdv3 *etcdv3) Service(name string) (*registry.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), etcdv3.config.WriteTimeout)
	defer cancel()

	// 以服务名称为前缀，获取该前缀的所有信息
	key := servicePath(name)
	res, err := etcdv3.client.Get(ctx, key, etcdcli.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(res.Kvs) == 0 {
		return nil, registry.ErrNodeNotFound
	}

	nodes := make([]*registry.Node, 0, len(res.Kvs))

	for _, kv := range res.Kvs {
		node := unpack(kv.Value)
		if node != nil {
			nodes = append(nodes, node)
		}
	}

	return &registry.Service{
		Name:  name,
		Nodes: nodes,
	}, nil
}

// Close 关闭
func (etcdv3 *etcdv3) Close() {
	etcdv3.closeOnce.Do(func() {
		if etcdv3.client != nil {
			etcdv3.client.Close()
		}
	})
}

func nodePath(serviceName string, node *registry.Node) string {
	return path.Join(prefix, serviceName, node.ID)
}

func servicePath(serviceName string) string {
	return path.Join(prefix, serviceName)
}

func pack(node *registry.Node) string {
	bytes, _ := json.Marshal(node)
	return string(bytes)
}

func unpack(bytes []byte) *registry.Node {
	var node *registry.Node
	json.Unmarshal(bytes, &node)
	return node
}

func (etcdv3 *etcdv3) put(ctx context.Context, service *registry.Service) error {
	for _, node := range service.Nodes {
		key := nodePath(service.Name, node)
		value := pack(node)

		if _, err := etcdv3.client.Put(ctx, key, value); err != nil {
			return err
		}
	}

	return nil
}

func (etcdv3 *etcdv3) putWithTTL(ctx context.Context, service *registry.Service, ttl time.Duration) error {
	// 创建一个租约
	l, err := etcdv3.client.Grant(ctx, int64(ttl.Seconds()))
	if err != nil {
		return err
	}

	for _, node := range service.Nodes {
		key := nodePath(service.Name, node)
		value := pack(node)

		if _, err = etcdv3.client.Put(ctx, key, value, etcdcli.WithLease(l.ID)); err != nil {
			return err
		}
	}

	return nil
}
