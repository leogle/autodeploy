package model

type Namespace struct {
	Name          string       `json:"name"`
	DeploymentSet []Deployment `json:"deploymentSet"`
}
