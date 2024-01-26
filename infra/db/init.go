package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var howellDb *gorm.DB

// ...
func Init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:345612@tcp(127.0.0.1:3306)/howell?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	howellDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return howellDb
}
