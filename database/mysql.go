package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

var MysqlDB *Mysql

func RegisterMysql() (*Mysql, error) {
	dsn := "root@tcp(127.0.0.1:3306)/auto_pharmacy?charset=utf8mb4&parseTime=True&loc=Local"
	bd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Mysql{}, err
	}
	return &Mysql{DB: bd}, nil
}
