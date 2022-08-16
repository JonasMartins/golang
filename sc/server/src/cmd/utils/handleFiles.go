package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	// rootPath = ?

	if err := os.Mkdir("public/images/"+uniqueFolderName, os.ModePerm); err != nil {
		log.Fatal(err)
		return "", err
	}

	tempFile, err := ioutil.TempFile("public/images/"+uniqueFolderName, "upload-*.png")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	tempFile.Write(content)
	return "public/images/" + uniqueFolderName, nil
}

// FUTURE, SEND THE FILES TO AWS

func GetRandonString() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}
