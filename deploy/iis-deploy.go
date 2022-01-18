package deploy

import "goclient/model"

type IISDeploy struct {
	DeployWorker
}

func (worker IISDeploy) Init(deployment model.Deployment) {
}
