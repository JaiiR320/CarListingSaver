package handler

import (
	"encoding/json"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func WriteJSON(c echo.Context, status int, v any) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)
	return json.NewEncoder(c.Response()).Encode(v)
}
