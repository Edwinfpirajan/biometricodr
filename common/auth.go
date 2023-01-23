package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("authorization")

		if authHeader == "" || authHeader[:6] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
