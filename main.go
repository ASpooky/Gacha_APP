package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/model"
	"github.com/ASpooky/ca_tech_dojo/router"
	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db := db.NewDB()
	e := router.NewRouter()
	uh := NewUserHandler(db)

	u := e.Group("/user")
	u.POST("/create", uh.CreateUser)
	u.GET("/get", uh.GetUser)
	u.PUT("/update", uh.UpdateUser)

	e.Logger.Fatal(e.Start(":8080"))
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	user := model.User{}

	//bodyから空のuserequest構造体に値をバインドする.
	if err := c.Bind(&user); err != nil {
		log.Println("err :", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not bind user struct!")
	}

	//validation

	//headerにx-tokenを設定する.

	text := user.Name + time.Now().String()

	hash := sha256.Sum256([]byte(text))
	user.Token = hex.EncodeToString(hash[:])

	if err := uh.db.Create(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err:Could not create user at db!")
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	user := model.User{}

	user.Token = c.Request().Header.Get("X-Token")

	if err := uh.db.Where("token=?", user.Token).First(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err:Could not get user at db!")
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	user := model.User{}

	//bodyから空のuserequest構造体に値をバインドする.
	if err := c.Bind(&user); err != nil {
		log.Println("err :", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not bind user struct!")
	}

	//reauest headerから値を取得.
	user.Token = c.Request().Header.Get("X-Token")

	result := uh.db.Model(user).Clauses(clause.Returning{}).Where("token=?", user.Token).Update("name", user.Name)
	if result.Error != nil {
		log.Println("err :", result.Error)
		return c.String(http.StatusInternalServerError, "err:Could not update user at db!")
	}
	if result.RowsAffected < 1 {
		return c.String(http.StatusInternalServerError, "err:Could not update user at db!")
	}

	return c.JSON(http.StatusOK, user)
}
