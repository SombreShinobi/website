package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/about"
)

func HandleShowAbout(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, about.Show(ok))
}
