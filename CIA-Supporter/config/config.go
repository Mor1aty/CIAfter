package config

type SupporterConfig struct {
	MysqlConfig `ini:"mysql"`
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
