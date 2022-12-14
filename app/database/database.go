package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    int    `gorm:"column:id"`
	Uid   string `gorm:"column:user_id"`
	UName string `gorm:"column:user_name"`
	UPass string `gorm:"column:user_password"`
}

func DBConnect() *gorm.DB {
	info := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(info), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	return db
}
