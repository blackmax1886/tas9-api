package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	USERNAME := "localuser"
	PASSWORD := "localpass"
	HOST := "127.0.0.1"
	PORT := "3306"
	DATABASE := "localdb"

	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DATABASE + "?charset=utf8mb4" + "&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil, err
	}

	return db, nil
}
