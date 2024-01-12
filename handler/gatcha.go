package handler

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ASpooky/ca_tech_dojo/model"
	"github.com/ASpooky/ca_tech_dojo/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GatchaHandler struct {
	db *gorm.DB
}

func NewGatchaHandler(db *gorm.DB) *GatchaHandler {
	return &GatchaHandler{
		db: db,
	}
}

type GatchaPlayRequest struct {
	Times int
}

type CharacterResponse struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type GatchaPlayResponse struct {
	Results []CharacterResponse `json:"results"`
}

// 確率導出の上限・下限
var MAX = 101
var MIN = 1

func (gh *GatchaHandler) PlayGatcha(c echo.Context) error {
	var req GatchaPlayRequest

	if err := c.Bind(&req); err != nil {
		log.Println("err :", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not bind request struct!")
	}

	var results []CharacterResponse

	//header tokenからuser情報の取得はmiddlewareにしとくといいかも
	var user model.User

	token := c.Request().Header.Get("X-Token")
	if err := gh.db.Where("token=?", token).First(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not find user")
	}

	for i := 0; i < req.Times; i++ {
		//1~100のrandomな整数→レアリティ
		randInt := rand.Intn(MAX-MIN) + MIN
		rarity := utils.NumToRarity(randInt)
		//log.Println(rarity)

		var emissions []model.Emission

		//レアリティからキャラクターID
		if err := gh.db.Where("rarity=?", rarity).Find(&emissions).Error; err != nil {
			log.Println("err :", err)
			return c.String(http.StatusInternalServerError, "err: Could not find emission")
		}

		lenEmissions := len(emissions)
		randEmission := rand.Intn(lenEmissions)

		var character model.Character

		//キャラクターIDからキャラクター
		if err := gh.db.Where("id=?", emissions[randEmission].CaracterID).First(&character).Error; err != nil {
			log.Println("err :", err)
			return c.String(http.StatusInternalServerError, "err: Could not find character")
		}
		//log.Println(character)

		//userIDとcharacterIDからposessionから抽出
		var possession model.Possession
		if err := gh.db.Where("user_id=? and character_id=?", user.ID, character.ID).First(&possession).Error; err != nil {
			//Possessionsの作成
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("new create")
				newPossession := model.Possession{
					UserID:      user.ID,
					CharacterID: character.ID,
					Quantity:    1,
				}
				gh.db.Create(&newPossession)
			} else {
				log.Println("err :", err)
				return c.String(http.StatusInternalServerError, "err: Could not find or create possessions")
			}
		} else {
			//Possessionsの更新
			newQuantity := possession.Quantity + 1
			possession.Quantity = newQuantity
			gh.db.Save(&possession)
		}

		results = append(results, CharacterResponse{CharacterID: strconv.FormatUint(uint64(character.ID), 10), Name: character.Name})
	}

	return c.JSON(http.StatusOK, GatchaPlayResponse{Results: results})
}
