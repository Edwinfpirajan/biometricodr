package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var blackListTokens = []string{} //aquí se almacenan los tokens invalidados

// Función para invalidar el token
func invalidateToken(token string) {
	blackListTokens = append(blackListTokens, token)
}

// Función de logout
func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("authorization")

	if authHeader == "" || authHeader[:6] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	tokenString := authHeader[7:]
	invalidateToken(tokenString)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Sesión cerrada con exito",
	})
}
