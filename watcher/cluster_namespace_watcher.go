package watcher

import (
	"encoding/json"
	"goclient/config"
	"goclient/deploy"
	"goclient/etcd"
	"goclient/ftp"
	"goclient/model"
	"goclient/utils"
	"log"
	"strings"
)

func WatchClusterNamespace() {
	key := config.GlobalConfig.GetClusterKey() + ".namespaces"
	client := new(etcd.EtcdClient)
	client.InitEtcdWatcher(key, func(typ int32, key string, value []byte) {
		if typ == 0 {
			log.Println("Key变化" + key)
			namespaces := []model.Namespace{}
			err := json.Unmarshal(value, &namespaces)
			if err != nil {
				return
			}
			for _, namespace := range namespaces {
				for _, deployment := range namespace.DeploymentSet {
					for _, node := range deployment.Nodes {
						if node == config.GlobalConfig.Name {
							if deployment.Status == model.DEPLOYMENT_STATUS_CREATE ||
								deployment.Status == model.DEPLOYMENT_STATUS_UPDATE {
								StartDeploy(deployment)
								deployment.Status = model.DEPLOYMENT_STATUS_DEPLOYED
								buf, _ := json.Marshal(deployment)
								client.Put(key, string(buf))
							}
						}
					}
				}
			}
		} else if typ == 2 {

		}
	})
}

func StartDeploy(deployment model.Deployment) {
	log.Println("开始部署" + deployment.Type)
	fileName := deployment.RemoteFile[strings.LastIndex(deployment.RemoteFile, "/"):]
	localPath := config.GlobalConfig.FileDir + "/" + fileName
	ftp.Download(deployment.RemoteFile, localPath)
	utils.Unzip(localPath, config.GlobalConfig.FileDir)
	var worker deploy.DeployWorker
	if deployment.Type == "SHELL" {
		worker = deploy.ShellDeploy{}
	}
	worker.Init(deployment)
	worker.Deploy(deployment)
	worker.Check(deployment)
}
