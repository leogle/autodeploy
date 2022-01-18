package deploy

import "goclient/model"

type DeployWorker interface {
	Init(deployment model.Deployment)
	Deploy(deployment model.Deployment)
	Check(deployment model.Deployment)
}
