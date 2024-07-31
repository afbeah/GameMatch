package main

import(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlers "GamesInsertion/Handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

//Rotas 
e.POST("/games", handlers.AddGame)
e.GET("/games", handlers.GetGames)

e.Logger.Fatal(e.Start(":8080"))

}
