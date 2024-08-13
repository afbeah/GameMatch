package main

import (
	handlers "GamesInsertion/src/Handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	//Rota para o Healthcheck
	e.GET("/health", handlers.HealthCheck)

	e.Logger.Fatal(e.Start(":8080"))

}
