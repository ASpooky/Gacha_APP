package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//parseTime=trueをつけることでtime.Time型をsqlでエラーなく読み込める.
	dsn := fmt.Sprintf("user:password@tcp(127.0.0.1:3306)/catech_db?parseTime=true")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + ";database can't connect")
	}
	fmt.Println("Connected!!")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
