package main

import(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlers "GamesInsertion/src/Handlers"
	"go.uber.org/zap"

)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("Logger initialized successfully", zap.String("ex", "logger simple ex"))

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

//Rotas 
e.POST("/games/addmatch", handlers.AddMatch)
e.GET("/games/getmatch", handlers.GetMatch)

e.Logger.Fatal(e.Start(":8080"))

}
