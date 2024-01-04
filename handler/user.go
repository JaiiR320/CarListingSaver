package handler

import (
	"github.com/JaiiR320/echotest/model"
	"github.com/JaiiR320/echotest/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	name string
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Name: "Jair",
	}

	return Render(c, user.Show(u))
}
