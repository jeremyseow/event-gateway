package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/jeremyseow/event-gateway/internal/utils"
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

func NewWriter(allUtilityClients *utils.AllUtilityClients, DirName func() string, FileName func() string) *FileWriter {
	// wrap in a recover to catch panic from funcs
	return &FileWriter{
		Logger: allUtilityClients.Logger,
		DirName: func() (dirName string) {
			defer func() {
				if err := recover(); err != nil {
					allUtilityClients.Logger.Error("error getting directory name, using default func GetDirName()", zap.String("logTag", logTag), zap.String("function", "DirName"))
					dirName = GetDirName()
				}
			}()
			return DirName()
		}, FileName: func() (fileName string) {
			defer func() {
				if err := recover(); err != nil {
					allUtilityClients.Logger.Error("error getting file name, using default func GetFileName()", zap.String("logTag", logTag), zap.String("function", "FileName"))
					fileName = GetFileName()
				}
			}()
			return FileName()
		},
	}
}

func (fw *FileWriter) Write(data []byte) (n int, err error) {
	currPath, err := os.Getwd()
	if err != nil {
		fw.Logger.Error("error when getting path", zap.String("logTag", logTag), zap.String("function", "main"))
		return 0, err
	}

	filePathName := path.Join(currPath, fw.DirName(), fw.FileName())
	dir := path.Dir(filePathName)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fw.Logger.Error("error when making directory", zap.String("logTag", logTag), zap.String("function", "Write"))
		return 0, err
	}

	if err := ioutil.WriteFile(filePathName, data, 0644); err != nil {
		fw.Logger.Error("error when writing file", zap.String("logTag", logTag), zap.String("function", "Write"))
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
