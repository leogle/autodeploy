package deploy

import (
	"bufio"
	"goclient/config"
	"goclient/model"
	"goclient/utils"
	"log"
	"os"
	"os/exec"
	"time"
)

type ShellDeploy struct {
	deployment model.Deployment
}

func (shell ShellDeploy) Init(deployment model.Deployment) {
	shell.deployment = deployment
	script := deployment.SetupScript
	log.Println("启动初始化:" + deployment.SetupScript)
	runScript(script, deployment)
}

func (shell ShellDeploy) Deploy(deployment model.Deployment) {
	script := deployment.DeploymentScript
	log.Println("启动部署:" + deployment.DeploymentScript)
	runScript(script, deployment)
}

func (shell ShellDeploy) Check(deployment model.Deployment) {
	script := deployment.TestScript
	log.Println("启动校验:" + deployment.TestScript)
	runScript(script, deployment)
}

func runScript(script string, deployment model.Deployment) {
	fileName := "script-" + deployment.Type + time.Now().Format("2006-01-02-15-04-05") + ".bat"
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	file.WriteString(script)
	defer file.Close()
	cmd := exec.Command("cmd", "/c", fileName)
	cmd.Dir = utils.PathCombine(config.GlobalConfig.WorkDir)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Panicf("Error creating StdoutPipe for Cmd:%s", err.Error())
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()
	err = cmd.Run()
	if err != nil {
		log.Printf("script error %s", err.Error())
	}
	err = os.RemoveAll(fileName)
	if err != nil {
		log.Printf("delete file error %s", err.Error())
	}
}

func daemon() {

}
