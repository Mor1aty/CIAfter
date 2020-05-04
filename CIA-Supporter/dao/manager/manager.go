package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao"
	"time"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/5/1 19:53
 * @Description TODO
 * 任务管理者服务 dao
 */

// 根据权限 id 获取任务
func FindTaskBySecret(secret string) ([]*bean.Task, error) {
	sqlStr := "SELECT * FROM task WHERE secret = ?"
	var tasks []*bean.Task
	err := dao.DB.Select(&tasks, sqlStr, secret)
	if err != nil {
		log.Printf("get task by secret [%s] failed, err: %v", secret, err)
		return nil, err
	}
	return tasks, nil
}

// 插入任务
func InsertTask(secret, testType, desc string, file int64, createTime time.Time) (int64, error) {
	sqlStr := "INSERT INTO task(secret, test_type, `desc`, file, status, create_time) VALUES(?, ?, ?, ?, 0, ?)"
	ret, err := dao.DB.Exec(sqlStr, secret, testType, desc, file, createTime)
	if err != nil {
		log.Printf("insert task failed, err: %v", err)
		return 0, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		log.Printf("get last insert Id failed, err: %v", err)
		return 0, err
	}
	return id, nil
}

// 推送任务
func PushTask(key, secret string, task *bean.TaskRedis) error {
	err := dao.RDB.LPush(fmt.Sprintf("%s_ids_%s", key, secret), task.TaskId).Err()
	if err != nil {
		log.Printf("push id[%d] into ci_task_ids failed, err: %v", task.TaskId, err)
		return err
	}
	taskJson, _ := json.Marshal(task)
	err = dao.RDB.Set(fmt.Sprintf("%s_%s_%d", key, secret, task.TaskId), taskJson, 0).Err()
	if err != nil {
		log.Printf("push task[%d] failed, err: %v", task.TaskId, err)
		return err
	}
	return nil
}
