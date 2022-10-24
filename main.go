package main

import (
	handler "golang-url-shortener/handler"
	"golang-url-shortener/store"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Go URL Shortener with Redis 🚀",
		})
	})

	e.POST("/encode", func(c echo.Context) error {
		return handler.CreateShortURL(c)
	})

	e.GET("/decode/:short-url", func(c echo.Context) error {
		return handler.ReturnLongURL(c)
	})

	store.InitializeStore()

	e.Logger.Fatal(e.Start(":1234"))
}
