package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	logTag = "file.FileWriter"
)

type FileWriter struct {
	Logger   *zap.Logger
	DirName  func() string
	FileName func() string
}

func NewWriter(Logger *zap.Logger, DirName func() string, FileName func() string) *FileWriter {
	return &FileWriter{Logger: Logger, DirName: DirName, FileName: FileName}
}

func (fw *FileWriter) Write(data []byte) (n int, err error) {
	currPath, err := os.Getwd()
	if err != nil {
		fw.Logger.Error("error when getting path", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return 0, err
	}

	fw.Logger.Info(fw.DirName())
	fw.Logger.Info(fw.FileName())

	filePathName := path.Join(currPath, fw.DirName(), fw.FileName())
	dir := path.Dir(filePathName)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fw.Logger.Error("error when making directory", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return 0, err
	}

	if err := ioutil.WriteFile(filePathName, data, 0644); err != nil {
		fw.Logger.Error("error when writing file", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return 0, err
	}

	return len(data), nil
}

func GetDirName() string {
	currTime := time.Now().UTC()
	filepath := fmt.Sprintf("output/year=%d/month=%d/day=%d/hour=%d", currTime.Year(), currTime.Month(), currTime.Day(), currTime.Hour())
	return filepath
}

func GetFileName() string {
	uuid := uuid.New().String()
	filename := fmt.Sprintf("%s-%s.txt", "events", uuid)
	return filename
}
