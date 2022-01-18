package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFileReadOnly(fullFilepath, content string) {
	// jika path belum ada
	if !IsPathExist(fullFilepath) {
		file := filepath.Base(fullFilepath)
		dirPath := strings.Replace(fullFilepath, file, "", 1)

		fmt.Println("creating a new nested file path")
		CreateNewNestedDirectory(dirPath)
	}

	// ubah permission file jadi writeable
	SetWritable(fullFilepath)

	f, err := os.OpenFile(fullFilepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("error opening file:", err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		fmt.Println("error writing file:", err)
	}

	// ubah permission file jadi read-only
	SetReadOnly(fullFilepath)
}

func SetWritable(filepath string) error {
	err := os.Chmod(filepath, 0222)
	if err != nil {
		fmt.Println("error change permission file to writeable:", err)
	}
	return err
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	if err != nil {
		fmt.Println("error change permission file to read-only:", err)
	}
	return err
}

func CreateNewNestedDirectory(folderPath string) error {
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println("error creating a new directory or file:", err)
	}
	return err
}

func IsPathExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
