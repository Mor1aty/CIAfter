package auth

import (
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 19:41
 * @Description TODO
 * 	内存 Auth
 */

type memoryAuth struct {
	storage map[string]Token
}

func newMemoryAuth() *memoryAuth {
	return &memoryAuth{storage: make(map[string]Token)}
}

// 获取 Token
func (auth memoryAuth) CheckToken(tokenCode string) *Token {
	token, ok := auth.storage[tokenCode]
	if !ok {
		return nil
	}
	if token.EffectiveTime == -1 {
		// 永久存储
		return &token
	}
	if time.Now().After(token.GenerationTime.Add(time.Hour * time.Duration(token.EffectiveTime))) {
		// Token 过期
		delete(auth.storage, tokenCode)
		return nil
	}
	return &token
}

// 存入 Token
func (auth memoryAuth) PutTokenStorage(token Token) string {
	token.GenerationTime = time.Now()
	if token.EffectiveTime == 0 {
		token.EffectiveTime = 5
	}
	tokenCode := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	token.Code = tokenCode
	auth.storage[tokenCode] = token
	return tokenCode
}

// 移除 Token
func (auth memoryAuth) RemoveToken(tokenCode string) {
	delete(auth.storage, tokenCode)
}
