package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func Unzip(file string, dest string) {
	r, err := zip.OpenReader(file)
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
