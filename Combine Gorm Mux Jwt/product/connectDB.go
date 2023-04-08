package product

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBMarket() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_market"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Product{})
	return db, nil
}