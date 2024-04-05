package handlers

import (
	"github.com/SombreShinobi/website/views/blog"
	"github.com/labstack/echo/v4"
)

func HandleShowBlog(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, blog.Show(ok))
}
