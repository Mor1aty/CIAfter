package executor

import (
	"encoding/json"
	"fmt"
	"log"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao"
	"strconv"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/14 11:31
 * @Description TODO
 * 任务执行者服务 dao
 */

// 获取任务 id
func PopId(idsKey string) (int64, error) {
	idStr, err := dao.RDB.RPop(idsKey).Result()
	if err != nil {
		log.Printf("get task id failed, err: %v\n", err)
		return 0, err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("task id [%s] error, err: %v\n", idStr, err)
		return 0, err
	}
	return id, nil
}

// 根据 key 获取任务
func FindTaskByKey(key string) (*bean.Task, error) {
	newTask, err := dao.RDB.Get(key).Result()
	if err != nil {
		log.Printf("get task [%s] failed, err: %v\n", key, err)
		return nil, err
	}
	var task bean.Task
	err = json.Unmarshal([]byte(newTask), &task)
	if err != nil {
		log.Printf("json unmarshal task [%s] failed, err: %v\n", key, err)
		return nil, err
	}
	return &task, nil
}

// 根据 key 删除任务
func DeleteTaskByKey(key string) error {
	err := dao.RDB.Del(key).Err()
	if err != nil {
		log.Printf("delete task [%s] failed, err: %v\n", key, err)
		return err
	}
	return nil
}

// 推送测试任务
func PushTestTask() {
	ids := []int64{10010, 10011, 10012, 10013}

	for _, id := range ids {
		dao.RDB.LPush("ci_task_ids", id)
		task := bean.Task{
			TaskId:   id,
			TaskType: "AAA",
			TaskFile: "xx.xx",
		}
		taskJson, _ := json.Marshal(task)
		dao.RDB.Set(fmt.Sprintf("ci_task_%d", id), taskJson, 0).Err()
	}
}

// 根据 id 获取手机信息
func FindPhoneById(id string) (*bean.Phone, error) {
	sqlStr := "SELECT id, client, test_type, secret, update_time FROM phone WHERE id = ?"
	var phone bean.Phone
	err := dao.DB.Get(&phone, sqlStr, id)
	if err != nil {
		log.Printf("get phone by id failed, err: %v\n", err)
		return nil, err
	}
	return &phone, nil
}

// 存入手机信息
func InsertPhone(phone *bean.Phone) error {
	log.Printf("%#v\n", phone)
	sqlStr := "INSERT INTO phone(id, client, test_type, secret, update_time) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.DB.Exec(sqlStr, phone.Id, phone.Client, phone.TestType, phone.Secret, phone.UpdateTime)
	if err != nil {
		log.Printf("insert phone failed, err: %v\n", err)
		return err
	}
	return nil
}

// 更新手机 client
func UpdatePhoneClient(id, client string) error {
	sqlStr := "UPDATE phone SET client = ? WHERE id = ?"
	_, err := dao.DB.Exec(sqlStr, client, id)
	if err != nil {
		log.Printf("update phone (client) failed, err: %v\n", err)
		return err
	}
	return nil
}

// 根据 ip 获取客户端
func FindClientByIp(ip string) (*bean.Client, error) {
	sqlStr := "SELECT ip, test_type, secret, update_time FROM client WHERE ip = ?"
	var client bean.Client
	err := dao.DB.Get(&client, sqlStr, ip)
	if err != nil {
		log.Printf("get client by ip failed, err: %v\n", err)
		return nil, err
	}
	return &client, nil
}
