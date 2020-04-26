package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/proto/supporter/publisher"
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

	sp := publisher.NewSupporterPublisherService("supporter", server.Client())
	findAllBusinessResp, err := sp.FindAllBusiness(context.TODO(), &publisher.FindAllBusinessReq{All: true})
	if err != nil {
		log.Printf("call supporter publisher failed, err: %v", err)
	}
	for _, business := range findAllBusinessResp.Businesses {
		fmt.Printf("%#v", business)
	}
	fmt.Println("--------------------------------------------")

	findAllBusinessResp, err = sp.FindAllBusiness(context.TODO(), &publisher.FindAllBusinessReq{IsStop: true})
	if err != nil {
		log.Printf("call supporter publisher failed, err: %v", err)
	}
	for _, business := range findAllBusinessResp.Businesses {
		fmt.Printf("%#v\n", business)
	}
	fmt.Println("--------------------------------------------")

	findAllBusinessResp, err = sp.FindAllBusiness(context.TODO(), &publisher.FindAllBusinessReq{})
	if err != nil {
		log.Printf("call supporter publisher failed, err: %v", err)
	}
	for _, business := range findAllBusinessResp.Businesses {
		fmt.Printf("%#v\n", business)
	}
	fmt.Println("--------------------------------------------")

}

func init() {
	log.SetPrefix("【CIA-Manager】")
	log.SetFlags(log.Lshortfile)
}
