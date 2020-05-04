package filecenter

import (
	"log"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 14:25
 * @Description TODO
 * 	文件中心服务 dao
 */

// 存入文件
func InsertFile(file *bean.File) (int64, error) {
	sqlStr := "INSERT INTO file(file_name, file_location, secret, create_time) VALUES (?, ?, ?, ?)"
	ret, err := dao.DB.Exec(sqlStr, file.FileName, file.FileLocation, file.Secret, file.CreateTime)
	if err != nil {
		log.Printf("insert file failed, err: %v", err)
		return 0, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		log.Printf("get last insert Id failed, err: %v", err)
		return 0, err
	}
	return id, nil
}

// 根据 id 获取文件
func FindFileById(id int64) (*bean.File, error) {
	sqlStr := "SELECT id, file_name, file_location, secret, create_time FROM file WHERE id = ?"
	var file bean.File
	err := dao.DB.Get(&file, sqlStr, id)
	if err != nil {
		log.Printf("get file by id failed, err: %v", err)
		return nil, err
	}
	return &file, nil
}

// 根据任务 id 和权限 id 获取文件
func FindFileByTaskAndSecret(task int64, secret string) ([]*bean.File, error) {
	sqlStr := "SELECT id, file_name, file_location, secret, create_time FROM file WHERE task = ? AND secret = ?"
	var files []*bean.File
	err := dao.DB.Select(&files, sqlStr, task, secret)
	if err != nil {
		log.Printf("get file by task and secret failed, err: %v", err)
		return nil, err
	}
	return files, nil
}
