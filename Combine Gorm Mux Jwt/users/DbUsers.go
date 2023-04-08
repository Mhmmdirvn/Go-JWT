package users

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDBUsers() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_users"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}