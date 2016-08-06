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
	e.POST("/address/modify", echo.HandlerFunc(self.AddressModify))
	e.POST("/address/delete", echo.HandlerFunc(self.AddressDel))
}

// 地址列表
func (AddressController) AddressList(ctx echo.Context) error {
	addressList, err := daoSql.GetAddressListByUid(uint64(10), false)
	if nil != err {
		// log
		return util.Fail(ctx, 10, "XXX")
	}

	return util.Render(ctx, "address/list", "送货地址", addressList)
}

// 地址修改
func (AddressController) AddressModify(ctx echo.Context) error {
	uid := util.Atoi(ctx.FormValue("uid"), 64, false).(uint64)

	// 插入新的地址信息
	trueName := ctx.FormValue("true_name")
	liveArea := ctx.FormValue("live_area")
	mobile := ctx.FormValue("mobile")
	addressText := ctx.FormValue("address")
	addressId := util.Atoi(ctx.FormValue("addressId"), 64, false).(uint64)

	// 显式提取地址信息
	addressInfo := daoSql.UserAddressInfo{
		TrueName: trueName,
		LiveArea: liveArea,
		Address:  addressText,
		Mobile:   mobile,
	}

	var err error
	var myAddress *daoSql.Address
	if 0 >= addressId {
		myAddress, err = daoSql.SaveMyAddress(uid, &addressInfo)
	} else {
		myAddress, err = daoSql.UpdateAddressById(addressId, &addressInfo)
		myAddress.MemberId = uid
		myAddress.Id = addressId
	}
	if nil != err {
		// log
		return util.Fail(ctx, 10, "更新出错")
	} else {
		return util.Success(ctx, myAddress)
	}
}

// 地址删除
func (AddressController) AddressDel(ctx echo.Context) error {
	uid := util.Atoi(ctx.FormValue("uid"), 64, false).(uint64)
	addressId := util.Atoi(ctx.FormValue("addressId"), 64, false).(uint64)

	err := daoSql.DelMyAddress(uid, addressId)
	if nil != err {
		// log
		return util.Fail(ctx, 10, "地址不存在或不是您的地址")
	} else {
		return util.Success(ctx, nil)
	}
}
