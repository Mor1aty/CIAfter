package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	pb "moriaty.com/cia/cia-common/proto/filecenter"
	supporter "moriaty.com/cia/cia-common/proto/supporter/filecenter"
	"os"
	"strings"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/9 20:37
 * @Description TODO
 *  文件中心 service
 */

var sfc supporter.SupporterFileCenterService

type FileCenter struct{}

// 注册服务
func Register(server micro.Service) error {

	sfc = supporter.NewSupporterFileCenterService("supporter", server.Client())

	err := pb.RegisterFileCenterHandler(server.Server(), new(FileCenter))
	if err != nil {
		log.Printf("register Upload failed, err: %v", err)
		return err
	}
	return nil
}

// 文件上传
func (fc *FileCenter) Upload(ctx context.Context, req *pb.UploadReq, resp *pb.UploadResp) (err error) {
	newLocation := fmt.Sprintf("D:\\File\\基于 CI 的 APP 自动化测试系统\\CIAfter\\file\\%s\\%d\\%s_%s",
		req.Secret, req.Task, strings.ReplaceAll(uuid.NewV4().String(), "-", ""), req.FileName)

	err = handleFile(req.TempFileLocation, newLocation)
	removeFile(req.TempFileLocation)
	if err != nil {
		log.Printf("upload %s(%s) failed, err: %v", req.FileName, req.TempFileLocation, err)
		removeFile(newLocation)
		return
	}
	insertFileResp, err := sfc.InsertFile(context.TODO(), &supporter.InsertFileReq{File: &supporter.File{
		FileName:     req.FileName,
		FileLocation: newLocation,
		Task:         req.Task,
		Secret:       req.Secret,
	}})
	if err != nil {
		log.Printf("call supporter InsertFile failed, err: %v", err)
		removeFile(newLocation)
		return
	}

	log.Printf("upload %s(%s)[%d] success", req.FileName, newLocation, insertFileResp.Id)
	return
}

// 处理文件
func handleFile(tempLocation, newLocation string) error {
	readFile, err := os.Open(tempLocation)
	if err != nil {
		log.Printf("open %s failed, err: %v", tempLocation, err)
		return err
	}
	defer readFile.Close()

	writeFile, err := os.OpenFile(newLocation, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("open %s failed, err: %v", newLocation, err)
		return err
	}
	defer writeFile.Close()

	reader := bufio.NewReader(readFile)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				writeFile.WriteString(line)
			}
			break
		}
		if err != nil {
			log.Printf("%s handle failed, err: %v", newLocation, err)
			return err
		}
		writeFile.WriteString(line)
	}
	return nil
}

// 删除文件
func removeFile(location string) {
	err := os.Remove(location)
	if err != nil {
		log.Printf("remove %s failed", location)
	}
}

// 根据 id 获取文件
func (fc *FileCenter) FindFile(ctx context.Context, req *pb.FindFileReq, resp *pb.FindFileResp) (err error) {
	findFileByIdResp, err := sfc.FindFileById(context.TODO(), &supporter.FindFileByIdReq{Id: req.Id})
	if err != nil {
		log.Printf("call supporter FindFileById failed, err: %v", err)
		return
	}
	if findFileByIdResp.File == nil {
		resp.File = nil
	} else {
		resp.File = &pb.File{
			Id:           findFileByIdResp.File.Id,
			FileName:     findFileByIdResp.File.FileName,
			FileLocation: findFileByIdResp.File.FileLocation,
			CreateTime:   findFileByIdResp.File.CreateTime,
		}
	}
	log.Printf("FindFile id[%d] success", req.Id)
	return
}

// 根据权限 id 与任务 id 获取文件
func (fc *FileCenter) FindTaskFile(ctx context.Context, req *pb.FindTaskFileReq, resp *pb.FindTaskFileResp) (err error) {
	findTaskFileResp, err := sfc.FindFileByTaskAndSecret(context.TODO(), &supporter.FindFileByTaskAndSecretReq{Task: req.Task, Secret: req.Secret})
	if err != nil {
		log.Printf("call supporter FindTaskFileResp failed, err: %v", err)
		return
	}
	files := make([]*pb.File, len(findTaskFileResp.Files))
	for index, file := range findTaskFileResp.Files {
		files[index] = &pb.File{
			Id:           file.Id,
			FileName:     file.FileName,
			FileLocation: file.FileLocation,
			CreateTime:   file.CreateTime,
		}
	}
	resp.Files = files
	log.Printf("FindFile task[%d] secret[%s] success", req.Task, req.Secret)
	return
}
