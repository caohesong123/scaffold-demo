package auth

import (
	"scaffold-demo/controllers/auth"

	"github.com/gin-gonic/gin"
)

//实现登录接口
func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login",auth.Login)
}

//实现退出接口
func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout",auth.Logout)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	//配置登录功能的路由策略
	authGroup := g.Group("/auth")
	//登录的功能
	login(authGroup)
	logout(authGroup)
}
