package controller

import (
	"fmt"

	daoConf "dao/conf"
	apiIndex "http/api"
	"logic"
	"util"

	"github.com/labstack/echo"
)

type CartController struct{}

// 注册路由
func (self CartController) RegisterRoute(e *echo.Group) {
	e.POST("/cart/list", echo.HandlerFunc(self.CartList))
}

// 购物车 首页
func (CartController) CartList(ctx echo.Context) error {
	goodsList, err := getCartGoodsList(ctx)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	// 收集 goodsId
	goodsIdList := []uint64{}
	for _, goodsInfo := range goodsList {
		goodsIdList = append(goodsIdList, goodsInfo.GoodsId)
	}

	// 获取 goodsId 信息
	goodsException, goodsIdMap, _ := logic.GetCartInfo(goodsIdList)
	// 校正库存
	logic.VerifyGoodsNum(goodsIdMap, goodsList)

	cartConf, err := daoConf.CartConf()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	cartData := apiIndex.Cart{}
	cartData.Tips = cartConf.Tips
	if 0 < len(goodsException) {
		cartData.Alert = fmt.Sprintf(cartConf.Alert, goodsException)
	} else {
		cartData.Alert = ""
	}

	// 遍历商品 填充接口数据
	for _, goodsInfo := range goodsList {
		if v, ok := goodsIdMap[goodsInfo.GoodsId]; ok {
			goodsTmp := apiIndex.CartGoods{
				GoodsInfo: &apiIndex.Goods{Goods: v},
				Selected:  util.Itoa(goodsInfo.Selected),
				GoodsNum:  util.Itoa(goodsInfo.GoodsNum),
			}
			cartData.GoodsList = append(cartData.GoodsList, &goodsTmp)
		} else {
			// log
			continue
		}
	}

	cartData.Format()
	return util.Success(ctx, cartData)
}

// 生成购物车页的 html 不提供外部接口
func (CartController) GenCartIndexHtml(ctx echo.Context) error {
	return util.Render(ctx, "cart/list", "购物车", map[string]interface{}{})
}
