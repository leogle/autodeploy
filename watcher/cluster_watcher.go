package watcher

import (
	"encoding/json"
	"goclient/config"
	"goclient/etcd"
	"goclient/model"
	"goclient/monitor"
	"log"
)

func WatchCluster() {
	key := config.GlobalConfig.GetClusterKey()
	client := new(etcd.EtcdClient)
	client.InitEtcdWatcher(key, func(typ int32, key string, value []byte) {
		if typ == 0 {
			cluster := model.Cluster{}
			err := json.Unmarshal(value, &cluster)
			if err == nil {
				//updateCluster(cluster)
				return
			}
			log.Println("cluster unmarshal fail")
		}
	})
}

func ReadCluster() {
	key := config.GlobalConfig.GetClusterKey()
	client := new(etcd.EtcdClient)
	value := client.Get(key)[key]
	cluster := model.Cluster{}
	err := json.Unmarshal([]byte(value), &cluster)
	if err == nil {
		updateCluster(cluster)
		return
	}
	log.Println("cluster unmarshal fail")
}

func updateCluster(cluster model.Cluster) {
	found := false
	for _, node := range cluster.Nodes {
		if node.Name == config.GlobalConfig.Name {
			node = monitor.GetSysInfo(node)
			found = true
			break
		}
	}
	if !found {
		node := model.Node{Name: config.GlobalConfig.Name}
		node = monitor.GetSysInfo(node)
		cluster.Nodes = append(cluster.Nodes, node)
	}
	client := new(etcd.EtcdClient)
	text, _ := json.Marshal(cluster)
	client.Put(config.GlobalConfig.GetClusterKey(), string(text))
}
