package config

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/12 19:58
 * @Description TODO
 * supporter 配置
 */
type SupporterConfig struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
	TaskConfig  `ini:"task"`
}

type MysqlConfig struct {
	Address     string `ini:"address"`
	Database    string `ini:"database"`
	Params      string `ini:"params"`
	Username    string `ini:"username"`
	Password    string `ini:"password"`
	MaxOpenConn int    `ini:"maxOpenConn"`
	MaxIdleConn int    `ini:"maxIdleConn"`
}

type RedisConfig struct {
	Address  string `ini:"address"`
	Password string `ini:"password"`
}

type TaskConfig struct {
	Key     string `ini:"key"`
	MaxTask int8   `ini:"max_task"`
}
