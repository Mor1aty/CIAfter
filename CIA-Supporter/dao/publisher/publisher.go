package publisher

import (
	"log"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/19 19:23
 * @Description TODO
 * 任务发布者服务 dao
 */

// 获取所有业务
func FindAllBusiness(isStop uint8) ([]*bean.Business, error) {
	sqlStr := "SELECT secret, secret_value, apk_name, pull_frequency, current_edition, file_url, is_stop, update_time FROM business"
	switch isStop {
	case 0:
		sqlStr += " WHERE is_stop = 0"
	case 1:
		sqlStr += " WHERE is_stop = 1"
	}
	var businesses []*bean.Business
	err := dao.DB.Select(&businesses, sqlStr)
	if err != nil {
		log.Printf("get all business failed, err: %v", err)
		return nil, err
	}
	return businesses, nil
}

// 获取所有测试类型
func FindBusinessTestBySecret(secret string) ([]*bean.BusinessTest, error) {
	sqlStr := "SELECT id, secret, test_type, test_option, update_time FROM business_test WHERE secret = ?"
	var businessTests []*bean.BusinessTest
	err := dao.DB.Select(&businessTests, sqlStr, secret)
	if err != nil {
		log.Printf("get business_test by secret failed, err: %v", err)
		return nil, err
	}
	return businessTests, nil
}

// 更新业务版本
func UpdateBusinessEdition(edition, secret string) error {
	sqlStr := "UPDATE business SET current_edition = ? WHERE secret = ?"
	_, err := dao.DB.Exec(sqlStr, edition, secret)
	if err != nil {
		log.Printf("update business (current_edition) failed, err: %v", err)
		return err
	}
	return nil
}
