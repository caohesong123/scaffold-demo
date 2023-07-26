// 项目总入口
package main

import (
	"scaffold-demo/config"

	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.加载程序的配置
	//2.配置gin

	logs.Error(nil, "打印info级别日志")
	r := gin.Default()
	r.Run(config.Port)

}
