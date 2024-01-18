package model

import "time"

type Character struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" gorm:"unique"`
}

type Possession struct {
	ID          uint      `gorm:"primary_key"`
	UserID      uint      `json:"userID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterID uint      `json:"characterID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Emission struct {
	ID         uint `gorm:"primary_key"`
	CaracterID uint `json:"characterID" gorm:"unique,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Rarity     int  `json:"rarity" gorm:"not null"`
}
