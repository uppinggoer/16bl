package controller

import (
	"fmt"
	"strconv"
	"strings"

	daoConf "dao/conf"
	. "global"
	apiIndex "http/api/cart"
	"logic"
	"util"

	"github.com/labstack/echo"
)

type CartController struct{}

// 注册路由
func (self CartController) RegisterRoute(e *echo.Group) {
	e.Get("/cart/list", echo.HandlerFunc(self.CartList))
}

// 购物车 首页
func (CartController) CartList(ctx echo.Context) error {
	goodsIdStr := ctx.QueryParam("goods_ids")
	goodsIdListTmp := strings.Split(goodsIdStr, ",")
	if 0 >= len(goodsIdListTmp) {
		return CartEmpty
	}

	goodsIdList := []int64{}
	for _, goodsIdStr := range goodsIdListTmp {
		goodsId, err := strconv.ParseInt(goodsIdStr, 10, 64)
		if nil != err {
			// log
		} else {
			goodsIdList = append(goodsIdList, int64(goodsId))
		}
	}

	goodsException, goodsIdMap, err := logic.GetCartInfo(goodsIdList)
	if nil != err {
		return err
	}

	cartConf, err := daoConf.CartConf()
	if nil != err {
		// log
		return err
	}

	cartData := apiIndex.Cart{}
	cartData.Tips = cartConf.Tips
	cartData.Alert = fmt.Sprintf(cartConf.Alert, goodsException)
	for _, goodsId := range goodsIdList {
		if v, ok := goodsIdMap[goodsId]; ok {
			cartData.GoodsList = append(cartData.GoodsList, *v)
		} else {
			// log
			continue
		}
	}
	return util.Success(ctx, cartData)
	return util.Render(ctx, "cart/list", cartData)
}
