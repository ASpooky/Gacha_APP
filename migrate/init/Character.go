package main

import (
	"fmt"
	"log"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/model"
)

type CharaInfo struct {
	Name   string
	Rarity int
}

type CharacterRequest struct {
	Name string
}

func main() {
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)

	characters := []CharaInfo{
		{Name: "a", Rarity: 1},
		{Name: "b", Rarity: 1},
		{Name: "c", Rarity: 1},
		{Name: "d", Rarity: 2},
		{Name: "e", Rarity: 2},
		{Name: "f", Rarity: 2},
		{Name: "g", Rarity: 3},
		{Name: "h", Rarity: 3},
		{Name: "i", Rarity: 3},
		{Name: "j", Rarity: 4},
		{Name: "k", Rarity: 4},
		{Name: "l", Rarity: 4},
		{Name: "m", Rarity: 5},
		{Name: "n", Rarity: 5},
		{Name: "o", Rarity: 5},
	}

	//create characters
	for _, chara := range characters {
		if err := dbCon.Create(&model.Character{Name: chara.Name}).Error; err != nil {
			log.Fatal("err:", err.Error())
		}
	}

	//create emmisions
	for _, chara := range characters {
		var target model.Character
		//find character id
		if err := dbCon.Where("name=?", chara.Name).First(&target).Error; err != nil {
			log.Fatal("err:", err.Error())
		}
		if err := dbCon.Create(&model.Emission{CaracterID: target.ID, Rarity: chara.Rarity}).Error; err != nil {
			log.Fatal("err:", err.Error())
		}
	}

	fmt.Println("Success characters table intialized!")
	return
}
