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
```

