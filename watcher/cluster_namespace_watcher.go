package watcher

import (
	"encoding/json"
	"goclient/config"
	"goclient/etcd"
	"goclient/ftp"
	"goclient/model"
	"goclient/utils"
	"strings"
)

func WatchClusterNamespace() {
	key := config.GlobalConfig.GetClusterKey() + ".namespace"
	client := new(etcd.EtcdClient)
	client.InitEtcdWatcher(key, func(typ int32, key string, value []byte) {
		if typ == 0 {
			namespace := model.Namespace{}
			err := json.Unmarshal(value, &namespace)
			if err == nil {
				utils.ForEach(namespace.DeploymentSet, func(deployment model.Deployment) {
					if utils.Any(deployment.Nodes, func(node string) bool {
						return node == config.GlobalConfig.Name
					}) {
						if deployment.Status == model.DEPLOYMENT_STATUS_CREATE ||
							deployment.Status == model.DEPLOYMENT_STATUS_UPDATE {
							StartDeploy(deployment)
						}
					}
				})
			}
		} else if typ == 2 {

		}
	})
}

func StartDeploy(deployment model.Deployment) {
	fileName := deployment.RemoteFile[strings.LastIndex(deployment.RemoteFile, "/"):]
	localPath := config.GlobalConfig.FileDir + "/" + fileName
	ftp.Download(deployment.RemoteFile, localPath)
	utils.Unzip(localPath, config.GlobalConfig.FileDir)
}
