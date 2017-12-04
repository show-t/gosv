package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB = dbConnect()

func dbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", dbResource)

	if err != nil {
		panic(err.Error())
	}

	return db
}
