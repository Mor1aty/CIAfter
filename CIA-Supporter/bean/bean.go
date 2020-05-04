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
	Secret       string    `db:"secret"`
	CreateTime   time.Time `db:"create_time"`
}

// 手机表（mysql）
type Phone struct {
	Id           string    `db:"id"`
	Client       string    `db:"client"`
	TestType     string    `db:"test_type"`
	Secret       string    `db:"secret"`
	PhoneType    string    `db:"phone_type"`
	PhoneEdition string    `db:"phone_edition"`
	UpdateTime   time.Time `db:"update_time"`
}

// 客户端表（mysql）
type Client struct {
	Ip         string    `db:"ip"`
	TestType   string    `db:"test_type"`
	Secret     string    `db:"secret"`
	UpdateTime time.Time `db:"update_time"`
}

// 业务表（business）
type Business struct {
	Secret         string    `db:"secret"`
	SecretValue    string    `db:"secret_value"`
	ApkName        string    `db:"apk_name"`
	PullFrequency  int64     `db:"pull_frequency"`
	CurrentEdition float32   `db:"current_edition"`
	FileUrl        string    `db:"file_url"`
	IsStop         uint8     `db:"is_stop"`
	UpdateTime     time.Time `db:"update_time"`
}

// 业务测试表（business_test）
type BusinessTest struct {
	Id         int64     `db:"id"`
	Secret     string    `db:"secret"`
	TestType   string    `db:"test_type"`
	TestOption int64     `db:"test_option"`
	UpdateTime time.Time `db:"update_time"`
}

// 任务表（task）
type Task struct {
	Id         int64     `db:"id"`
	Secret     string    `db:"secret"`
	TestType   string    `db:"test_type"`
	Client     string    `db:"client"`
	Desc       string    `db:"desc"`
	File       int64     `db:"file"`
	Status     int32     `db:"status"`
	Result     int32     `db:"result"`
	ResultDesc string    `db:"result_desc"`
	CreateTime time.Time `db:"create_time"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
}

// 任务（redis）
type TaskRedis struct {
	TaskId   int64  `json:"task_id"`
	TaskType string `json:"task_type"`
	TaskFile string `json:"task_file"`
}
