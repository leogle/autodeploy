package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Name       string       `yaml:"name"`
	ApiVersion string       `yaml:"apiVersion"`
	Etcd       EtcdConfig   `yaml:"etcd"`
	Server     ServerConfig `yaml:"server"`
	FileDir    string       `yaml:"fileDir"`
	Ftp        FtpConfig    `yaml:"ftp"`
}

// GetClusterKey 获取集群Key
func (conf *Config) GetAppVersion() string {
	return "1.0.0"
}

// GetClusterKey 获取集群Key
func (conf *Config) GetClusterKey() string {
	return conf.ApiVersion + ".cluster"
}

// GetClusterNamespaceKey 获取集群命名空间Key
func (conf *Config) GetClusterNamespaceKey() string {
	return conf.ApiVersion + ".cluster.namespaces"
}

type EtcdConfig struct {
	Host      []string `yaml:"host"`
	WatchList []string `yaml:"watchList"`
}

type ServerConfig struct {
	IP string `yaml:"ip"`
}

type FtpConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var GlobalConfig Config

func ReadConfig() Config {
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.app: %#v\n", _config.Name)
	GlobalConfig = _config
	return _config
}
