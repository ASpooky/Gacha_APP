package handler

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ASpooky/ca_tech_dojo/model"
	"github.com/ASpooky/ca_tech_dojo/types"
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

	var possessions []model.Possession
	emissions := c.Get("emissions").(types.EmissionsByRarity)

	for i := 0; i < req.Times; i++ {
		//1~100のrandomな整数→レアリティ
		randInt := rand.Intn(MAX-MIN) + MIN
		rarity := utils.NumToRarity(randInt)
		//log.Println(rarity)

		//レアリティからキャラクターID

		lenEmissions := len(emissions[rarity])
		randEmission := rand.Intn(lenEmissions)

		var character model.Character

		//キャラクターIDからキャラクター
		if err := gh.db.Where("id=?", emissions[rarity][randEmission]).First(&character).Error; err != nil {
			log.Println("err :", err)
			return c.String(http.StatusInternalServerError, "err: Could not find character")
		}
		//log.Println(character)

		possessions = append(possessions, model.Possession{UserID: user.ID, CharacterID: character.ID})

		//userIDとcharacterIDからposessionから抽出
		results = append(results, CharacterResponse{CharacterID: strconv.FormatUint(uint64(character.ID), 10), Name: character.Name})
	}

	if err := gh.db.Create(&possessions).Error; err != nil {
		log.Println("err:", err)
		return c.String(http.StatusInternalServerError, "err: Could not record possessions")
	}

	return c.JSON(http.StatusOK, GatchaPlayResponse{Results: results})
}
