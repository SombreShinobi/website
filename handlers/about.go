package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/about"
)

type AboutHanlder struct{}

func (aHandler *AboutHanlder) Show(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, about.Show(ok))
}
