package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "moriaty.com/cia/cia-common/proto/filecenter"
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
	service := micro.NewService(
		micro.Name("manager"),
	)
	service.Init()

	findFile := pb.NewFindFileService("supporter", service.Client())
	findFileResp, err := findFile.FindFileById(context.TODO(), &pb.FindFileByIdReq{Id: 1})
	if err != nil {
		fmt.Printf("call supporter FindFile failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", findFileResp.File)

	findTaskFile := pb.NewFindTaskFileService("supporter", service.Client())
	findTaskFileResp, err := findTaskFile.FindFileByTaskAndSecret(context.TODO(), &pb.FindFileByTaskAndSecretReq{Task: 1, Secret: "FFFFFF"})
	if err != nil {
		fmt.Printf("call supporter FindTaskFile failed, err: %v\n", err)
		return
	}
	for _, file := range findTaskFileResp.Files {
		fmt.Printf("%#v\n", file)
	}

	insertFile := pb.NewInsertFileService("supporter", service.Client())
	insertFileResp, err := insertFile.InsertFile(context.TODO(), &pb.InsertFileReq{File: &pb.File{
		FileName:     "测试文件3",
		FileLocation: "/zz/zz/zz",
		Task:         2,
		Secret:       "AAAAAA",
	}})
	if err != nil {
		fmt.Printf("call supporter InsertFile failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", insertFileResp.Id)
}
