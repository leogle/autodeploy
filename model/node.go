package model

import (
	"time"
)

type Node struct {
	Name        string    `json:"name"`
	HostName    string    `json:"hostName"`
	Ip          string    `json:"ip"`
	Status      string    `json:"status"`
	OS          string    `json:"os"`
	CpuCore     int       `json:"cpuCore"`
	CpuUsage    string    `json:"cpu"`
	Memory      uint64    `json:"memory"`
	MemoryUsage uint64    `json:"memoryUsage"`
	Disk        string    `json:"disk"`
	AppVersion  string    `json:"appVersion"`
	UpdateTime  time.Time `json:"updateTime"`
}
