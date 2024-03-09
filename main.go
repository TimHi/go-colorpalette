package main

import (
	"github.com/timhi/golorpalette/src"
)

func main() {
	src.GetColors("https://f4.bcbits.com/img/0018109799_2.jpg")

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	_, err := src.GetColors("https://cdn.discordapp.com/attachments/122739462210846721/1216001747496079443/image0.jpg?ex=65fecce0&is=65ec57e0&hm=3e799129fa778c14fc88566484ebc234aae9abd6296b2857d477d4471fdadb53&")
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}

	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
