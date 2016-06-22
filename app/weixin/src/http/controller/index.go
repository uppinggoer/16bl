package controller

import (
	// "fmt"

	apiIndex "http/api/index"
	"logic"
	"util"

	"github.com/labstack/echo"
)

type IndexController struct{}

// 注册路由
func (self IndexController) RegisterRoute(e *echo.Group) {
	e.Get("/", echo.HandlerFunc(self.Index))
}

// Index 首页
func (IndexController) Index(ctx echo.Context) error {
	homeConf, goodsIdMap, err := logic.GetHomeData(ctx)
	if nil != err {
		// log
		return util.Fail(ctx, 10, "XXX")
	}

	homeData := apiIndex.Home{}
	homeData.Banner = homeConf.Banner
	homeData.Nav = homeConf.Nav
	for _, itemConf := range homeConf.Class {
		classConf := apiIndex.Class{}
		classConf.Img = itemConf.Img
		classConf.Name = itemConf.Name
		classConf.Color = itemConf.Color

		for _, goodsId := range itemConf.GoodsIdList {
			if v, ok := goodsIdMap[goodsId]; ok {
				classConf.GoodsList = append(classConf.GoodsList, *v)
			} else {
				// log id not exists
			}
		}

		homeData.Class = append(homeData.Class, classConf)
	}

	// time.Sleep(3 * time.Second)
	// return util.Success(ctx, homeData)
	return util.Render(ctx, "home/index", "便利", homeData)
}
