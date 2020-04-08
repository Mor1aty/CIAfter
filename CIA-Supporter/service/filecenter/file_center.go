package filecenter

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	pb "moriaty.com/cia/cia-common/proto/filecenter"
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

// 注册服务
func Register(service micro.Service) error {
	// 注册 InsertFile（存入文件）
	err := pb.RegisterInsertFileHandler(service.Server(), new(InsertFile))
	if err != nil {
		log.Printf("register InsertFile failed, err: %v\n", err)
		return err
	}

	// 注册 FindFile（根据 id 获取文件）
	err = pb.RegisterFindFileHandler(service.Server(), new(FindFile))
	if err != nil {
		log.Printf("register FindFile failed, err: %v\n", err)
		return err
	}

	// 注册 FindTaskFile（根据任务 id 和权限 id 获取文件）
	err = pb.RegisterFindTaskFileHandler(service.Server(), new(FindTaskFile))
	if err != nil {
		log.Printf("register FindTaskFile failed, err: %v\n", err)
		return err
	}
	return nil
}

// 存入文件
type InsertFile struct{}

func (insertFile *InsertFile) InsertFile(ctx context.Context, req *pb.InsertFileReq, resp *pb.InsertFileResp) (err error) {
	id, err := filecenter.InsertFile(&bean.File{
		FileName:     req.File.FileName,
		FileLocation: req.File.FileLocation,
		Task:         req.File.Task,
		Secret:       req.File.Secret,
		CreateTime:   time.Now(),
	})
	if err != nil {
		log.Printf("insert file failed, err: %v\n", err)
		return err
	}
	resp.Id = id
	return nil
}

// 根据 id 获取文件
type FindFile struct{}

func (findFile *FindFile) FindFileById(ctx context.Context, req *pb.FindFileByIdReq, resp *pb.FindFileByIdResp) (err error) {
	file, err := filecenter.FindFileById(req.Id)
	if err != nil {
		log.Printf("find file by id failed, err: %v\n", err)
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
	return nil
}

// 根据任务 id 和权限 id 获取文件
type FindTaskFile struct{}

func (findTaskFile *FindTaskFile) FindFileByTaskAndSecret(ctx context.Context, req *pb.FindFileByTaskAndSecretReq, resp *pb.FindFileByTaskAndSecretResp) (err error) {
	tmpFiles, err := filecenter.FindFileByTaskAndSecret(req.Task, req.Secret)
	if err != nil {
		log.Printf("find file by task and secret failed, err: %v\n", err)
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
	return nil
}
