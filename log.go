package chili

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	// 定义日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.ForceFormatting = true
	// 定义日志时间格式
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000"
	// 控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false

	logrus.SetFormatter(formatter)

	// 日志级别
	//level := os.Getenv("log.debug")
	//if level == "true" {
	logrus.SetLevel(logrus.DebugLevel)
	//}

	logrus.Info("foo")
	logrus.Debug("bar")
	logrus.Error("error")
}
