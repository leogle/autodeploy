package model

import "time"

type Deployment struct {
	Id     int32
	Type   string `json:"type"`   //类型:IIS、CMD、Nginx
	Status string `json:"status"` //状态：Created、Published、Deployed、Stoped、Updated、Error
	//远程文件
	RemoteFile string
	//
	SetupScript string
	//部署文件
	DeploymentScript string
	//
	TestScript string
	//
	Nodes []string
}

const DEPLOYMENT_STATUS_CREATE = "Created"
const DEPLOYMENT_STATUS_UPDATE = "Update"

type DeploymentLog struct {
	Timepoint    time.Time
	Message      string
	DeploymentId int32
}
