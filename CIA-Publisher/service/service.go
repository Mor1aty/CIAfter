package service

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"io"
	"io/ioutil"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	"moriaty.com/cia/cia-common/proto/manager"
	supporter "moriaty.com/cia/cia-common/proto/supporter/publisher"
	"moriaty.com/cia/cia-publisher/bean"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/24 16:56
 * @Description TODO
 * 任务发布者 service
 */

var (
	sp supporter.SupporterPublisherService
	ms manager.ManagerService
)

// 初始化 service
func Init(server micro.Service) {
	sp = supporter.NewSupporterPublisherService("supporter", server.Client())
	ms = manager.NewManagerService("manager", server.Client())
}

// 拉取处理 apk
func PullHandleApk() {
	findAllBusinessResp, err := sp.FindAllBusiness(context.TODO(), &supporter.FindAllBusinessReq{})
	if err != nil {
		log.Printf("call supporter FindAllBusiness failed, err: %v", err)
		return
	}
	for _, business := range findAllBusinessResp.Businesses {
		business := business
		go func() {
			pullHandleApk(business)
		}()
	}
}

func pullHandleApk(business *supporter.Business) {
	if business.BusinessTests == nil {
		log.Printf("business test with secret [%s] not found", business.Secret)
		return
	}
	_, err := os.Stat(constant.TEMP_FILE_LOCATION + business.Secret)
	if os.IsNotExist(err) {
		err = os.Mkdir(constant.TEMP_FILE_LOCATION+business.Secret, os.ModePerm)
		if err != nil {
			log.Printf("create dir [%s] failed, err: %v", business.Secret, err)
			return
		}
	}
	for {
		for _, businessTest := range business.BusinessTests {
			fileInfo, err := callFileInfo(fmt.Sprintf("%sfile/%s", business.FileUrl, business.ApkName))
			if err != nil {
				log.Printf("get file info failed, err: %v", err)
				continue
			}
			if fileInfo.Edition > business.CurrentEdition {
				tmpFileLocation, err := pullFile(fmt.Sprintf("%s%s", business.FileUrl, fileInfo.Name), business.Secret, fileInfo.Name)
				if err != nil {
					log.Printf("get the newest file failed, err: %v", err)
					continue
				}
				business.CurrentEdition = fileInfo.Edition
				switch businessTest.TestOption {
				case constant.TEST_OPTION_SINGLE:
					tmpZipName, tmpZipLocation, err := singleFileHandle(tmpFileLocation, businessTest.TestType)
					if err != nil {
						log.Printf("compress file failed, err: %v", err)
						continue
					}
					log.Printf("single file handle success, zip file: %s(%s)", tmpZipName, tmpZipLocation)
					createTaskResp, err := ms.CreateTask(context.TODO(), &manager.CreateTaskReq{
						Secret:       business.Secret,
						TestType:     businessTest.TestType,
						Desc:         fmt.Sprintf("%s(%s)[%0.2f] %s", business.Secret, businessTest.TestType, business.CurrentEdition, time.Now().Format(constant.DATE_TYPE_01)),
						FileName:     tmpZipName,
						FileLocation: tmpFileLocation,
					})
					if err != nil {
						log.Printf("call Manager CreateTask failed, err: %v", err)
						removeAllFile(constant.TEMP_FILE_LOCATION + business.Secret + string(os.PathSeparator) + fileInfo.Name)
						continue
					}
					removeAllFile(constant.TEMP_FILE_LOCATION + business.Secret + string(os.PathSeparator) + fileInfo.Name)
					_, err = sp.UpdateBusinessEdition(context.TODO(), &supporter.UpdateBusinessEditionReq{
						Edition: business.CurrentEdition,
						Secret:  business.Secret,
					})
					if err != nil {
						log.Printf("call supporter UpdateBusinessEdition failed, err: %v", err)
						continue
					}
					log.Printf("create task [%d] success", createTaskResp.Id)
				case constant.TEST_OPTION_MULTI:
				default:
				}
			} else {
				log.Printf("[%s][%s] current edition [%0.2f][%0.2f] is the newest", business.Secret, business.ApkName, business.CurrentEdition, fileInfo.Edition)
			}
			time.Sleep(time.Second * 5)
		}
		time.Sleep(time.Second * time.Duration(business.PullFrequency))
	}
}

// 获取文件最新版本信息
func callFileInfo(url string) (*bean.FileInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("call [%s] failed, err: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[%s] read body failed, err: %v", url, err)
		return nil, err
	}
	var fileUrlResp *bean.FileUrlResp
	err = json.Unmarshal(body, &fileUrlResp)
	if err != nil {
		log.Printf("[%s] json unmarshal failed, err: %v", body, err)
		return nil, err
	}
	if fileUrlResp.Code != 201 {
		log.Printf("return code: %d, return msg: %s", fileUrlResp.Code, fileUrlResp.Msg)
		return nil, errors.New(fmt.Sprintf("code: %d, return msg: %s", fileUrlResp.Code, fileUrlResp.Msg))
	}
	return &fileUrlResp.Data, nil
}

// 获取文件
func pullFile(fileUrl, secret, fileName string) (string, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		log.Printf("call [%s] failed, err: %v", fileUrl, err)
		return "", err
	}
	defer resp.Body.Close()
	tempFile := constant.TEMP_FILE_LOCATION + secret + string(os.PathSeparator) + fileName
	file, err := os.OpenFile(tempFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("create temp failed, err: %v", err)
		return "", err
	}
	defer file.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[%s] read body failed, err: %v", fileUrl, err)
		return "", err
	}
	_, err = file.WriteString(string(body))
	if err != nil {
		log.Printf("write temp file failed, err: %v", err)
		return "", err
	}
	return tempFile, nil
}

// 单文件处理
func singleFileHandle(apkFile, testType string) (string, string, error) {
	zipFile, err := os.Create(apkFile + ".zip")
	if err != nil {
		log.Printf("create zip file failed, err: %v", err)
		return "", "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileList := []string{
		apkFile, constant.TEST_SCRIPTS_LOCATION + testType + ".py",
	}

	for _, fileLocation := range fileList {
		file, err := os.Open(fileLocation)
		if err != nil {
			log.Printf("open file [%s] failed, err: %v", apkFile, err)
			return "", "", err
		}

		fileInfo, err := file.Stat()
		if err != nil {
			log.Printf("get file [%s] info failed, err: %v", apkFile, err)
			file.Close()
			return "", "", err
		}
		// 写入文件的头信息
		fileHead, err := zip.FileInfoHeader(fileInfo)

		fileWriter, err := zipWriter.CreateHeader(fileHead)
		if err != nil {
			log.Printf("writer file [%s] head failed, err: %v", apkFile, err)
			file.Close()
			return "", "", err
		}
		// 写入文件内容
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			log.Printf("writer file [%s] content failed, err: %v", apkFile, err)
			file.Close()
			return "", "", err
		}
		file.Close()
	}
	_, fileName := filepath.Split(zipFile.Name())
	return fileName, apkFile + ".zip", nil
}

// 删除文件
func removeAllFile(location string) {
	err := os.Remove(location)
	if err != nil {
		log.Printf("remove %s failed, err: %v", location, err)
		return
	}
	err = os.Remove(location + ".zip")
	if err != nil {
		log.Printf("remove %s failed, err: %v", location+".zip", err)
		return
	}
	log.Printf("remove %s success", location)
}
