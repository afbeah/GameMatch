package handlers

import (
	"net/http"
	//"sync"

	"github.com/labstack/echo/v4"
	models "GamesInsertion/Models"
)

var (
	games []models.Play
	//mu sync.Mutex
)

func AddGame(c echo.Context) error {
	var game models.Play
	if err := c.Bind(&game); err != nil {
		return err
	}

//mu.Lock()
game.ID = len(games) + 1
games = append(games, game)
//mu.Unlock()

return c.JSON(http.StatusCreated, game)

}

func GetGames(c echo.Context) error {
	//mu.Lock()
	//defer mu.Unlock()

	return c.JSON(http.StatusOK, games)
}
