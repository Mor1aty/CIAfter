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
		log.Printf("register Executor failed, err: %v", err)
		return err
	}
	taskCfg = cfg
	return nil
}

// 获取任务 key
func (se *SupporterExecutor) FindTaskKey(ctx context.Context, req *pb.FindTaskKeyReq, resp *pb.FindTaskKeyResp) (err error) {
	idsKey := fmt.Sprintf("%s_ids_%s", taskCfg.Key, req.Secret)
	id, err := executor.PopId(idsKey)
	if err != nil {
		if "redis: nil" == err.Error() {
			log.Printf("task key not found")
			return nil
		}
		log.Printf("find task key [%s] failed, err: %v", idsKey, err)
		return err
	}
	resp.Key = fmt.Sprintf("%s_%s_%d", taskCfg.Key, req.Secret, id)
	log.Printf("find task key [%s] success", idsKey)
	return nil
}

// 获取任务
func (se *SupporterExecutor) FindTaskByKey(ctx context.Context, req *pb.FindTaskByKeyReq, resp *pb.FindTaskByKeyResp) (err error) {
	task, err := executor.FindTaskByKey(req.Key)
	if err != nil {
		if "redis: nil" == err.Error() {
			log.Printf("task by key [%s] not found", req.Key)
			resp.Task = nil
			return nil
		}
		log.Printf("find task by key [%s] failed, err: %v", req.Key, err)
		return err
	}
	resp.Task = &pb.Task{
		TaskId:   task.TaskId,
		TaskType: task.TaskType,
		TaskFile: task.TaskFile,
	}
	log.Printf("find task by key [%s] success", req.Key)
	return nil
}

// 删除任务
func (se *SupporterExecutor) DeleteTaskByKey(ctx context.Context, req *pb.DeleteTaskByKeyReq, resp *pb.DeleteTaskByKeyResp) (err error) {
	err = executor.DeleteTaskByKey(req.Key)
	if err != nil {
		log.Printf("delete task by key [%s] failed, err: %v", req.Key, err)
		return err
	}
	log.Printf("delete task by key [%s] success", req.Key)
	return nil
}

// 存入手机信息
func (se *SupporterExecutor) InsertPhone(ctx context.Context, req *pb.InsertPhoneReq, resp *pb.InsertPhoneResp) (err error) {
	for _, phone := range req.Phones {
		tempPhone, err := executor.FindPhoneById(phone.Id)
		if err != nil {
			if "sql: no rows in result set" == err.Error() {
				log.Printf("%#v", phone)
				err = executor.InsertPhone(&bean.Phone{
					Id:           phone.Id,
					Client:       phone.Client,
					TestType:     phone.TestType,
					Secret:       phone.Secret,
					PhoneType:    phone.PhoneType,
					PhoneEdition: phone.PhoneEdition,
					UpdateTime:   time.Now(),
				})
				if err != nil {
					log.Printf("insert phone %s failed, err: %v", phone.Id, err)
					return err
				}
			} else {
				log.Printf("find phone by id [%s] failed, err: %v", phone.Id, err)
				return err
			}
		} else {
			if tempPhone.Client != phone.Client {
				err = executor.UpdatePhoneClient(phone.Id, phone.Client)
				if err != nil {
					log.Printf("update phone [%s] client [%s] failed, err: %v", phone.Id, phone.Client, err)
					return err
				}
			}
		}
	}
	log.Printf("insert phones success")
	return nil
}

// 根据 ip 获取客户端
func (se *SupporterExecutor) FindClientByIp(ctx context.Context, req *pb.FindClientByIpReq, resp *pb.FindClientByIpResp) (err error) {
	client, err := executor.FindClientByIp(req.Ip)
	if err != nil {
		if "sql: no rows in result set" == err.Error() {
			log.Printf("find client by ip [%s] not found", req.Ip)
			resp.Client = nil
			return nil
		}
		log.Printf("find client by ip [%s] failed, err: %v", req.Ip, err)
		return err
	}
	resp.Client = &pb.Client{
		Ip:         client.Ip,
		TestType:   client.TestType,
		Secret:     client.Secret,
		UpdateTime: client.UpdateTime.Format(constant.DATE_TYPE_01),
	}
	log.Printf("find client by ip [%s] success", req.Ip)
	return
}

// 更新任务开始信息
func (se *SupporterExecutor) UpdateTaskStartById(ctx context.Context, req *pb.UpdateTaskStartByIdReq, resp *pb.UpdateTaskStartByIdResp) (err error) {
	err = executor.UpdateTaskStartById(req.Id, req.Client)
	if err != nil {
		log.Printf("upload task [%d] start failed, err: %v", req.Id, err)
		return err
	}
	return nil
}

// 更新任务结束信息
func (se *SupporterExecutor) UpdateTaskEndById(ctx context.Context, req *pb.UpdateTaskEndByIdReq, resp *pb.UpdateTaskEndByIdResp) (err error) {
	result := 1
	if req.IsSuccess {
		result = 0
	}
	err = executor.UpdateTaskEndById(req.Id, result, req.ResultDesc, req.ResultLocation, req.ResultImageLocation)
	if err != nil {
		log.Printf("upload task [%d] end failed, err: %v", req.Id, err)
		return err
	}
	return nil
}
