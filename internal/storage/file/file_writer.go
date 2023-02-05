package file

import (
	"io/ioutil"
	"os"
	"path"

	"go.uber.org/zap"
)

const (
	logTag = "file.FileWriter"
)

type FileWriter struct {
	Logger *zap.Logger
}

func NewWriter(logger *zap.Logger) *FileWriter {
	return &FileWriter{Logger: logger}
}

func (fw *FileWriter) Write(dirName string, fileName string, data []byte) error {
	currPath, err := os.Getwd()
	if err != nil {
		fw.Logger.Error("error when getting path", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return err
	}

	filePathName := path.Join(currPath, dirName, fileName)
	dir := path.Dir(filePathName)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fw.Logger.Error("error when making directory", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return err
	}

	if err := ioutil.WriteFile(filePathName, data, 0644); err != nil {
		fw.Logger.Error("error when writing file", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return err
	}

	return nil
}
