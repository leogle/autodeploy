package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func Unzip(file string, dest string) {
	r, err := zip.OpenReader(file)
	defer r.Close()
	if err != nil {
		log.Println(err)
		return
	}

	for _, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(dest+k.Name, 0644)
			if err != nil {
				log.Println(err)
			}
			continue
		}
		if strings.Index(k.FileHeader.Name, "/") > 0 {
			fileName := k.FileHeader.Name
			for strings.Index(fileName, "/") > 0 {
				index := strings.Index(fileName, "/")
				dir := fileName[:index]
				os.MkdirAll(PathCombine(dest, dir), 0644)
				fileName = fileName[index+1:]
			}
		}
		r, err := k.Open()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("unzip: ", k.Name)
		defer r.Close()
		NewFile, err := os.Create(dest + k.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
}
