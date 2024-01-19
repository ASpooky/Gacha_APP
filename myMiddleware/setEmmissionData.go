package myMiddleware

import (
	"github.com/ASpooky/ca_tech_dojo/types"
	"github.com/labstack/echo/v4"
)

func SetEmmissionDataMiddleware(data types.EmissionsByRarity) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("emissions", data)
			return next(c)
		}
	}
}
