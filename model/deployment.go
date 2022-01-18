package model

import "time"

type Deployment struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`   //类型:IIS、CMD、Nginx
	Status string `json:"status"` //状态：Created、Published、Deployed、Stoped、Updated、Error
	//远程文件
	RemoteFile string `json:"remoteFile"`
	//
	SetupScript string `json:"setupScript"`
	//部署文件
	DeploymentScript string `json:"deploymentScript"`
	//
	TestScript string `json:"testScript"`
	//
	Nodes []string `json:"nodes"`
}

const DEPLOYMENT_STATUS_CREATE = "Created"
const DEPLOYMENT_STATUS_UPDATE = "Update"
const DEPLOYMENT_STATUS_DEPLOYED = "Deployed"

type DeploymentLog struct {
	Timepoint    time.Time
	Message      string
	DeploymentId int32
}
