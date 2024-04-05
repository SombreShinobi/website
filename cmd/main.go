package main

import (
	"github.com/labstack/echo/v4"
	"main.go/handlers"
)

func main() {
	e := echo.New()

	e.Static("/dist", "dist")

	e.GET("/", handlers.HandleShowHome)
	e.GET("/home", handlers.HandleShowHome)
	e.GET("/blog", handlers.HandleShowBlog)
	e.GET("/about", handlers.HandleShowAbout)

	e.Logger.Fatal(e.Start(":7331"))
}
