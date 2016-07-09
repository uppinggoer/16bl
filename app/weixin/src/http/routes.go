package http

import (
	. "http/controller"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Group) {
	new(IndexController).RegisterRoute(e)
	new(CartController).RegisterRoute(e)
	new(ShopController).RegisterRoute(e)
	new(OrderController).RegisterRoute(e)

	// 测试框架配合 util.fake_context 使用  正式环境中不可过滤掉，目前未做处理
	new(TestController).RegisterRoute(e)
}
