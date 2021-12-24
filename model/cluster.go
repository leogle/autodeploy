package model

// Cluster 集群信息
type Cluster struct {
	//集群节点
	Nodes []Node `json:"nodes"`
	//集群命名空间
	Namespaces []Namespace `json:"namespaces"`
}
