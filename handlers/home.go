package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/home"
)

func HandleShowHome(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, home.Show(ok))
}
