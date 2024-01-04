package main

import (
	"context"

	"github.com/JaiiR320/echotest/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.Use(withUser)

	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	key := "user"
	value := "Jair"
	return func(c echo.Context) error {
		c.Set(key, value)
		ctx := context.WithValue(c.Request().Context(), key, value)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
