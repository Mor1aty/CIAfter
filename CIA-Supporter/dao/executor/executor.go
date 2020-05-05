package executor

import (
	"encoding/json"
	"log"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao"
	"strconv"
	"time"
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
		log.Printf("get task id failed, err: %v", err)
		return 0, err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("task id [%s] error, err: %v", idStr, err)
		return 0, err
	}
	return id, nil
}

// 根据 key 获取任务
func FindTaskByKey(key string) (*bean.TaskRedis, error) {
	newTask, err := dao.RDB.Get(key).Result()
	if err != nil {
		log.Printf("get task [%s] failed, err: %v", key, err)
		return nil, err
	}
	var task bean.TaskRedis
	err = json.Unmarshal([]byte(newTask), &task)
	if err != nil {
		log.Printf("json unmarshal task [%s] failed, err: %v", key, err)
		return nil, err
	}
	return &task, nil
}

// 根据 key 删除任务
func DeleteTaskByKey(key string) error {
	err := dao.RDB.Del(key).Err()
	if err != nil {
		log.Printf("delete task [%s] failed, err: %v", key, err)
		return err
	}
	return nil
}

// 根据 id 获取手机信息
func FindPhoneById(id string) (*bean.Phone, error) {
	sqlStr := "SELECT * FROM phone WHERE id = ?"
	var phone bean.Phone
	err := dao.DB.Get(&phone, sqlStr, id)
	if err != nil {
		log.Printf("get phone by id failed, err: %v", err)
		return nil, err
	}
	return &phone, nil
}

// 存入手机信息
func InsertPhone(phone *bean.Phone) error {
	sqlStr := "INSERT INTO phone(id, client, test_type, secret, phone_type, phone_edition, update_time) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := dao.DB.Exec(sqlStr, phone.Id, phone.Client, phone.TestType, phone.Secret, phone.PhoneType, phone.PhoneEdition, phone.UpdateTime)
	if err != nil {
		log.Printf("insert phone failed, err: %v", err)
		return err
	}
	return nil
}

// 更新手机 client
func UpdatePhoneClient(id, client string) error {
	sqlStr := "UPDATE phone SET client = ? WHERE id = ?"
	_, err := dao.DB.Exec(sqlStr, client, id)
	if err != nil {
		log.Printf("update phone (client) failed, err: %v", err)
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
		log.Printf("get client by ip failed, err: %v", err)
		return nil, err
	}
	return &client, nil
}

// 更新任务开始信息
func UpdateTaskStartById(id int64, client string) error {
	sqlStr := "UPDATE task SET client = ?, status = 1, start_time = ?  WHERE id = ?"
	_, err := dao.DB.Exec(sqlStr, client, time.Now(), id)
	if err != nil {
		log.Printf("update task (%d) start failed, err: %v", id, err)
		return err
	}
	return nil
}

// 更新任务结束信息
func UpdateTaskEndById(id int64, result int, resultDesc, resultLocation, resultImageLocation string) error {
	sqlStr := "UPDATE task SET status = 2, result = ?, result_desc = ?, result_location = ?, result_image_location = ?, end_time = ?  WHERE id = ?"
	_, err := dao.DB.Exec(sqlStr, result, resultDesc, resultLocation, resultImageLocation, time.Now(), id)
	if err != nil {
		log.Printf("update task (%d) end failed, err: %v", id, err)
		return err
	}
	return nil
}
