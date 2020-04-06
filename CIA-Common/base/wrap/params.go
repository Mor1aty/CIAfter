package wrap

import (
	"encoding/json"
	"log"
	"mime/multipart"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 12:58
 * @Description TODO
 * 	参数封装
 */

type WrapParams struct {
	param map[string]interface{}
}

func NewWrapParams(param map[string]interface{}) *WrapParams {
	return &WrapParams{param: param}
}

func (wp WrapParams) String() string {
	paramJson, err := json.Marshal(wp.param)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(paramJson)
}

func (wp WrapParams) Json() {

}

// 存入字段
func (wp WrapParams) Put(key string, value interface{}) {
	wp.param[key] = value
}

// 获取参数数量
func (wp WrapParams) Size() int {
	return len(wp.param)
}

// 判断参数是否为空
func (wp WrapParams) IsEmpty() bool {
	return len(wp.param) == 0
}

// 判断参数是否包含 key
func (wp WrapParams) ContainsKey(key string) bool {
	_, ok := wp.param[key]
	return ok
}

// 获取 key 获取对应 value
func (wp WrapParams) Get(key string) interface{} {
	return wp.param[key]
}

// 根据 key 获取对应 value，转化为 boolean
func (wp WrapParams) GetBoolValue(key string) bool {
	return wp.param[key].(bool)
}

// 根据 key 获取对应 value，转化为 int
func (wp WrapParams) GetIntValue(key string) int {
	return wp.param[key].(int)
}

// 根据 key 获取对应 value，转化为 string
func (wp WrapParams) GetString(key string) string {
	return wp.param[key].(string)
}

// 根据 key 获取对应 value，转化为 File
func (wp WrapParams) GetFileValue(key string) *multipart.FileHeader {
	return wp.param[key].(*multipart.FileHeader)
}

// 清除所有参数
func (wp WrapParams) Clear() {
	wp.param = make(map[string]interface{})
}

// 删除 key 对应的参数
func (wp WrapParams) remove(key string) {
	delete(wp.param, key)
}

// 获取参数 key 集合
func (wp WrapParams) KeySet() []string {
	keys := make([]string, 0, len(wp.param))
	for key := range wp.param {
		keys = append(keys, key)
	}
	return keys
}
