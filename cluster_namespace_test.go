package main

import (
	"goclient/model"
	"goclient/watcher"
	"testing"
)

func TestMain2(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		depolyment := model.Deployment{}
		depolyment.RemoteFile = "ftp://10.10.10.34/test/FlachClick1.0.0.1.zip"
		watcher.StartDeploy(depolyment)
	})

}
