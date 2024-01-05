package web

import (
	"encoding/json"
	"net/http"

	"github.com/JaiiR320/carlistingsaver/handler"
	"github.com/JaiiR320/carlistingsaver/types"
	"github.com/JaiiR320/carlistingsaver/view"
	"github.com/labstack/echo/v4"
)

type WebHandler struct {
}

func (h *WebHandler) HandleShowDashboard(c echo.Context) error {
	resp, err := http.Get("http://localhost:3000/api/listing")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var listings []types.Listing

	err = json.NewDecoder(resp.Body).Decode(&listings)
	if err != nil {
		return err
	}

	return handler.Render(c, view.Dashboard(listings))
}
