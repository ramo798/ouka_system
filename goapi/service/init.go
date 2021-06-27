package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "my_database"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	var err error
	db, err := gorm.Open(DBMS, CONNECT)
	db.SingularTable(true)
	if err != nil {
		panic(err.Error())
	}
	return db
}
