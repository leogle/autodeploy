package main

import (
	"go.etcd.io/etcd/clientv3"
)

/*func main() {
	config.ReadConfig()
	client := new(etcd.EtcdClient)
	client.Put("test.key.subkey", "testval")

	kvs := client.Get("test.key")
	for k, v := range kvs {
		log.Printf("k %s v %s", k, v)
	}

	watchChan := client.Watch("test.key")
	//var val clientv3.WatchResponse
	go readChan(watchChan)

}*/

func readChan(watchChan clientv3.WatchChan) {
	/*	for val := range watchChan {
		for _, evt := range val.Events {
			log.Printf("evt:k: %s v: %s", evt.Kv.Key, evt.Kv.Value)
		}
	}*/
}
