package service

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	fileCenter "moriaty.com/cia/cia-common/proto/filecenter"
	pb "moriaty.com/cia/cia-common/proto/manager"
	supporter "moriaty.com/cia/cia-common/proto/supporter/manager"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/5/4 15:24
 * @Description TODO
 * 	任务管理者 service
 */

var (
	sms supporter.SupporterManagerService
	fcs fileCenter.FileCenterService
)

type Manager struct{}

func Register(server micro.Service) error {
	sms = supporter.NewSupporterManagerService("supporter", server.Client())
	fcs = fileCenter.NewFileCenterService("filecenter", server.Client())
	err := pb.RegisterManagerHandler(server.Server(), new(Manager))
	if err != nil {
		log.Printf("register Upload failed, err: %v", err)
		return err
	}
	return nil
}

// 创建任务
func (m *Manager) CreateTask(ctx context.Context, req *pb.CreateTaskReq, resp *pb.CreateTaskResp) (err error) {
	uploadResp, err := fcs.Upload(context.TODO(), &fileCenter.UploadReq{
		FileName:         req.FileName,
		TempFileLocation: req.FileLocation,
		Secret:           req.Secret,
	})
	if err != nil {
		log.Printf("call filecenter Upload failed, err: %v", err)
		return err
	}
	insertTaskResp, err := sms.InsertTask(context.TODO(), &supporter.InsertTaskReq{
		Secret:   req.Secret,
		TestType: req.TestType,
		Desc:     req.Desc,
		File:     uploadResp.Id,
	})
	if err != nil {
		log.Printf("insert task [%d] failed, err: %v", uploadResp.Id, err)
		return err
	}
	_, err = sms.PushTask(context.TODO(), &supporter.PushTaskReq{
		TaskId:   insertTaskResp.Id,
		TaskType: req.TestType,
		TaskFile: uploadResp.FileLocation,
	})
	if err != nil {
		log.Printf("push task failed, err: %v", err)
		return err
	}
	resp.Id = insertTaskResp.Id
	log.Printf("create task [%d] success", insertTaskResp.Id)
	return nil
}
