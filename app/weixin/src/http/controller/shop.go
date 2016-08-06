package controller

import (
	// . "global"

	apiIndex "http/api"
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
			ClassName: item.ClassName,
			ClassId:   item.ClassId,
		}
		if v, ok := goodsIdMap[item.ClassId]; ok {
			classItem.GoodsList = make([]*apiIndex.Goods, len(v))
			for idx, goodsItem := range v {
				classItem.GoodsList[idx] = &apiIndex.Goods{Goods: goodsItem}
			}
		}

		cartData.ClassList = append(cartData.ClassList, &classItem)
	}

	cartData.Format()
	// return util.Success(ctx, cartData)
	return util.Render(ctx, "shop/list", "商品详情", cartData)
}
