package executor

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	pb "moriaty.com/cia/cia-common/proto/supporter/executor"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/config"
	"moriaty.com/cia/cia-supporter/dao/executor"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/14 14:10
 * @Description TODO
 * 任务执行者服务 service
 */

var taskCfg *config.TaskConfig

type SupporterExecutor struct{}

// 注册服务
func Register(service micro.Service, cfg *config.TaskConfig) error {
	err := pb.RegisterSupporterExecutorHandler(service.Server(), new(SupporterExecutor))
	if err != nil {
		log.Printf("register Executor failed, err: %v\n", err)
		return err
	}
	taskCfg = cfg
	return nil
}

// 获取任务 key
func (se *SupporterExecutor) FindTaskKey(ctx context.Context, req *pb.FindTaskKeyReq, resp *pb.FindTaskKeyResp) (err error) {
	idsKey := fmt.Sprintf("%s_ids", taskCfg.Key)
	id, err := executor.PopId(idsKey)
	if err != nil {
		if "redis: nil" == err.Error() {
			log.Printf("task key not found\n")
			return nil
		}
		log.Printf("find task key [%s] failed, err: %v\n", idsKey, err)
		return err
	}
	resp.Key = fmt.Sprintf("%s_%d", taskCfg.Key, id)
	log.Printf("find task key [%s] success\n", idsKey)
	return nil
}

// 获取任务
func (se *SupporterExecutor) FindTaskByKey(ctx context.Context, req *pb.FindTaskByKeyReq, resp *pb.FindTaskByKeyResp) (err error) {
	task, err := executor.FindTaskByKey(req.Key)
	if err != nil {
		if "redis: nil" == err.Error() {
			log.Printf("task by key [%s] not found\n", req.Key)
			resp.Task = nil
			return nil
		}
		log.Printf("find task by key [%s] failed, err: %v\n", req.Key, err)
		return err
	}
	resp.Task = &pb.Task{
		TaskId:   task.TaskId,
		TaskType: task.TaskType,
		TaskFile: task.TaskFile,
	}
	log.Printf("find task by key [%s] success\n", req.Key)
	return nil
}

// 删除任务
func (se *SupporterExecutor) DeleteTaskByKey(ctx context.Context, req *pb.DeleteTaskByKeyReq, resp *pb.DeleteTaskByKeyResp) (err error) {
	err = executor.DeleteTaskByKey(req.Key)
	if err != nil {
		log.Printf("delete task by key [%s] failed, err: %v\n", req.Key, err)
		return err
	}
	log.Printf("delete task by key [%s] success\n", req.Key)
	return nil
}

// 推送测试任务
func (se *SupporterExecutor) PushTestTask(ctx context.Context, req *pb.PushTestTaskReq, resp *pb.PushTestTaskResp) (err error) {
	executor.PushTestTask()
	return nil
}

// 存入手机信息
func (se *SupporterExecutor) InsertPhone(ctx context.Context, req *pb.InsertPhoneReq, resp *pb.InsertPhoneResp) (err error) {
	for _, phone := range req.Phones {
		tempPhone, err := executor.FindPhoneById(phone.Id)
		if err != nil {
			if "sql: no rows in result set" == err.Error() {
				err = executor.InsertPhone(&bean.Phone{
					Id:         phone.Id,
					Client:     phone.Client,
					TestType:   phone.TestType,
					Secret:     phone.Secret,
					UpdateTime: time.Now(),
				})
				if err != nil {
					log.Printf("insert phone %s failed, err: %v\n", phone.Id, err)
					return err
				}
			} else {
				log.Printf("find phone by id [%s] failed, err: %v\n", phone.Id, err)
				return err
			}
		} else {
			if tempPhone.Client != phone.Client {
				err = executor.UpdatePhoneClient(phone.Id, phone.Client)
				if err != nil {
					log.Printf("update phone [%s] client [%s] failed, err: %v\n", phone.Id, phone.Client, err)
					return err
				}
			}
		}
	}
	log.Printf("insert phones success\n")
	return nil
}

// 根据 ip 获取客户端
func (se *SupporterExecutor) FindClientByIp(ctx context.Context, req *pb.FindClientByIpReq, resp *pb.FindClientByIpResp) (err error) {
	client, err := executor.FindClientByIp(req.Ip)
	if err != nil {
		if "sql: no rows in result set" == err.Error() {
			log.Printf("find client by ip [%s] not found\n", req.Ip)
			resp.Client = nil
			return nil
		}
		log.Printf("find client by ip [%s] failed, err: %v\n", req.Ip, err)
		return err
	}
	resp.Client = &pb.Client{
		Ip:         client.Ip,
		TestType:   client.TestType,
		Secret:     client.Secret,
		UpdateTime: client.UpdateTime.Format(constant.DATE_TYPE_01),
	}
	log.Printf("find client by ip [%s] success\n", req.Ip)
	return
}
