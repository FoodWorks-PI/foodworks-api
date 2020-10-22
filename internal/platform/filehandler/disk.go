package filehandler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type FileHandler interface {
	Upload(upload *graphql.Upload) (string, error)
	Delete() (string, error)
	Download(w http.ResponseWriter, r *http.Request, fileName string)
}

type DiskUploader struct{}

var mediaPath string

func NewDiskUploader() FileHandler {
	mediaPath = "images"
	strict := os.FileMode(0766)
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		err := os.MkdirAll(mediaPath, strict)
		if err != nil {
			log.Err(err)
			log.Fatal().Msgf("Unable to create dir: %s", mediaPath)
		}
	}
	return &DiskUploader{}
}

func (du *DiskUploader) Delete() (string, error) {
	return "", nil
}

func (du *DiskUploader) Download(w http.ResponseWriter, r *http.Request, fileName string) {
	filePath := path.Join(mediaPath, fileName)
	http.ServeFile(w, r, filePath)
}

func (du *DiskUploader) Upload(upload *graphql.Upload) (string, error) {
	content, err := ioutil.ReadAll(upload.File)
	if err != nil {
		return "", err
	}
	fileUUID := uuid.New().String()
	fileExtension := filepath.Ext(upload.Filename)
	fileName := fileUUID + fileExtension
	fileFullPath := filepath.Join(mediaPath, fileName)
	out, err := os.Create(fileFullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = out.Write(content)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
