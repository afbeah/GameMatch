package handler

import (
	model "GamesInsertion/src/Model"
	service "GamesInsertion/src/Service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MatchHandler struct {
	Service service.MatchService
}

func NewMatchHandler(s service.MatchService) *MatchHandler {
	return &MatchHandler{Service: s}
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "UP",
	})
}

func (h *MatchHandler) AddGame(c echo.Context) error {
	var match model.Match

	if err := c.Bind(&match); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := h.Service.AddGame(&match); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, match)
}

func (h *MatchHandler) GetGame(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	match, err := h.Service.GetGame(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, match)
}

func (h *MatchHandler) UpdateGame(c echo.Context) error {
	var match model.Match
	if err := c.Bind(&match); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := h.Service.UpdateGame(&match); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, match)
}

func (h *MatchHandler) DeleteGame(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.Service.DeleteGame(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *MatchHandler) ListGames(c echo.Context) error {
    matches, err := h.Service.GetAllGames()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not retrieve games"})
    }
    return c.JSON(http.StatusOK, matches) 
}
