package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/views/blog"
)

func HandleShowBlog(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, blog.Show(ok))
}
