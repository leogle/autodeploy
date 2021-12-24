package main

import (
	"fmt"
	"go.etcd.io/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"go.etcd.io/etcd/clientv3"
	"goclient/config"
	"log"
	"time"
)

func main1() {
	config := config.ReadConfig()
	fmt.Println(config.Name)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2479", "localhost:2579"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Println("connect to etcd failed, err:", err)
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	putResp, err := cli.Put(ctx, "fruit", "orange")

	cancel()

	if err != nil {
		log.Println("put to ectcd failed:", err)
		return
	}
	log.Println("putResp:", putResp)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	getResp, err := cli.Get(ctx, "fruit")
	cancel()
	if err != nil {
		log.Printf("get from etcd failed, err:%v\n", err)
		return
	}

	log.Println("getResp:", getResp)

	for _, ev := range getResp.Kvs {
		log.Printf("%s:%s\n", ev.Key, ev.Value)
	}

}
