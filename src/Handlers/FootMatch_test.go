package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"bytes"
    "encoding/json"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    "GamesInsertion/src/Models"
)

func TestHealthCheck(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, HealthCheck(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.JSONEq(t, `{"status":"UP"}`, rec.Body.String())
    }
}

func TestAddMatch(t *testing.T) {
    e := echo.New()

    match := models.Match{ // Ajuste o modelo conforme necessário
        // Preencha com dados apropriados para o teste
    }
    jsonMatch, _ := json.Marshal(match)

    req := httptest.NewRequest(http.MethodPost, "/games/addmatch", bytes.NewBuffer(jsonMatch))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, AddMatch(c)) {
        assert.Equal(t, http.StatusCreated, rec.Code)
        assert.JSONEq(t, string(jsonMatch), rec.Body.String())
    }
}

func TestGetMatch(t *testing.T) {
    e := echo.New()

    // Configuração de um exemplo de match
    match := models.Match{} // Preencha conforme necessário

    // Defina o match no contexto, se necessário
    // Por exemplo, usando um middleware para definir o estado ou banco de dados fake

    req := httptest.NewRequest(http.MethodGet, "/games/getmatch", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, GetMatch(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.JSONEq(t, `{} /* aqui vai o JSON do match que você espera */`, rec.Body.String())
    }
}