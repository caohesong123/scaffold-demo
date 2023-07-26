// 项目总入口
package main

import (
	"fmt"
	"scaffold-demo/config"

	"scaffold-demo/utils/jwtutils"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.加载程序的配置
	//2.配置gin

	logs.Error(nil, "打印info级别日志")
	r := gin.Default()
	//测试生成jwt token是否可用
	ss, _ := jwtutils.GenToken("song")
	fmt.Println("是否能生成TOKEN", ss)
	r.Run(config.Port)

}
