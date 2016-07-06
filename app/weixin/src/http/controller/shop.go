package controller

import (
	// . "global"

	apiIndex "http/api/shop"
	"logic"
	"util"

	"github.com/labstack/echo"
)

type ShopController struct{}

// 注册路由
func (self ShopController) RegisterRoute(e *echo.Group) {
	e.GET("/shop/list", echo.HandlerFunc(self.ShopList))
}

// 购物车 首页
func (ShopController) ShopList(ctx echo.Context) error {
	// 商品列表基础信息
	classList, goodsIdMap, err := logic.GetShopData(ctx)
	if err != nil {
		return util.Fail(ctx, 1, err.Error())
	}

	cartData := apiIndex.Shop{}
	for _, item := range classList {
		classItem := apiIndex.Class{
			ClassName: item.Name,
			ClassId:   item.Id,
		}
		if v, ok := goodsIdMap[item.Id]; ok {
			classItem.GoodsList = v
		}

		cartData.ClassList = append(cartData.ClassList, classItem)
	}

	// return util.Success(ctx, cartData)
	return util.Render(ctx, "shop/list", "商品详情", cartData)
}
