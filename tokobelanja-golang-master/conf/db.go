package conf

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	userName := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	hostIP := os.Getenv("HOST")
	portServer := os.Getenv("PORTDB")
	dbName := os.Getenv("DATABASE")

	// read db
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, hostIP, portServer, dbName)

	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}

	return db
}
