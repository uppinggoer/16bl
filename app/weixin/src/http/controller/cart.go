package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	e.POST("/cart/list", echo.HandlerFunc(self.CartList))
}

// 购物车 首页
func (CartController) CartList(ctx echo.Context) error {
	goodsListStr := ctx.FormValue("goods_list")

	goodsList := []map[string]string{}
	err := json.Unmarshal([]byte(goodsListStr), &goodsList)
	if err != nil {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	if 0 >= len(goodsList) {
		return util.Fail(ctx, 10, CartEmpty.Error())
	}

	// 收集 goodsId
	goodsIdList := []int64{}
	for _, goodsInfo := range goodsList {
		goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 0)
		if nil != err {
			// log
		} else {
			goodsIdList = append(goodsIdList, goodsId)
		}
	}
	// 获取 goodsId 信息
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
	if 0 < len(goodsException) {
		cartData.Alert = fmt.Sprintf(cartConf.Alert, goodsException)
	} else {
		cartData.Alert = ""
	}
	for _, goodsInfo := range goodsList {
		goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 0)
		if nil != err {
			// log
			continue
		}

		var selected = goodsInfo["selected"]
		goodsNum, _ := strconv.ParseInt(goodsInfo["goods_num"], 10, 0)
		if v, ok := goodsIdMap[goodsId]; ok {
			if goodsNum > int64(v.Storage) {
				goodsNum = int64(v.Storage)
			}
			goodsTmp := apiIndex.Goods{GoodsInfo: *v, Selected: selected, GoodsNum: strconv.FormatInt(goodsNum, 10)}
			cartData.GoodsList = append(cartData.GoodsList, goodsTmp)
		} else {
			// log
			continue
		}
	}
	return util.Success(ctx, cartData)
}

// 生成购物车页的 html 不提供外部接口
func (CartController) GenCartIndexHtml(ctx echo.Context) error {
	return util.Render(ctx, "cart/list", map[string]interface{}{})
}
