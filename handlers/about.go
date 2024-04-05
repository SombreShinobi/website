package handlers

import (
	"github.com/SombreShinobi/website/views/about"
	"github.com/labstack/echo/v4"
)

func HandleShowAbout(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, about.Show(ok))
}
