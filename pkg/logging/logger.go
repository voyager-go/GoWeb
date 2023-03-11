package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(logPath string) {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(io.MultiWriter(os.Stdout, createLogFile(logPath)))
	Log.SetLevel(logrus.DebugLevel)
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.PathMap{
			logrus.ErrorLevel: logPath,
			logrus.FatalLevel: logPath,
			logrus.PanicLevel: logPath,
		},
		&logrus.JSONFormatter{},
	))
}

func createLogFile(logPath string) *os.File {
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		os.MkdirAll(logPath, 0755)
	}
	dateStr := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s/%s.log", logPath, dateStr)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatalf("Failed to open log file: %v", err)
	}
	return file
}
