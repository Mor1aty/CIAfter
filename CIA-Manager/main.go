package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
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
	ret, err := fc.FindFile(context.TODO(), &filecenter.FindFileReq{Id: 10})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", ret.File)

	ret1, err := fc.FindTaskFile(context.TODO(), &filecenter.FindTaskFileReq{Task: 3, Secret: "AAaAAAA"})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range ret1.Files {
		fmt.Printf("%#v\n", file)
	}

}

func init() {
	log.SetPrefix("【CIA-Manager】")
	log.SetFlags(log.Lshortfile)
}
