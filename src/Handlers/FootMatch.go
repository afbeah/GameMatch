package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	models "GamesInsertion/src/Models"
)

type Games struct{
	game []models.Match
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
