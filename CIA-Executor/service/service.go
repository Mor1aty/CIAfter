package service

import (
	"context"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	supporter "moriaty.com/cia/cia-common/proto/supporter/executor"
	"net"
	"os/exec"
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

var se supporter.SupporterExecutorService

// 初始化 service
func Init(server micro.Service) error {
	se = supporter.NewSupporterExecutorService("supporter", server.Client())
	return nil
}

// 消费任务
func ConsumeTask() {
	log.Printf("consume task start")
	defer log.Printf("consume task end")
	time.Sleep(time.Second * 2)
	for {
		findTaskKeyResp, err := se.FindTaskKey(context.TODO(), &supporter.FindTaskKeyReq{})
		if err != nil {
			log.Printf("call supporter FindTaskKey failed, err: %v", err)
			return
		}
		if findTaskKeyResp.Key == "" {
			log.Printf("no task to consume")
		} else {
			findTaskByKeyResp, err := se.FindTaskByKey(context.TODO(), &supporter.FindTaskByKeyReq{Key: findTaskKeyResp.Key})
			if err != nil {
				log.Printf("call supporter FindTaskByKey failed, err: %v", err)
				return
			}
			if findTaskByKeyResp.Task == nil {
				log.Printf("get task by key [%s] not found", findTaskKeyResp.Key)
			} else {
				log.Printf("consume task [%d] start", findTaskByKeyResp.Task.TaskId)
				time.Sleep(time.Second * 10)
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
	_, err := se.DeleteTaskByKey(context.TODO(), &supporter.DeleteTaskByKeyReq{Key: key})
	if err != nil {
		return err
	}
	return nil
}

// 推送任务
func PushTestTask() {
	se.PushTestTask(context.TODO(), &supporter.PushTestTaskReq{})
}

// 初始化当前设备的手机信息
func InitPhone() error {
	ip, err := getIP()
	if err != nil {
		log.Printf("get ip failed, err: %v", err)
		return err
	}

	findClientByIpResp, err := se.FindClientByIp(context.TODO(), &supporter.FindClientByIpReq{Ip: ip})
	if err != nil {
		log.Printf("call supporter FindClientByIp failed, err: %v", err)
		return err
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
				Id:       temp[0],
				Client:   findClientByIpResp.Client.Ip,
				TestType: findClientByIpResp.Client.TestType,
				Secret:   findClientByIpResp.Client.Secret,
			})
		}
	}
	if len(phones) == 0 {
		log.Printf("no phone needs to be registered")
		return nil
	}

	_, err = se.InsertPhone(context.TODO(), &supporter.InsertPhoneReq{Phones: phones})
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
