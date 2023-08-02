// 项目总入口
package main

import (
	"scaffold-demo/config"
	"scaffold-demo/middlewares"

	"scaffold-demo/utils/logs"

	"scaffold-demo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.加载程序的配置
	//2.配置gin

	logs.Error(nil, "打印info级别日志")
	r := gin.Default()
	r.Use(middlewares.JwtAuth)
	// //测试生成jwt token是否可用
	// ss, _ := jwtutils.GenToken("song")
	// fmt.Println("是否能生成TOKEN", ss)
	// //验证解析Token的方法
	// claims, err := jwtutils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNvbmciLCJpc3MiOiJkb3RiYWxvIiwic3ViIjoic29uZyIsImV4cCI6MTY5MDM2MDQ2NCwibmJmIjoxNjkwMzUzMjY0LCJpYXQiOjE2OTAzNTMyNjR9.DJgMu-zkYc-KQr_pjNKKlkJKgCJtUH4Iw3IomUfbgGY")
	// if err != nil {
	// 	//解析失败
	// 	fmt.Println("解析失败：",err.Error())
	// }else {
	// 	fmt.Println("解析成功：。。。。。。",claims)
	// }
	routers.RegisterRouters(r)
	r.Run(config.Port)
	

}
