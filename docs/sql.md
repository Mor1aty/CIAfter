# 数据库（ci_after）

```mysql
CREATE DATABASE `ci_after` CHARSET utf8;
```

## 文件表（file）

|    字段名     |      类型       |             约束             |   描述   |
| :-----------: | :-------------: | :--------------------------: | :------: |
|      id       | bigint unsigned |             主键             | 自增主键 |
|   file_name   |   varchar(50)   |             非空             |  文件名  |
| file_location |  varchar(200)   |             非空             | 文件位置 |
|     task      | bigint unsigned |   外键非空，`任务表`(`id`)   | 任务 id  |
|    secret     |   varchar(50)   | 外键非空，`业务表`(`secret`) | 权限 id  |
|  create_time  |    datetime     |             非空             | 创建时间 |

```mysql
CREATE TABLE `file`(
	`id` BIGINT UNSIGNED AUTO_INCREMENT,
    `file_name` VARCHAR(50) NOT NULL,
    `file_location` VARCHAR(200) NOT NULL,
    `task` BIGINT UNSIGNED NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `create_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;
```



## 手机表（phone）

|   字段名    |    类型     |             约束             |    描述     |
| :---------: | :---------: | :--------------------------: | :---------: |
|     id      | varchar(30) |             主键             | 手机设备 id |
|   client    | varchar(20) |  外键非空，`客户端表`(`ip`)  |   client    |
|  test_type  | varchar(50) |             非空             |  测试类型   |
|   secret    | varchar(50) | 外键非空，`业务表`(`secret`) |   权限 id   |
| update_time |  datetime   |             非空             |  更新时间   |

```mysql
CREATE TABLE `phone`(
	`id` VARCHAR(30) NOT NULL,
    `client` VARCHAR(20) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;
```



## 客户端表（client）

|   字段名    |    类型     |             约束             |   描述   |
| :---------: | :---------: | :--------------------------: | :------: |
|     ip      | varchar(20) |             主键             |    ip    |
|  test_type  | varchar(50) |             非空             | 测试类型 |
|   secret    | varchar(50) | 外键非空，`业务表`(`secret`) | 权限 id  |
| update_time |  datetime   |             非空             | 更新时间 |

```mysql
CREATE TABLE `client`(
	`ip` VARCHAR(20) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`ip`)
)CHARSET utf8 ENGINE INNODB;
```



## 业务表（business）

|     字段名      |       类型       | 约束 |              描述              |
| :-------------: | :--------------: | :--: | :----------------------------: |
|     secret      |   varchar(50)    | 主键 |            权限 id             |
|  secret_value   |   varchar(100)   | 非空 |             权限值             |
|    apk_name     |   varchar(100)   | 非空 |           apk 文件名           |
| pull_frequency  |   int unsigned   | 非空 |     拉取频率，单位：分/次      |
| current_edition |   varchar(100)   | 非空 |            当前版本            |
|    file_url     |   varchar(200)   | 非空 |            文件位置            |
|     is_stop     | tinyint unsigned | 非空 | 停用标记，0 为未停用，1 为停用 |
|   update_time   |     datetime     | 非空 |            更新时间            |

```mysql
CREATE TABLE `business`(
	`secret` VARCHAR(50) NOT NULL,
    `secret_value` VARCHAR(100) NOT NULL,
    `apk_name` VARCHAR(100) NOT NULL,
    `pull_frequency` INT UNSIGNED NOT NULL,
    `current_edition` VARCHAR(100) NOT NULL,
    `file_url` VARCHAR(200) NOT NULL,
    `is_stop` TINYINT UNSIGNED NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`secret`)
)CHARSET utf8 ENGINE INNODB;
```



## 业务测试表（business_test）

|   字段名    |      类型       |             约束             |                      描述                       |
| :---------: | :-------------: | :--------------------------: | :---------------------------------------------: |
|     id      | bigint unsigned |           自增主键           |                       id                        |
|   secret    |   varchar(50)   | 外键非空，`业务表`(`secret`) |                     权限 id                     |
|  test_type  |   varchar(50)   |             非空             |                    测试类型                     |
| test_option |  int unsigned   |             非空             | 测试操作，0 为单测试文件，1 为多测试文件处理... |
| update_time |    datetime     |             非空             |                    更新时间                     |

```mysql
CREATE TABLE `business_test`(
	`id` BIGINT UNSIGNED AUTO_INCREMENT,
    `secret` VARCHAR(50) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `test_option` INT UNSIGNED NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;
```





## 执行 sql

```mysql
CREATE DATABASE `ci_after` CHARSET utf8;
USE `ci_after`;

DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`(
	`id` BIGINT UNSIGNED AUTO_INCREMENT,
    `file_name` VARCHAR(50) NOT NULL,
    `file_location` VARCHAR(200) NOT NULL,
    `task` BIGINT UNSIGNED NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `create_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;

DROP TABLE IF EXISTS `phone`;
CREATE TABLE `phone`(
	`id` VARCHAR(30) NOT NULL,
    `client` VARCHAR(20) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;

DROP TABLE IF EXISTS `client`;
CREATE TABLE `client`(
	`ip` VARCHAR(20) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `secret` VARCHAR(50) NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`ip`)
)CHARSET utf8 ENGINE INNODB;

DROP TABLE IF EXISTS `business`;
CREATE TABLE `business`(
	`secret` VARCHAR(50) NOT NULL,
    `secret_value` VARCHAR(100) NOT NULL,
    `apk_name` VARCHAR(100) NOT NULL,
    `pull_frequency` INT UNSIGNED NOT NULL,
    `current_edition` VARCHAR(100) NOT NULL,
    `file_url` VARCHAR(200) NOT NULL,
    `is_stop` TINYINT UNSIGNED NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`secret`)
)CHARSET utf8 ENGINE INNODB;

DROP TABLE IF EXISTS `business_test`;
CREATE TABLE `business_test`(
	`id` BIGINT UNSIGNED AUTO_INCREMENT,
    `secret` VARCHAR(50) NOT NULL,
    `test_type` VARCHAR(50) NOT NULL,
    `test_option` INT UNSIGNED NOT NULL,
    `update_time` DATETIME NOT NULL,
    PRIMARY KEY(`id`)
)CHARSET utf8 ENGINE INNODB;
```

