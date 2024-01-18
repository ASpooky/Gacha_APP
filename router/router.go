package router

import (
	"log"
	"net/http"

	"github.com/ASpooky/ca_tech_dojo/db"
	"github.com/ASpooky/ca_tech_dojo/handler"
	"github.com/ASpooky/ca_tech_dojo/model"
	"github.com/ASpooky/ca_tech_dojo/types"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func setEmmissionDataMiddleware(data types.EmissionsByRarity) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("emissions", data)
			return next(c)
		}
	}
}

func NewRouter() *echo.Echo {
	db := db.NewDB()

	var emissions []model.Emission

	if err := db.Find(&emissions).Error; err != nil {
		log.Fatalln("err :", err)
	}

	//rarity別にキャラクターIDを格納するmap
	//contextにこれを格納してhandlerで使用する.
	emissionsByRarity := make(types.EmissionsByRarity)

	for _, v := range emissions {
		emissionsByRarity[v.Rarity] = append(emissionsByRarity[v.Rarity], v.CaracterID)
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	uh := handler.NewUserHandler(db)
	gh := handler.NewGatchaHandler(db)
	ch := handler.NewCharacterHandler(db)

	u := e.Group("/user")
	u.POST("/create", uh.CreateUser)
	u.GET("/get", uh.GetUser)
	u.PUT("/update", uh.UpdateUser)

	g := e.Group("/gacha")
	g.Use(setEmmissionDataMiddleware(emissionsByRarity))
	g.POST("/draw", gh.PlayGatcha)

	c := e.Group("/character")
	c.GET("/list", ch.GetCharacterList)

	return e
}
