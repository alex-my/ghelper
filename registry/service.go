package registry

import "encoding/json"

// Service 表示一种服务
type Service struct {
	// Name 服务的名称
	Name string `json:"name"`

	// Nodes 该服务下的节点信息
	Nodes []*Node `json:"nodes"`
}

// String ..
func (service *Service) String() string {
	bytes, _ := json.Marshal(service)
	return string(bytes)
}

// Node 表示提供服务的一个节点信息
type Node struct {
	// ID 用于辨认同一个服务下的不同节点，需要唯一性
	ID      string `json:"id"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// String ..
func (node *Node) String() string {
	bytes, _ := json.Marshal(node)
	return string(bytes)
}
