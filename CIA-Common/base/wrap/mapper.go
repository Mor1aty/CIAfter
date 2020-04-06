package wrap

import "moriaty.com/cia/cia-common/base/constant"

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 12:57
 * @Description TODO
 * 	返回方法封装
 */

// 成功封装
func OK(code uint16, msg string, data ...interface{}) *Wrapper {
	return &Wrapper{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 执行成功封装
func OKExec(msg string, data ...interface{}) *Wrapper {
	return &Wrapper{
		Code: constant.CODE_EXEC_SUCCESS,
		Msg:  msg,
		Data: data,
	}
}

// 获取成功封装
func OKObtain(msg string, data interface{}) *Wrapper {
	return &Wrapper{
		Code: constant.CODE_OBTAIN_SUCCESS,
		Msg:  msg,
		Data: data,
	}
}

// 错误封装
func Error(code uint16, msg string) *Wrapper {
	return &Wrapper{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// 执行错误封装
func ErrorExec(msg string) *Wrapper {
	return &Wrapper{
		Code: constant.CODE_EXEC_FAILURE,
		Msg:  msg,
		Data: nil,
	}
}

// 获取错误封装
func ErrorObtain(msg string) *Wrapper {
	return &Wrapper{
		Code: constant.CODE_OBTAIN_FAILURE,
		Msg:  msg,
		Data: nil,
	}
}
