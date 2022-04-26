package chili

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 定义日志格式
	formatter := &logrus.TextFormatter{}
	logrus.SetFormatter(formatter)

	// 日志级别
	level := os.Getenv("log.debug")
	if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// 控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false

	logrus.Info("info")
	logrus.Debug("debug")
}
