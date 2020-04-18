package main

import (
	"github.com/micro/go-micro"
	"gopkg.in/ini.v1"
	"log"
	"moriaty.com/cia/cia-supporter/config"
	"moriaty.com/cia/cia-supporter/dao"
	"moriaty.com/cia/cia-supporter/service/executor"
	"moriaty.com/cia/cia-supporter/service/filecenter"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 10:59
 * @Description TODO
 * CIA-Supporter
 */
/**
1、数据库操作服务支持
  1. 文件中心服务支持
  2. 任务执行者服务支持
2、Redis 操作服务支持
  1. 任务执行者服务支持
*/
func main() {
	// 加载配置文件
	cfg := new(config.SupporterConfig)
	err := ini.MapTo(&cfg, "./config/config.ini")
	if err != nil {
		log.Fatalf("load ini failed, err: %v", err)
	}
	log.Println("config success")

	// 初始化 mysql
	err = dao.InitMysql(&cfg.MysqlConfig)
	if err != nil {
		log.Fatalf("init mysql failed, err: %v", err)
	}
	log.Println("init mysql success")

	// 初始化 redis
	//err = dao.InitRedis(&cfg.RedisConfig)
	//if err != nil {
	//	log.Fatalf("init redis failed, err: %v", err)
	//}
	log.Println("init redis success")

	// 启动服务
	server := micro.NewService(
		micro.Name("supporter"),
	)
	server.Init()

	// 注册文件中心服务
	err = filecenter.Register(server)
	if err != nil {
		log.Fatalf("register service filecenter failed, err: %v", err)
	}

	// 注册任务执行者服务
	err = executor.Register(server, &cfg.TaskConfig)
	if err != nil {
		log.Fatalf("register service executor failed, err: %v", err)
	}

	// 服务运行
	err = server.Run()
	if err := server.Run(); err != nil {
		log.Fatalf("run supporter failed, err: %v\n", err)
	}
}

func init() {
	log.SetPrefix("【CIA-Supporter】")
	log.SetFlags(log.Lshortfile)
}
