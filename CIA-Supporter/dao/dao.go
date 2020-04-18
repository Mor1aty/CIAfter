package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"moriaty.com/cia/cia-supporter/config"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/7 11:08
 * @Description TODO
 * dao 初始化
 */

var (
	DB  *sqlx.DB
	RDB *redis.Client
)

// 初始化数据库信息
func InitMysql(mysqlCfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Address, mysqlCfg.Database, mysqlCfg.Params)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		// dsn 格式不正确
		return err
	}
	DB.SetMaxOpenConns(mysqlCfg.MaxOpenConn)
	DB.SetMaxIdleConns(mysqlCfg.MaxIdleConn)
	return nil
}

// 初始化 Redis 信息
func InitRedis(redisCfg *config.RedisConfig) (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
	})
	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
