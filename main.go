package main

import (
	"github.com/labstack/echo"
	"log"
	"tune/spotifyclient"
)

func main() {
	client, err := spotifyclient.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	handler := NewHandler(client)
	e := echo.New()

	e.GET("/current", handler.Current)

	e.Logger.Fatal(e.Start(":8080"))
}
