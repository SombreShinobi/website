package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/home"
)

type HomeHandler struct{}

func (hHandler *HomeHandler) Show(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, home.Show(ok))
}
