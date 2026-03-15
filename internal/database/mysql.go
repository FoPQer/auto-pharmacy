package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Mysql struct {
	DB *gorm.DB
}

var MysqlDB *Mysql

func RegisterMysql() (*Mysql, error) {
	dsn := "root@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	bd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Mysql{}, err
	}
	return &Mysql{DB: bd}, nil
}
