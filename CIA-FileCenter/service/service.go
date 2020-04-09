package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
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

var currentClient client.Client

type FileCenter struct{}

// 注册服务
func Register(server micro.Service) error {
	currentClient = server.Client()
	// 注册 Upload（文件上传）
	err := pb.RegisterFileCenterHandler(server.Server(), new(FileCenter))
	if err != nil {
		log.Printf("register Upload failed, err: %v\n", err)
		return err
	}
	return nil
}

// 文件上传
func (fc *FileCenter) Upload(ctx context.Context, req *pb.UploadReq, resp *pb.UploadResp) (err error) {
	newLocation := fmt.Sprintf("D:\\File\\基于 CI 的 APP 自动化测试系统\\CIAfter\\file\\%s_%s",
		strings.ReplaceAll(uuid.NewV4().String(), "-", ""), req.FileName)

	err = handleFile(req.TempFileLocation, newLocation)
	removeFile(req.TempFileLocation)
	if err != nil {
		log.Printf("upload %s(%s) failed, err: %v\n", req.FileName, req.TempFileLocation, err)
		removeFile(newLocation)
		return
	}

	sfc := supporter.NewSupporterFileCenterService("supporter", currentClient)
	insertFileResp, err := sfc.InsertFile(context.TODO(), &supporter.InsertFileReq{File: &supporter.File{
		FileName:     req.FileName,
		FileLocation: newLocation,
		Task:         req.Task,
		Secret:       req.Secret,
	}})
	if err != nil {
		fmt.Printf("call supporter InsertFile failed, err: %v\n", err)
		removeFile(newLocation)
		return
	}

	log.Printf("upload %s(%s)[%d] success\n", req.FileName, newLocation, insertFileResp.Id)
	return
}

// 处理文件
func handleFile(tempLocation, newLocation string) error {
	readFile, err := os.Open(tempLocation)
	if err != nil {
		log.Printf("open %s failed, err: %v\n", tempLocation, err)
		return err
	}
	defer readFile.Close()

	writeFile, err := os.OpenFile(newLocation, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("open %s failed, err: %v\n", newLocation, err)
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
			log.Printf("%s handle failed, err: %v\n", newLocation, err)
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
		log.Printf("remove %s failed\n", location)
	}
}
