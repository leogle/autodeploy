package monitor

import (
	"goclient/config"
	"goclient/model"
	"os"
	"runtime"
	"time"
)

func GetSysInfo(node model.Node) model.Node {
	node.AppVersion = config.GlobalConfig.GetAppVersion()
	node.OS = runtime.GOOS
	node.HostName, _ = os.Hostname()
	node.CpuCore = runtime.GOMAXPROCS(0)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	node.Memory = m.TotalAlloc / 1024
	node.MemoryUsage = m.TotalAlloc / 1024
	node.UpdateTime = time.Now()
	return node
}
