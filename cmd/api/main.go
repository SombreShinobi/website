package main

import (
	"os"

	"github.com/SombreShinobi/website/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.Static("/dist", "dist")

	e.GET("/", handlers.HandleShowHome)
	e.GET("/home", handlers.HandleShowHome)
	e.GET("/blog", handlers.HandleShowBlog)
	e.GET("/about", handlers.HandleShowAbout)

	e.Logger.Fatal(e.Start(os.Getenv("URL")))
}
