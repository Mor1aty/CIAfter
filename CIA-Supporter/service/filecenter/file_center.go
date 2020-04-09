package filecenter

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	pb "moriaty.com/cia/cia-common/proto/supporter/filecenter"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao/filecenter"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 15:51
 * @Description TODO
 * 文件中心服务 service
 */

type SupporterFileCenter struct{}

// 注册服务
func Register(service micro.Service) error {
	// 注册 InsertFile（存入文件）
	err := pb.RegisterSupporterFileCenterHandler(service.Server(), new(SupporterFileCenter))
	if err != nil {
		log.Printf("register FileCenter failed, err: %v\n", err)
		return err
	}
	return nil
}

// 存入文件
func (sfc *SupporterFileCenter) InsertFile(ctx context.Context, req *pb.InsertFileReq, resp *pb.InsertFileResp) (err error) {
	id, err := filecenter.InsertFile(&bean.File{
		FileName:     req.File.FileName,
		FileLocation: req.File.FileLocation,
		Task:         req.File.Task,
		Secret:       req.File.Secret,
		CreateTime:   time.Now(),
	})
	if err != nil {
		log.Printf("insert file  %s(%s) failed, err: %v\n", req.File.FileName, req.File.FileLocation, err)
		return err
	}
	resp.Id = id
	log.Printf("insert file %s(%s)[%d] success\n", req.File.FileName, req.File.FileLocation, id)
	return nil
}

// 根据 id 获取文件
func (sfc *SupporterFileCenter) FindFileById(ctx context.Context, req *pb.FindFileByIdReq, resp *pb.FindFileByIdResp) (err error) {
	file, err := filecenter.FindFileById(req.Id)
	if err != nil {
		log.Printf("find file by id [%d] failed, err: %v\n", req.Id, err)
		return err
	}
	resp.File = &pb.File{
		Id:           file.Id,
		FileName:     file.FileName,
		FileLocation: file.FileLocation,
		Task:         file.Task,
		Secret:       file.Secret,
		CreateTime:   file.CreateTime.Format(constant.DATE_TYPE_01),
	}
	log.Printf("find file by id [%d] success\n", req.Id)
	return nil
}

// 根据任务 id 和权限 id 获取文件
func (sfc *SupporterFileCenter) FindFileByTaskAndSecret(ctx context.Context, req *pb.FindFileByTaskAndSecretReq, resp *pb.FindFileByTaskAndSecretResp) (err error) {
	tmpFiles, err := filecenter.FindFileByTaskAndSecret(req.Task, req.Secret)
	if err != nil {
		log.Printf("find file by task [%d] and secret [%s] failed, err: %v\n", req.Task, req.Secret, err)
		return err
	}
	files := make([]*pb.File, len(tmpFiles))
	for index, tmpFile := range tmpFiles {
		files[index] = &pb.File{
			Id:           tmpFile.Id,
			FileName:     tmpFile.FileName,
			FileLocation: tmpFile.FileLocation,
			Task:         tmpFile.Task,
			Secret:       tmpFile.Secret,
			CreateTime:   tmpFile.CreateTime.Format(constant.DATE_TYPE_01),
		}
	}
	resp.Files = files
	log.Printf("find file by task [%d] and secret [%s] success\n", req.Task, req.Secret)
	return nil
}
