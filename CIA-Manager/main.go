package main

import (
	"context"
	"github.com/micro/go-micro"
	"moriaty.com/cia/cia-common/proto/filecenter"
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
	任务管理者 CIA-Manager
  	1. 提供生成机制
  	2. 生成任务
  	3. 将任务分发到任务容器，方便任务执行者执行任务
  	4. 提供任务结果获取机制
*/

func main() {
	server := micro.NewService(
		micro.Name("manager"),
	)
	server.Init()

	fc := filecenter.NewFileCenterService("filecenter", server.Client())
	fc.Upload(context.TODO(), &filecenter.UploadReq{
		FileName:         "README.md",
		TempFileLocation: "D:\\File\\基于 CI 的 APP 自动化测试系统\\CIAfter\\temp\\README.md",
		Task:             4,
		Secret:           "FFFFFF",
	})
}
