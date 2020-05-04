package main

import (
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-publisher/service"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/19 18:40
 * @Description TODO
 * CIA-Publisher
 */
/**
1、拉取 apk
  1. 获取拉取配置
  2. 根据配置拉取 apk
2、整合 zip
  1. 根据配置将 apk 整合成 zip
  2. 调用任务生成
*/
func main() {
	server := micro.NewService(
		micro.Name("publisher"),
	)
	server.Init()
	service.Init(server)
	service.PullHandleApk()
	select {}
}

func init() {
	log.SetPrefix("【CIA-Publisher】")
	log.SetFlags(log.Lshortfile)
}
