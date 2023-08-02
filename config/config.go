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
	Port       string
	JwtSigKey  string
	JwtExpTime int64 // jwt token过期时间，单位：分钟
	Username   string
	Password   string
)

// 数据格式规范
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData

}

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
	//获取JWT加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "song")
	//获取JWT过期时间的配置
	viper.SetDefault("JWT_EXPIRE_TIME", 120)

	//获取用户名密码配置
	//MD5加密
	//默认值：admin song ，部署时请修改账户密码
	viper.SetDefault("USER_NAME_NAME", "21232F297A57A5A743894A0E4A801FC3")
	viper.SetDefault("PASS_WORD", "683EB609607A439B0561DCBB4C8329E8")

	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	//加载日志格式
	initLogConfig(logLevel)
	Port = viper.GetString("Post")
	JwtSigKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	//获取用户名密码
	Username = viper.GetString("USER_NAME_NAME")
	Password = viper.GetString("PASS_WORD")
	logs.Info(map[string]interface{}{"用户名": Username, "密码": Password}, "用户名密码默认变量值打印信息")
	initLogConfig(logLevel)

}
