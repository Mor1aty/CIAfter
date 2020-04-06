package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moriaty.com/cia/cia-common/auth"
	"moriaty.com/cia/cia-common/base/wrap"
	"moriaty.com/cia/cia-common/valid"
	"net/http"
)

/**
	任务管理者 CIA-Manager
  	1. 提供生成机制
  	2. 生成任务
  	3. 将任务分发到任务容器，方便任务执行者执行任务
  	4. 提供任务结果获取机制
*/

type UserData struct {
	account, password string
}

func main() {

	// init
	r := gin.Default()
	memoryAuth := auth.NewAuth("memory")

	// 登录测试
	r.POST("/login", valid.MiddleWare([]valid.Param{
		{Name: "account", Methods: []valid.Method{valid.PHONE}},
		{Name: "password"},
	}), func(c *gin.Context) {
		params := valid.GetWrapParams(c)
		fmt.Println(params)
		token := auth.Token{EffectiveTime: -1, Data: &UserData{
			account:  params.GetString("account"),
			password: params.GetString("password"),
		}}
		tokenCode := memoryAuth.PutTokenStorage(token)
		c.JSON(http.StatusOK, wrap.OKExec("登录成功", tokenCode))
	})

	// 访问测试
	r.GET("/login/after", auth.MiddleWare(memoryAuth, "user"), func(c *gin.Context) {
		token := auth.GetToken(c, "user")
		c.JSON(http.StatusOK, wrap.OKExec("获取成功", token))
	})

	r.Run(":8080")
}
