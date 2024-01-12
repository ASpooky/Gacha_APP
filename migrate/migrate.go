package main

import (
	"fmt"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/model"
)

func main() {
	dbCon := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(&model.User{})
	dbCon.AutoMigrate(&model.Character{})
	dbCon.AutoMigrate(&model.Possession{})
	dbCon.AutoMigrate(&model.Emission{})
}
