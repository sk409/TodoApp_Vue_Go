package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "host=database user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&User{}, &Todo{})
}
