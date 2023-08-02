// 路由层：管理程序的路由信息
package routers

import (
	"github.com/gin-gonic/gin"
	"scaffold-demo/routers/auth"
)

//注册路由的方法

func RegisterRouters(r *gin.Engine) {
	//登录的路由的配置
	//1.登录：login
	//2.退出：logout
	//3. /api/auth/login
	//   /api/auth/logout
	// /api/auth/
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}