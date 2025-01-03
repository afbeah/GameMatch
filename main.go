package main

import (
	handler "GamesInsertion/src/Handler"
	model "GamesInsertion/src/Model"
	service "GamesInsertion/src/Service"
	"net/http"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("footmatch.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.AutoMigrate(&model.Match{},)
}

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Logger initialized successfully")

	InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())


	matchService := service.NewMatchService(db)
	matchHandler := handler.NewMatchHandler(matchService)


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Team!")
	})

	e.GET("/health", handler.HealthCheck)

	matchs := e.Group("/matches")
	matchs.GET("", matchHandler.ListGames)
	matchs.POST("", matchHandler.AddGame)
	matchs.GET("/:id", matchHandler.GetGame)
	matchs.PUT("/:id", matchHandler.UpdateGame)
	matchs.DELETE("/:id", matchHandler.DeleteGame)

	e.Logger.Fatal(e.Start(":8088"))

}
