// 配置层：存放程序的配置信息
package config

import (
	"scaffold-demo/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port string
)

func initLogConfig(logLevel string) {
	//配置程序的日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//配置日志格式为JSON格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
	})
	//文件名字和行号加进去
	logrus.SetReportCaller(true)
}

func init() {
	logs.Debug(nil, "加载程序配置日志")
	//在部署环境中配置环境变量设置日志级别export LOG_LEVEL=DEBUG;
	viper.SetDefault("LOG_LEVEL", "Debug")
	//获取程序启动端口号的配置
	viper.SetDefault("Post", ":8080")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("Post")
	initLogConfig(logLevel)

}