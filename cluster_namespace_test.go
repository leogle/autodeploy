package main

import (
	"goclient/model"
	"goclient/watcher"
	"testing"
)

func TestMain2(t *testing.T) {

	t.Run("test2", func(t *testing.T) {
		depolyment := model.Deployment{}
		depolyment.Type = "SHELL"
		depolyment.SetupScript = "echo hello"
		depolyment.DeploymentScript = "FlaskClick.exe"
		depolyment.TestScript = "echo end"
		depolyment.RemoteFile = "ftp://10.10.10.34/test/FlachClick1.0.0.1.zip"
		watcher.StartDeploy(depolyment)

	})

}
