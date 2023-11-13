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
	//AutoMigrateでTable名を指定する場合は??
}
