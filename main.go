package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timhi/golorpalette/src"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.GET("/colors/:url", func(c echo.Context) error {
		url := c.Param("url")
		log.Println("Request received")
		colors, err := src.GetColors(url)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, colors)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
