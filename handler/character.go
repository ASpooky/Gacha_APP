package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ASpooky/ca_tech_dojo/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CharacterHandler struct {
	db *gorm.DB
}

func NewCharacterHandler(db *gorm.DB) *CharacterHandler {
	return &CharacterHandler{db}
}

type PossessionResponse struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
	Quantity        int    `json:"quantity"`
}

type GetCharacterListResponse struct {
	Characters []PossessionResponse `json:"characters"`
}

type Possession struct {
	model.Possession
	model.Character
}

func (ch *CharacterHandler) GetCharacterList(c echo.Context) error {
	user := model.User{}

	token := c.Request().Header.Get("X-Token")
	if err := ch.db.Where("token=?", token).First(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not find user")
	}

	var possessions []Possession

	//キャラクターテーブルとポゼッションテーブルを結合して抽出
	if err := ch.db.Table("possessions").
		Select("possessions.id, characters.name, possessions.user_id, possessions.character_id, possessions.created_at, possessions.updated_at").
		Joins("JOIN characters ON characters.id = possessions.character_id").
		Where("possessions.user_id=?", user.ID).
		Order("possessions.character_id ASC").
		Find(&possessions).
		Error; err != nil {
		log.Println("err:", err)
		return c.String(http.StatusInternalServerError, "Could not find possessions")
	}

	//listで出力するため
	//capacityの指定
	result := make([]PossessionResponse, 0, len(possessions))

	//指定された出力フォーマットに型変換
	current_CharacterID := uint(0)
	current_index := -1
	for _, p := range possessions {
		if p.CharacterID != current_CharacterID {
			result = append(result, PossessionResponse{
				UserCharacterID: strconv.FormatUint(uint64(p.Possession.ID), 10),
				CharacterID:     strconv.FormatUint(uint64(p.Possession.CharacterID), 10),
				Name:            p.Character.Name,
				Quantity:        1,
			})
			current_CharacterID = uint(p.CharacterID)
			current_index += 1
		} else {
			result[current_index].Quantity += 1
		}

	}

	//c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, GetCharacterListResponse{Characters: result})
}
