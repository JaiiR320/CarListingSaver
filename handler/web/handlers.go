package web

import (
	"github.com/JaiiR320/carlistingsaver/handler"
	"github.com/JaiiR320/carlistingsaver/view"
	"github.com/labstack/echo/v4"
)

type WebHandler struct {
}

func (h *WebHandler) HandleShowHome(c echo.Context) error {
	return handler.Render(c, view.Dashboard())
}
