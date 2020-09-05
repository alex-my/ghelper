package etcdv3_test

import (
	"testing"

	"github.com/alex-my/ghelper/registry"
	"github.com/alex-my/ghelper/registry/etcdv3"
)

func newServices(name string) *registry.Service {
	return &registry.Service{
		Name: name,
		Nodes: []*registry.Node{
			&registry.Node{
				ID:      "1",
				Address: "127.0.0.1",
				Port:    9001,
			},
			&registry.Node{
				ID:      "2",
				Address: "127.0.0.1",
				Port:    9002,
			},
			&registry.Node{
				ID:      "3",
				Address: "127.0.0.1",
				Port:    9003,
			},
		},
	}
}

func TestRegister(t *testing.T) {
	r, err := etcdv3.NewEtcdv3(
		registry.WithAddrs("127.0.0.1:7179", "127.0.0.1:7279", "127.0.0.1:7379"),
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	names := []string{"test-name-register-1", "test-name-register-2", "test-name-register-3"}

	for _, name := range names {
		services := newServices(name)
		if err := r.Register(services); err != nil {
			t.Fatal(err.Error())
		}

		service, err := r.Service(name)
		if err != nil {
			t.Fatal(err.Error())
		}

		t.Log(service.String())
	}
}

func TestDeregister(t *testing.T) {
	r, err := etcdv3.NewEtcdv3(
		registry.WithAddrs("127.0.0.1:7179", "127.0.0.1:7279", "127.0.0.1:7379"),
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := "test-name-deregister"

	// 注册
	services := newServices(name)
	if err := r.Register(services); err != nil {
		t.Fatal(err.Error())
	}

	// 获取数量
	service, err := r.Service(name)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 注销
	if err := r.Deregister(service); err != nil {
		t.Fatal(err.Error())
	}

	// 重新获取
	service, err = r.Service(name)
	if err != nil && err != registry.ErrNodeNotFound {
		t.Fatal(err.Error())
	}

	if service != nil && len(service.Nodes) > 0 {
		t.Fatal("deregister failed")
	}
}
