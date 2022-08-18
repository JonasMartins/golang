package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleUploads(content []byte, name string) (string, error) {

	upload := os.Getenv("UPLOAD_STORAGE")
	if upload == "" {
		log.Fatal("please add a valid env")
		return "", errors.New(`Need a valid env, "UPLOAD_STORAGE" variable required`)
	}

	// if upload == local, then store locally, else call the
	// method that uploads to aws

	path, err := manageStoreFileLocally(content, name)
	if err != nil {
		return "", err
	}

	return path, nil
}

func manageStoreFileLocally(content []byte, name string) (string, error) {
	uniqueFolderName := GetRandonString()
	rootPath, err := GetRootApplicationPath()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var path string = rootPath + "public/images/" + uniqueFolderName + "/"

	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal(err)
		return "", err
	}

	tempFile, err := os.Create(path + name)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	tempFile.Write(content)
	return "public/images/" + uniqueFolderName + "/" + name, nil
}

// FUTURE, SEND THE FILES TO AWS

func GetRandonString() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func GetRootApplicationPath() (string, error) {
	var path string = ""
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	arrDir := strings.Split(currDir, "/")
	var i int = 0
	for i = (len(arrDir) - 1); i >= 0; i-- {
		if arrDir[i] == "server" {
			break
		}
	}

	for j := 0; j <= i; j++ {
		path += arrDir[j] + "/"
	}

	return path, nil

}
