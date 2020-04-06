package valid

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/6 13:12
 * @Description TODO
 * valid 中使用的参数实体
 */

type Method uint8

const (
	NOTNULL Method = iota
	EMAIL
	NUMBER
	PHONE
	FILE
)

type Param struct {
	Name    string
	Methods []Method
}
