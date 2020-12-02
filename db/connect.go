package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser = "root"
	dbPass = "dongnv2804"
	dbName = "quizdb"
)

// Dbconn : connect to database
func Dbconn() (db *gorm.DB) {
	dns := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
