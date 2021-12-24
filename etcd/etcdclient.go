package etcd

import (
	"fmt"
	"go.etcd.io/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"go.etcd.io/etcd/clientv3"
	"goclient/config"
	"log"
	"time"
)

type EtcdInterface interface {
	Put(key string, value string)
	Get(key string) string
	Watch(key string)
}

type EtcdClient struct {
	client *clientv3.Client
}

func (client EtcdClient) Put(key string, value string) {
	cli := getClient()
	if cli != nil {
		defer cli.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		putResp, err := cli.Put(ctx, key, value)
		cancel()
		if err != nil {
			log.Println("put to ectcd failed:", err)
			return
		}
		log.Println("putResp:", putResp)
	}
}

func (client EtcdClient) Get(key string) map[string]string {
	cli := getClient()
	if cli != nil {
		defer cli.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		getRes, err := cli.Get(ctx, key)
		cancel()
		if err == nil {
			log.Println("put to ectcd failed:", err)
			dict := make(map[string]string)
			for _, kv := range getRes.Kvs {
				dict[string(kv.Key)] = string(kv.Value)
			}
			return dict
		}
	}
	return nil
}

func (client EtcdClient) Watch(key string) clientv3.WatchChan {
	cli := getClient()
	if cli != nil {
		defer cli.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cancel()
		return cli.Watch(ctx, key)
	}
	return nil
}

type WatchFunc func(typ int32, key string, value []byte)

func (client EtcdClient) InitEtcdWatcher(key string, callbackFunc WatchFunc) {
	/*	for _, key := range config.GlobalConfig.Etcd.WatchList {
		go watchKey(key)
	}*/
	go watchKey(key, callbackFunc)
}

func watchKey(key string, callbackFunc WatchFunc) {
	// 初始化连接etcd
	cli := getClient()
	//logs.Debug("开始监控key:", key)
	// Watch操作
	wch := cli.Watch(context.Background(), key)
	for resp := range wch {
		for _, ev := range resp.Events {
			fmt.Printf("Type: %v, Key:%v, Value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
			callbackFunc(int32(ev.Type), string(ev.Kv.Key), ev.Kv.Value)
		}
	}
}

func getClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.GlobalConfig.Etcd.Host,
		DialTimeout: 5 * time.Second,
	})

	if err == nil {
		return cli
	}
	return nil
}
