package auth

import (
	"github.com/gin-gonic/gin"
	"moriaty.com/cia/cia-common/base/constant"
	"moriaty.com/cia/cia-common/base/wrap"
	"net/http"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 20:34
 * @Description TODO
 * 	auth 权限校验中间件
 */

// 登录校验
func MiddleWare(auth Auth, name string, level ...uint8) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := auth.CheckToken(c.Request.Header.Get("authorization"))
		if token == nil {
			c.Abort()
			c.JSON(http.StatusOK, wrap.Error(constant.CODE_TOKEN_FAILURE, constant.MSG_TOKEN_FAILED))
			return
		}
		if len(level) != 0 && level[0] != token.Level {
			c.Abort()
			c.JSON(http.StatusOK, wrap.Error(constant.CODE_LEVEL_FAILURE, constant.MSG_LEVEL_FAILED))
			return
		}
		c.Keys = make(map[string]interface{})
		c.Keys[name] = token
		c.Next()
	}
}

// 获取 Token
func GetToken(c *gin.Context, name string) *Token {
	return c.Keys[name].(*Token)

}
