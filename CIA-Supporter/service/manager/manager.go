package manager

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	pb "moriaty.com/cia/cia-common/proto/supporter/manager"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/config"
	"moriaty.com/cia/cia-supporter/dao/filecenter"
	"moriaty.com/cia/cia-supporter/dao/manager"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/5/1 21:18
 * @Description TODO
 * 任务管理者服务 service
 */

var taskCfg *config.TaskConfig

type SupporterManager struct{}

// 注册服务
func Register(service micro.Service) error {
	err := pb.RegisterSupporterManagerHandler(service.Server(), new(SupporterManager))
	if err != nil {
		log.Printf("register Manager failed, err: %v", err)
		return err
	}
	return nil
}

// 根据权限 id 获取任务
func (sm *SupporterManager) FindTaskBySecret(ctx context.Context, req *pb.FindTaskBySecretReq, resp *pb.FindTaskBySecretResp) (err error) {
	tmpTasks, err := manager.FindTaskBySecret(req.Secret)
	if err != nil {
		log.Printf("find task by secret [%s] failed, err: %v", req.Secret, err)
		return err
	}
	if len(tmpTasks) == 0 {
		log.Printf("find task by secret [%s] not found", req.Secret)
		resp.Tasks = nil
		return nil
	}

	tasks := make([]*pb.Task, len(tmpTasks))
	for index, tmpTask := range tmpTasks {
		tmpFile, err := filecenter.FindFileById(tmpTask.File)
		var file *pb.File
		if err != nil {
			if "sql: no rows in result set" == err.Error() {
				log.Printf("find file by id [%d] not found", tmpTask.File)
				file = nil
			} else {
				log.Printf("find file by id [%d] failed, err: %v", tmpTask.File, err)
				file = nil
			}
		} else {
			file = &pb.File{
				Id:           tmpFile.Id,
				FileName:     tmpFile.FileName,
				FileLocation: tmpFile.FileLocation,
				Secret:       tmpFile.Secret,
				CreateTime:   tmpFile.CreateTime.Format(constant.DATE_TYPE_01),
			}
		}
		tasks[index] = &pb.Task{
			Id:                  tmpTask.Id,
			Secret:              tmpTask.Secret,
			TestType:            tmpTask.TestType,
			Client:              tmpTask.Client,
			Desc:                tmpTask.Desc,
			File:                file,
			Status:              tmpTask.Status,
			Result:              tmpTask.Result,
			ResultDesc:          tmpTask.ResultDesc,
			ResultLocation:      tmpTask.ResultLocation,
			ResultImageLocation: tmpTask.ResultImageLocation,
			CreateTime:          tmpTask.CreateTime.Format(constant.DATE_TYPE_01),
			StartTime:           tmpTask.StartTime.Format(constant.DATE_TYPE_01),
			EndTime:             tmpTask.EndTime.Format(constant.DATE_TYPE_01),
		}
	}
	resp.Tasks = tasks
	log.Printf("find task by secret [%s] success", req.Secret)
	return nil
}

// 插入任务
func (sm *SupporterManager) InsertTask(ctx context.Context, req *pb.InsertTaskReq, resp *pb.InsertTaskResp) (err error) {
	id, err := manager.InsertTask(req.Secret, req.TestType, req.Desc, req.File, time.Now())
	if err != nil {
		log.Printf("insert task %s(%s) failed, err: %v", req.TestType, req.Secret, err)
		return err
	}
	resp.Id = id
	log.Printf("insert task %s(%s)[%d] success", req.TestType, req.Secret, id)
	return nil
}

// 推送任务
func (sm *SupporterManager) PushTask(ctx context.Context, req *pb.PushTaskReq, resp *pb.PushTaskResp) (err error) {
	err = manager.PushTask(taskCfg.Key, req.Secret, &bean.TaskRedis{
		TaskId:   req.TaskId,
		TaskType: req.TaskType,
		TaskFile: req.TaskFile,
	})
	if err != nil {
		log.Printf("push task[%d] failed, err: %v", req.TaskId, err)
		return err
	}
	return nil
}
