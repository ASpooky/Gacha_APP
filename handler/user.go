package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"log"

	"github.com/ASpooky/ca_tech_dojo/model"
	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

type UserCreateRequest struct {
	Name string
}

//+a:各関数の前に認証middlewareを追加してtokenで認証,idで検索できるようする.

func (uh *UserHandler) CreateUser(c echo.Context) error {
	var req model.UserCreateRequest

	//bodyから空のuserequest構造体に値をバインドする.
	if err := c.Bind(&req); err != nil {
		log.Println("err :", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not bind request struct!")
	}

	//uuidを使用する
	token := uuid.New().String()

	user := model.User{
		Name:  req.Name,
		Token: token,
	}

	if err := uh.db.Create(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err:Could not create user at db!")
	}

	return c.JSON(http.StatusOK, model.UserCreateResponse{
		Token: user.Token,
	})
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	var user model.User

	token := c.Request().Header.Get("X-Token")

	if err := uh.db.Where("token=?", token).First(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err:Could not get user at db!")
	}

	return c.JSON(http.StatusOK, model.UserGetResponse{
		Name: user.Name,
	})
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	var req model.UserUpdateRequest
	var user model.User

	//bodyから空のuserequest構造体に値をバインドする.
	if err := c.Bind(&req); err != nil {
		log.Println("err :", err.Error())
		return c.String(http.StatusInternalServerError, "err: Could not bind request struct!")
	}

	//reauest headerから値を取得.
	token := c.Request().Header.Get("X-Token")

	//modelのgetとupdateを分ける.
	//一緒にしてしまうとレコードがなかったのか更新できなかったのかのエラーが分かりづらくなる.
	if err := uh.db.Where("token=?", token).First(&user).Error; err != nil {
		log.Println("err:", err.Error())
		return c.String(http.StatusInternalServerError, "err:Could not get user at db!")
	}

	user.Name = req.Name

	result := uh.db.Save(&user)
	if result.Error != nil {
		log.Println("err :", result.Error)
		return c.String(http.StatusInternalServerError, "err:Could not update user at db!")
	}
	if result.RowsAffected < 1 {
		return c.String(http.StatusInternalServerError, "err:Could not update user at db!")
	}

	return c.String(http.StatusOK, "OK !")
}
