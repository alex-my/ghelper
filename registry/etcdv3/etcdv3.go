// Package etcdv3 ...
package etcdv3

import (
	"context"
	"encoding/json"
	"errors"
	"path"
	"sync"

	etcdcli "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	"github.com/alex-my/ghelper/registry"

	hash "github.com/mitchellh/hashstructure"
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

	// leases 存储租约信息
	leases map[string]etcdcli.LeaseID

	// hashs 存储服务散列信息
	hashs map[string]uint64

	closeOnce sync.Once

	lock sync.Mutex
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
		leases: make(map[string]etcdcli.LeaseID),
		hashs:  make(map[string]uint64),
	}, nil
}

// Register 注册一个服务
func (etcdv3 *etcdv3) Register(service *registry.Service) error {
	// 服务必须带有节点信息
	if len(service.Nodes) == 0 {
		return errors.New("at least one node")
	}

	// 如果已有租约，则延长租约，不用新建
	var leaseKeep bool
	leaseID, ok := etcdv3.leases[service.Name]
	if ok {
		ctx, cancel := context.WithTimeout(context.Background(), etcdv3.config.WriteTimeout)
		defer cancel()

		if _, err := etcdv3.client.KeepAliveOnce(ctx, leaseID); err != nil {
			if err != rpctypes.ErrLeaseNotFound {
				return err
			}
		} else {
			leaseKeep = true
		}
	}

	var err error

	// 判断服务内容是否发生了变化
	h, err := hash.Hash(service, nil)
	if err != nil {
		return err
	}

	// 如果服务内容没有发生改变，且又成功延长了租约时间，则不需要重新 put 到 etcd 中
	etcdv3.lock.Lock()
	oh, ok := etcdv3.hashs[service.Name]
	etcdv3.lock.Unlock()

	if ok && oh == h && leaseKeep {
		return nil
	}

	// 向 etcd put 数据
	ctx, cancel := context.WithTimeout(context.Background(), etcdv3.config.WriteTimeout)
	defer cancel()

	var lease *etcdcli.LeaseGrantResponse
	if etcdv3.config.TTL > 0 {
		lease, err = etcdv3.client.Grant(ctx, int64(etcdv3.config.TTL.Seconds()))
		if err != nil {
			return err
		}
	}

	for _, node := range service.Nodes {
		key := nodePath(service.Name, node)
		value := pack(node)

		var err error
		if lease != nil {
			_, err = etcdv3.client.Put(ctx, key, value, etcdcli.WithLease(lease.ID))
		} else {
			_, err = etcdv3.client.Put(ctx, key, value)
		}

		if err != nil {
			return err
		}
	}

	// 记录一些数据
	etcdv3.lock.Lock()
	etcdv3.hashs[service.Name] = h
	if lease != nil {
		etcdv3.leases[service.Name] = lease.ID
	}

	etcdv3.lock.Unlock()

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
