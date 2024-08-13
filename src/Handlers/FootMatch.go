package handlers

import (
	"net/http"

	models "GamesInsertion/src/Models"

	"github.com/labstack/echo/v4"
)

type Games struct {
	game []models.Match
}

// Healthcheck
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "UP",
	})
}

func AddMatch(c echo.Context) error {
	var game models.Match
	if err := c.Bind(&game); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, game)

}

func GetMatch(c echo.Context) error {
	var game models.Match
	return c.JSON(http.StatusOK, game)
}
