// Package registry 服务注册与发现
package registry

import "errors"

var (
	// ErrNodeNotFound 服务节点未找到
	ErrNodeNotFound = errors.New("node not found")
)

// Registry 服务器注册与发现接口
type Registry interface {
	// Register 注册一个服务
	Register(service *Service) error

	// Deregister 注销一个服务
	Deregister(service *Service) error

	// Service 获取指定名称的服务列表
	Service(name string) (*Service, error)

	// Close 关闭
	Close()
}
