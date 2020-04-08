package bean

import "time"

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 15:18
 * @Description TODO
 * 数据表结构体
 */

// 文件表
type File struct {
	Id           int64     `db:"id"`
	FileName     string    `db:"file_name"`
	FileLocation string    `db:"file_location"`
	Task         int64     `db:"task"`
	Secret       string    `db:"secret"`
	CreateTime   time.Time `db:"create_time"`
}
