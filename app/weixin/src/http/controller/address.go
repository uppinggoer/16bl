package controller

import (
	"github.com/labstack/echo"
)

type AddressController struct{}

// 注册路由
func (self AddressController) RegisterRoute(e *echo.Group) {
	e.GET("/address/list", echo.HandlerFunc(self.AddressList))
}

// 地址列表
func (AddressController) AddressList(ctx echo.Context) error {
	return nil
}
