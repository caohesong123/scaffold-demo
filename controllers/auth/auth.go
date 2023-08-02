package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutils"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录的逻辑
func Login(r *gin.Context) {
	//1.获取前端传递的用户名和密码
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		
		r.JSON(200, gin.H{
			"message": err.Error(),
			"status":  401,
		})
		return
	}
	logs.Info(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "开始验证登录信息")
	//验证用户名和密码是否正确
	//数据库 环境变量
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		//认证成功
		//生成Jwt的TOKEN
		//token失败
		ss, err := jwtutils.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err.Error()}, "用户名密码正确但是token失败")
			r.JSON(200, gin.H{
				"status":  401,
				"message": "生成token失败",
			})
			return
		}
		//token正常生成，返回给前端
		logs.Info(map[string]interface{}{"用户名": userInfo.Username}, "登录成功")
		// data := make(map[string]interface{})
		// data["token"] = ss
		returnData.Message ="登录成功"
		returnData.Data["token"] = ss
		r.JSON(200, returnData)
		return
	} else {
		//用户名密码错误
		r.JSON(200, gin.H{
			"status":  401,
			"message": "用户名密码错误",
		})
	}
}
func Logout(r *gin.Context) {
	//1.退出
	//实现逻辑
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Info(nil, "退出账户成功")

}
