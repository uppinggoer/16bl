package http

import (
	. "http/controller"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Group) {
	new(IndexController).RegisterRoute(e)
}
