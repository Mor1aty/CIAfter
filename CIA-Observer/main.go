package main

import (
	"fmt"
	"os"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/5/4 18:03
 * @Description TODO
 * CIA-Observer
 */
func main() {
	err := os.RemoveAll("D:\\File\\基于 CI 的 APP 自动化测试系统\\CIAfter\\temp\\AAAAAA\\")
	if err != nil {
		fmt.Println(err)
	}
}
