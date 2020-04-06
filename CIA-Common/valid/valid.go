package valid

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	"moriaty.com/cia/cia-common/base/wrap"
	"net/http"
	"regexp"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 17:28
 * @Description TODO
 * valid 参数校验中间件
 */

// 参数校验
func MiddleWare(params []Param) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramMap := make(map[string]interface{})
		if c.Request.Method == "GET" {
			for _, param := range params {
				value, ok := c.GetQuery(param.Name)
				if !ok {
					c.Abort()
					c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+constant.MSG_PARAM_NULL))
					return
				}
				res := handleParam(value, param.Methods)
				if res != "" {
					c.Abort()
					c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+res))
					return
				}
				paramMap[param.Name] = value
			}

		} else {
			if c.Request.Header.Get("Content-Type") == "application/json" {
				body, err := ioutil.ReadAll(c.Request.Body)
				if err != nil {
					c.Abort()
					log.Println(err)
					c.JSON(http.StatusOK, wrap.Error(constant.CODE_SERVER_EXCEPTION, "服务器异常"))
					return
				}
				err = json.Unmarshal(body, &paramMap)
				if err != nil {
					c.Abort()
					c.JSON(http.StatusOK, wrap.Error(constant.CODE_SERVER_EXCEPTION, "服务器异常"))
					return
				}
				for _, param := range params {
					_, ok := paramMap[param.Name]
					if !ok {
						c.Abort()
						c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+constant.MSG_PARAM_NULL))
						return
					}
					res := handleParam(paramMap[param.Name], param.Methods)
					if res != "" {
						c.Abort()
						c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+res))
						return
					}
				}
			} else {
				for _, param := range params {
					isFile := handleFile(param.Methods)
					if isFile {
						file, err := c.FormFile(param.Name)
						if err != nil {
							if "http: no such file" == err.Error() {
								c.Abort()
								log.Println(err)
								c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+constant.MSG_PARAM_NULL))
								return
							}
							c.Abort()
							log.Println(err)
							c.JSON(http.StatusOK, wrap.Error(constant.CODE_SERVER_EXCEPTION, "服务器异常"))
							return
						}
						paramMap[param.Name] = file
					} else {
						value, ok := c.GetPostForm(param.Name)
						if !ok {
							c.Abort()
							c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+constant.MSG_PARAM_NULL))
							return
						}
						res := handleParam(value, param.Methods)
						if res != "" {
							c.Abort()
							c.JSON(http.StatusOK, wrap.Error(constant.CODE_PARAM_FAILURE, param.Name+res))
							return
						}
						paramMap[param.Name] = value
					}
				}
			}
		}
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["params"] = wrap.NewWrapParams(paramMap)
		c.Next()
	}
}

// 获取参数
func GetWrapParams(c *gin.Context) *wrap.WrapParams {
	return c.Keys["params"].(*wrap.WrapParams)

}

// 处理参数
func handleParam(value interface{}, methods []Method) string {
	for _, method := range methods {
		switch method {
		case EMAIL:
			if compile := regexp.MustCompile(constant.REGULAR_MAIL); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return constant.MSG_PARAM_EMAIL
			}
		case NUMBER:
			if compile := regexp.MustCompile(constant.REGULAR_NUMBER); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return constant.MSG_PARAM_NUMBER
			}
		case PHONE:
			if compile := regexp.MustCompile(constant.REGULAR_PHONE); !compile.MatchString(fmt.Sprintf("%v", value)) {
				return constant.MSG_PARAM_PHONE
			}
		}
	}
	return ""
}

// 判断参数是否为文件
func handleFile(methods []Method) bool {
	if len(methods) == 1 && methods[0] == FILE {
		return true
	}
	return false
}
