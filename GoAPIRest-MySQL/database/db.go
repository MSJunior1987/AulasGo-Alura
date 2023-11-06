package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	erro error
)

func ConectarComBancoDeDados() {
	stringDeConexao := "root:root@tcp(127.0.0.1:3306)/root?charset=utf8mb4&parseTime=True&loc=Local"
	DB, erro = gorm.Open(mysql.Open(stringDeConexao), &gorm.Config{})
}
