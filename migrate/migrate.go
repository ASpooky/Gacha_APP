package main

import (
	"fmt"
	"log"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/model"
)

func main() {
	dbCon := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbCon)
	if err := dbCon.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("Err:", err)
	}
	if err := dbCon.AutoMigrate(&model.Character{}); err != nil {
		log.Fatalln("Err:", err)
	}
	if err := dbCon.AutoMigrate(&model.Possession{}); err != nil {
		log.Fatalln("Err:", err)
	}
	if err := dbCon.AutoMigrate(&model.Emission{}); err != nil {
		log.Fatalln("Err:", err)
	}
}
