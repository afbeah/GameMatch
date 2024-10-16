package main

import (
	"GamesInsertion/src/Handler"
	"GamesInsertion/src/Service"
	"net/http"
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
	logger.Info("Logger initialized successfully")

	//Criando instancia do Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Criando Service e Handler
	matchService := service.NewMatchService()
	matchHandler := handler.NewMatchHandler(matchService)

	//Rota de verificação do Projeto
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Team!")
	})

	e.GET("/health", handler.HealthCheck)

	//Rotas dos Jogos
	matchs := e.Group("/matchs")
	matchs.POST("", matchHandler.AddGame)
	matchs.GET("/:id", matchHandler.GetGame)
	matchs.PUT("/", matchHandler.UpdateGame)
	matchs.DELETE("/:id", matchHandler.DeleteGame)

	//Inicializando o servidor na porta 8088
	e.Logger.Fatal(e.Start(":8088"))

}
