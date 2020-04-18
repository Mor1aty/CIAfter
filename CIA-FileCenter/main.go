package main

import (
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-filecenter/service"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 10:48
 * @Description TODO
 *    CIA-FileCenter
 */
/**
CIA-FileCenter
 1、文件上传
  1. 提供文件上传的接口
  2. 文件分类更名存储
  3. 数据库存储文件信息
 2、文件获取
  1. 根据权限 id 与任务 id 获取文件
  2. 根据 id 获取文件
*/
func main() {

	// 启动服务
	server := micro.NewService(
		micro.Name("filecenter"),
	)
	server.Init()

	// 注册文件中心服务
	err := service.Register(server)
	if err != nil {
		log.Fatalf("register service failed, err: %v", err)
	}

	// 服务运行
	err = server.Run()
	if err := server.Run(); err != nil {
		log.Fatalf("run filecenter failed, err: %v\n", err)
	}

}

func init() {
	log.SetPrefix("【CIA-FileCenter】")
	log.SetFlags(log.Lshortfile)
}
