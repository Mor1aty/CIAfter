package main

import (
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-manager/service"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 17:22
 * @Description TODO
 * CIA-Manager
 */
/**
1、生成任务
  1. 组织任务信息
  2. 任务信息存入数据库
  3. 任务信息存入任务容器
2、任务信息
  1. 获取任务执行信息
  2. 获取任务执行结果
*/

func main() {
	// 启动服务
	server := micro.NewService(
		micro.Name("manager"),
	)
	server.Init()

	// 注册任务管理者服务
	err := service.Register(server)
	if err != nil {
		log.Fatalf("register service failed, err: %v", err)
	}

	// 服务运行
	err = server.Run()
	if err := server.Run(); err != nil {
		log.Fatalf("run manager failed, err: %v", err)
	}
}

func init() {
	log.SetPrefix("【CIA-Manager】")
	log.SetFlags(log.Lshortfile)
}
