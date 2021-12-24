package main

import (
	"github.com/kardianos/service"
	"goclient/config"
	"goclient/watcher"
	"log"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}
func (p *program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}
func (p *program) run() {
	//读取配置
	config.ReadConfig()
	watcher.StartWatch()
	/*client := new(etcd.EtcdClient)
	client.InitEtcdWatcher("test.key", func(typ int32, key string, value []byte) {
		if typ == 0 {
			log.Println(value)
		}
	})*/
}

func main() {
	cfg := &service.Config{
		Name:        "serverlet",
		DisplayName: "serverlet service",
		Description: "This is an example Go service.",
	}
	prg := &program{}
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
