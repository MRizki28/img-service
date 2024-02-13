package Config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_img_service"))
	if err != nil {
		panic(err)
	}

	DB = database
}