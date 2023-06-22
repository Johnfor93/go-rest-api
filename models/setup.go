package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"os"
	"fmt"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

  	if err != nil {
    	fmt.Printf("Error loading .env file")
		return
  	}

	databaseURI := os.Getenv("database_uri")

	database, err := gorm.Open(mysql.Open(databaseURI))

	if err!=nil{
		panic(err)
	}

	database.AutoMigrate(&Contact{})

	DB = database
}