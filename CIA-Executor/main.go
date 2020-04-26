package main

import (
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-executor/service"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/12 17:04
 * @Description TODO
 * CIA-Executor
 */
/**
1、初始化手机信息
  1. 初始化时将插在客户端上的手机信息上传数据库
2、执行任务
  1. 获取执行的任务
  2. 执行任务
  3. 将任务结果上传到文件中心
*/

func main() {
	// 启动服务
	server := micro.NewService(
		micro.Name("executor"),
	)
	server.Init()

	// 初始化 service
	err := service.Init(server)
	if err != nil {
		log.Fatalf("init service failed, err: %v", err)
	}

	err = service.InitPhone()
	if err != nil {
		log.Fatalf("init phone failed, err: %v", err)
	}

	service.PushTestTask()
	log.Println("push data success")

	go service.ConsumeTask()

	select {}
}

func init() {
	log.SetPrefix("【CIA-Supporter】")
	log.SetFlags(log.Lshortfile)
}
