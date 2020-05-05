package service

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	supporter "moriaty.com/cia/cia-common/proto/supporter/executor"
	"moriaty.com/cia/cia-executor/bean"
	"moriaty.com/cia/cia-executor/config"
	"moriaty.com/cia/cia-executor/handle"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/14 11:24
 * @Description TODO
 * 任务执行者服务 service
 */

var (
	ses           supporter.SupporterExecutorService
	currentPhones = make([]*bean.Phone, 0)
	clientCfg     *config.ClientConfig
)

// 初始化 service
func Init(server micro.Service, cfg *config.ClientConfig) error {
	ses = supporter.NewSupporterExecutorService("supporter", server.Client())
	clientCfg = cfg
	return nil
}

// 消费任务
func ConsumeTask() {
	log.Printf("consume task start")
	defer log.Printf("consume task end")
	time.Sleep(time.Second * 2)
	for {
		findTaskKeyResp, err := ses.FindTaskKey(context.TODO(), &supporter.FindTaskKeyReq{Secret: clientCfg.Secret})
		if err != nil {
			log.Printf("call supporter FindTaskKey failed, err: %v", err)
			return
		}
		if findTaskKeyResp.Key == "" {
			log.Printf("no task to consume")
		} else {
			findTaskByKeyResp, err := ses.FindTaskByKey(context.TODO(), &supporter.FindTaskByKeyReq{Key: findTaskKeyResp.Key})
			if err != nil {
				log.Printf("call supporter FindTaskByKey failed, err: %v", err)
				return
			}
			if findTaskByKeyResp.Task == nil {
				log.Printf("get task by key [%s] not found", findTaskKeyResp.Key)
			} else {
				log.Printf("consume task [%d] start", findTaskByKeyResp.Task.TaskId)
				var currentPhone *bean.Phone = nil
				for {
					for _, phone := range currentPhones {
						if phone.IsOccupy == false {
							currentPhone = phone
							break
						}
					}
					if currentPhone == nil {
						log.Printf("phones are all occupied")
					} else {
						break
					}
					time.Sleep(time.Second * 10)
				}

				_, err = ses.UpdateTaskStartById(context.TODO(), &supporter.UpdateTaskStartByIdReq{
					Id:     findTaskByKeyResp.Task.TaskId,
					Client: currentPhone.Ip,
				})
				if err != nil {
					log.Printf("call supporter UpdateTaskStartById failed, err: %v", err)
					return
				}

				apkFile := findTaskByKeyResp.Task.TaskFile
				dest, file := filepath.Split(apkFile)
				testFile := dest + string(os.PathSeparator) + strings.Split(file, ".")[0] + string(os.PathSeparator) + findTaskByKeyResp.Task.TaskType + ".py"
				result := constant.RESULT_FILE_LOCATION + string(findTaskByKeyResp.Task.TaskId)

				err = handle.Exec(apkFile, testFile, currentPhone.Type, currentPhone.Edition, currentPhone.Id, file, result)
				isSuccess := false
				resultDesc := ""
				if err != nil {
					log.Printf("exec failed")
					resultDesc = "执行失败，错误：" + err.Error()
				} else {
					log.Printf("exec success")
					isSuccess = true
					resultDesc = "执行成功"
				}
				_, err := ses.UpdateTaskEndById(context.TODO(), &supporter.UpdateTaskEndByIdReq{
					Id:                  findTaskByKeyResp.Task.TaskId,
					IsSuccess:           isSuccess,
					ResultDesc:          resultDesc,
					ResultLocation:      apkFile,
					ResultImageLocation: result,
				})
				if err != nil {
					log.Printf("call supporter UpdateTaskEndById failed, err: %v", err)
					return
				}

				log.Printf("consume task [%d] end", findTaskByKeyResp.Task.TaskId)
			}
			err = deleteTask(findTaskKeyResp.Key)
			if err != nil {
				log.Printf("call supporter DeleteTaskByKey failed, err: %v", err)
				return
			}
		}
		time.Sleep(time.Second * 5)
	}
}

// 删除任务
func deleteTask(key string) error {
	_, err := ses.DeleteTaskByKey(context.TODO(), &supporter.DeleteTaskByKeyReq{Key: key})
	if err != nil {
		return err
	}
	return nil
}

// 初始化当前设备的手机信息
func InitPhone() error {
	ip, err := getIP()
	if err != nil {
		log.Printf("get ip failed, err: %v", err)
		return err
	}

	findClientByIpResp, err := ses.FindClientByIp(context.TODO(), &supporter.FindClientByIpReq{Ip: ip})
	if err != nil {
		log.Printf("call supporter FindClientByIp failed, err: %v", err)
		return err
	}

	if findClientByIpResp.Client == nil {
		log.Printf("client ip [%s] not found", ip)
		return errors.New("current client is not registered")
	}

	cmd := exec.Command("adb", "devices")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("adb std out failed, err: %v", err)
		return err
	}

	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		log.Printf("adb exec out failed, err: %v", err)
		return err
	}

	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Printf("adb out read failed, err: %v", err)
		return err
	}

	phones := make([]*supporter.Phone, 0)

	outSplit := strings.Split(string(out), "\r\n")
	for _, s := range outSplit {
		temp := strings.Split(s, "\t")
		if len(temp) == 2 {
			phones = append(phones, &supporter.Phone{
				Id:           temp[0],
				Client:       findClientByIpResp.Client.Ip,
				TestType:     findClientByIpResp.Client.TestType,
				Secret:       findClientByIpResp.Client.Secret,
				PhoneType:    clientCfg.PhoneType,
				PhoneEdition: clientCfg.PhoneEdition,
			})
			currentPhones = append(currentPhones, &bean.Phone{
				Id:      temp[0],
				Type:    clientCfg.PhoneType,
				Edition: clientCfg.PhoneEdition,
				Ip:      ip,
			})
		}
	}
	if len(phones) == 0 {
		log.Printf("no phone needs to be registered")
		return nil
	}

	_, err = ses.InsertPhone(context.TODO(), &supporter.InsertPhoneReq{Phones: phones})
	if err != nil {
		log.Printf("call supporter InsertPhone failed, err: %v", err)
		return err
	}
	return nil
}

// 获取本机 ip
func getIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
