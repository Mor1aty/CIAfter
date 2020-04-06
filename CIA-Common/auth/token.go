package auth

import "time"

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 19:35
 * @Description TODO
 * 	Auth 接口
 */

type Auth interface {
	CheckToken(tokenCode string) *Token
	PutTokenStorage(token Token) string
	RemoveToken(tokenCode string)
}

func NewAuth(name string) Auth {
	switch name {
	case "memory":
		return newMemoryAuth()
	default:
		return nil
	}
}

type Token struct {
	// token code
	Code string `json:"code"`
	// 有效时间，单位:时，-1 为永久存储
	EffectiveTime int16 `json:"effective_time"`
	// 生成时间
	GenerationTime time.Time `json:"generation_time"`
	// 权限等级
	Level uint8 `json:"level"`
	// 数据
	Data interface{} `json:"data"`
}
