package config

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/5/4 22:56
 * @Description TODO
 * executor 配置
 */
type ExecutorConfig struct {
	ClientConfig `ini:"client"`
}

type ClientConfig struct {
	Secret       string `ini:"secret"`
	PhoneType    string `ini:"phone_type"`
	PhoneEdition string `ini:"phone_edition"`
}
