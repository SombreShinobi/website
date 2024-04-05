package main

import (
	"github.com/labstack/echo/v4"
	"main.go/handlers"
)

func main() {
	e := echo.New()

	e.Static("/dist", "dist")

	homeHandler := handlers.HomeHandler{}
	e.GET("/", homeHandler.Show)
	e.GET("/home", homeHandler.Show)

	blogHanlder := handlers.BlogHandler{}
	e.GET("/blog", blogHanlder.Show)

	aboutHanlder := handlers.AboutHanlder{}
	e.GET("/about", aboutHanlder.Show)

	e.Logger.Fatal(e.Start(":7331"))
}
