package bean

import "time"

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 15:18
 * @Description TODO
 * 数据结构体
 */

// 文件表（mysql）
type File struct {
	Id           int64     `db:"id"`
	FileName     string    `db:"file_name"`
	FileLocation string    `db:"file_location"`
	Task         int64     `db:"task"`
	Secret       string    `db:"secret"`
	CreateTime   time.Time `db:"create_time"`
}

// 手机表（mysql）
type Phone struct {
	Id         string    `db:"id"`
	Client     string    `db:"client"`
	TestType   string    `db:"test_type"`
	Secret     string    `db:"secret"`
	UpdateTime time.Time `db:"update_time"`
}

// 客户端表（client）
type Client struct {
	Ip         string    `db:"ip"`
	TestType   string    `db:"test_type"`
	Secret     string    `db:"secret"`
	UpdateTime time.Time `db:"update_time"`
}

// 任务（redis）
type Task struct {
	TaskId   int64  `json:"task_id"`
	TaskType string `json:"task_type"`
	TaskFile string `json:"task_file"`
}
