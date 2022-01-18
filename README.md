# 基于 golang http 标准库的http服务最佳实践

##  常用命令

```shell
# 初始化项目 go mod
$ go mod init GoHttpServerBestPractice

```

## 技术选型

| 名称                    | 链接                                               | 安装方式                                      | star  | 说明                    |
| ----------------------- | -------------------------------------------------- | --------------------------------------------- | ----- | ----------------------- |
| viper                   | [链接](https://github.com/spf13/viper)             | go get github.com/spf13/viper                 | 16.1k | golang 配置文件解决方案 |
| go-playground/validator | [链接](https://github.com/go-playground/validator) | go get github.com/go-playground/validator/v10 | 8.9k  | 表单验证                |
| Mysql相关库             | [链接](https://github.com/go-sql-driver/mysql)     | go get github.com/go-sql-driver/mysql         | 11.6K | mysql驱动               |
|                         | [链接](https://github.com/jmoiron/sqlx)            | go get github.com/jmoiron/sqlx                | 10.9K | 基于标准库 sql 的扩展   |

