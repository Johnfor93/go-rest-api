package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi?parseTime=true"))

	if err!=nil{
		panic(err)
	}

	database.AutoMigrate(&Contact{})

	DB = database
}