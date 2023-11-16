package router

import (
	"net/http"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewRouter() *echo.Echo {
	db := db.NewDB()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	uh := handler.NewUserHandler(db)

	u := e.Group("/user")
	u.POST("/create", uh.CreateUser)
	u.GET("/get", uh.GetUser)
	u.PUT("/update", uh.UpdateUser)

	return e
}
