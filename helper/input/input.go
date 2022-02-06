package input

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"path/filepath"
	"strings"

	_ "image/png"

	"golang-base-code/helper/file"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/spf13/cast"
)

func File(c *gin.Context, inputName, fileStoragePath, fileName string) error {
	gologger.Info("Start processing upload file")
	f, err := c.FormFile(inputName)
	if err != nil {
		gologger.Error(err.Error())
		return err
	}

	if !file.IsPathExist(fileStoragePath) {
		gologger.Info("Creating a new nested file path")
		file.CreateNewNestedDirectory(fileStoragePath)
	}

	fullPathWithFileName := filepath.Join(fileStoragePath, fmt.Sprintf("%s%s", fileName, filepath.Ext(f.Filename)))

	fi, _ := f.Open()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, fi)
	isImage, _ := isImage(buf.Bytes())
	gologger.Info(fmt.Sprintf("Is image: %t", isImage))
	if isImage {
		gologger.Info("Start resize image")
		image, _, err := image.Decode(buf)
		if err != nil {
			gologger.Error(err.Error())
			gologger.Info("Finish resize image")
			return err
		}

		// resize image
		newImage := imaging.Fit(image, 200, 200, imaging.Lanczos)
		// save image
		imaging.Save(newImage, fullPathWithFileName)
		gologger.Info("Finish resize image")
		gologger.Info("Finish processing upload file")
		return nil
	}

	err = c.SaveUploadedFile(f, fullPathWithFileName)
	if err != nil {
		gologger.Error(err.Error())
		gologger.Info("Finish processing upload file")
		return err
	}
	gologger.Info("Finish processing upload file")
	return nil
}

func isImage(incipit []byte) (bool, string) {
	// image formats and magic numbers
	var magicTable = map[string]string{
		"\xff\xd8\xff":      "image/jpeg",
		"\x89PNG\r\n\x1a\n": "image/png",
		"GIF87a":            "image/gif",
		"GIF89a":            "image/gif",
	}

	incipitStr := []byte(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(cast.ToString(incipitStr), magic) {
			return true, mime
		}
	}

	return false, ""
}
