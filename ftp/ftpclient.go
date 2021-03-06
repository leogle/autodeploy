package ftp

import (
	"github.com/jlaffaye/ftp"
	"goclient/config"
	"io"
	"log"
	"os"
	"strings"
)

func Download(filename string, localPath string) {
	log.Println("下载FTP文件:" + filename)
	var ip string
	if strings.Index(filename, "ftp://") == 0 {
		ip = filename[6 : strings.Index(filename[6:], "/")+6]
	}
	if ip == "" {
		ip = config.GlobalConfig.Ftp.Host
	}
	remotefile := filename[strings.Index(filename, ip)+len(ip):]
	c, err := ftp.Dial(ip)
	if err == nil {
		err = c.Login(config.GlobalConfig.Ftp.User, config.GlobalConfig.Ftp.Password)
		if err == nil {
			r, err := c.Retr(remotefile)
			if err != nil {
				panic(err)
			}
			defer r.Close()
			file, _ := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE, 0777)
			io.Copy(file, r)
		}
	}
	log.Println("下载FTP文件成功:" + filename)
}

func GetFtpPath() {

}
