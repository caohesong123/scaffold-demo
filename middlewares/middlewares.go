// 中间件层
package middlewares

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutils"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func JwtAuth(r *gin.Context) {
	//1.除了login和logout之外的所有的接口，都要验证是否请求携带token
	requestUrl := r.FullPath()
	logs.Info(map[string]interface{}{"请求路径": requestUrl},"")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Info(map[string]interface{}{"请求路径": requestUrl},"登录和退出不验证TOKEN")
		r.Next()
		return
	}
	returnData := config.NewReturnData()
	//TOKEN
	//其他接口验证token
	//获取是否写道TOKEN
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		//说明请求没有携带token
		returnData.Status = 401
		returnData.Message ="请求未携带token,请登录后尝试"
		r.JSON(200,returnData)
		r.Abort()
		return	
	}
	//token不为空，要去验证token是否合法
	claims, err := jwtutils.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "token验证未通过"
		r.JSON(200,returnData)
		r.Abort()
		return
	}
	//验证成功
	r.Set("claims",claims)
	r.Next()
}
