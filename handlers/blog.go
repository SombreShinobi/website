package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/blog"
)

type BlogHandler struct{}

func (bHandler *BlogHandler) Show(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, blog.Show(ok))
}
