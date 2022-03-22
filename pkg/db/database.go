package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() (*gorm.DB, error) {

	dsn := "root:secret@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
