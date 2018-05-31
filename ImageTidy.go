package main

import (
	"os"
	"log"
	"path/filepath"
	"flag"
	"strings"
	"io"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "images directory path")
	flag.Parse()
}

func main() {
	//validation
	if len(path) < 1 {
		log.Fatalln("error: path is empty.")
	} else {
		log.Println("path: " + path)
	}

	//tmp dir create
	root, _ := os.Getwd()
	tmpDir := root + "/tmp/"
	hdpiDir := tmpDir + "/drawable-hdpi/"
	xhdpiDir := tmpDir + "/drawable-xhdpi/"
	xxhdpiDir := tmpDir + "/drawable-xxhdpi/"

	if err := os.RemoveAll(tmpDir); err != nil {
		log.Println("error: tmpDir dir not deleted")
		log.Fatalln(err)
	}

	if err := os.MkdirAll(tmpDir, 0777); err != nil {
		log.Println("error: tmpDir dir not created")
		log.Fatalln(err)
	}

	if err := os.MkdirAll(hdpiDir, 0777); err != nil {
		log.Println("error: hdpiDir dir not created")
		log.Fatalln(err)
	}

	if err := os.MkdirAll(xhdpiDir, 0777); err != nil {
		log.Println("error: xhdpiDir dir not created")
		log.Fatalln(err)
	}

	if err := os.MkdirAll(xxhdpiDir, 0777); err != nil {
		log.Println("error: xxhdpiDir dir not created")
		log.Fatalln(err)
	}

	//file list load & copy
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			infoName := rename(info.Name())

			var dstPath string
			if strings.Contains(infoName, "@1x") {
				dstPath = hdpiDir + strings.Replace(infoName, "@1x", "", 1)
			}
			if strings.Contains(infoName, "@2x") {
				dstPath = xhdpiDir + strings.Replace(infoName, "@2x", "", 1)
			}
			if strings.Contains(infoName, "@3x") {
				dstPath = xxhdpiDir + strings.Replace(infoName, "@3x", "", 1)
			}
			if len(dstPath) < 1 {
				return nil
			}

			log.Println(path)
			src, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer src.Close()

			dst, err := os.Create(dstPath)
			if err != nil {
				panic(err)
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalln("Error on filepath.Walk : ", err)
	}
}

func rename(value string) string {
	var result = value
	result = strings.Replace(result, "-", "_", -1)
	result = strings.Replace(result, "ON", "_on", -1)
	result = strings.Replace(result, "NAVI", "_navi", -1)
	result = strings.ToLower(result)
	return result
}
