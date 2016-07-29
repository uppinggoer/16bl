package controller

import (
	"util"

	daoSql "dao/sql"

	"github.com/labstack/echo"
)

type AddressController struct{}

// 注册路由
func (self AddressController) RegisterRoute(e *echo.Group) {
	e.GET("/address/list", echo.HandlerFunc(self.AddressList))
}

// 地址列表
func (AddressController) AddressList(ctx echo.Context) error {
	addressList, err := daoSql.GetAddressListByUid(uint64(10), false)
	if nil != err {
		// log
		return util.Fail(ctx, 10, "XXX")
	}

	return util.Render(ctx, "address/list", "送货地址", addressList)
	// return util.Success(ctx, addressList)
}
