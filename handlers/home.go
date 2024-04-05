package handlers

import (
	"github.com/SombreShinobi/website/views/home"
	"github.com/labstack/echo/v4"
)

func HandleShowHome(ctx echo.Context) error {
	_, ok := ctx.Request().Header["Hx-Request"]
	return render(ctx, home.Show(ok))
}
