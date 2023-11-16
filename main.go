package main

import (
	"github.com/ASpooky/ca_tech_dojo/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
