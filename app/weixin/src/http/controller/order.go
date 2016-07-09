package controller

import (
	// "fmt"

	daoConf "dao/conf"
	"encoding/json"
	"fmt"
	. "global"
	apiIndex "http/api/order"
	"logic"
	"strconv"
	"util"

	"github.com/labstack/echo"
)

type OrderController struct{}

// 注册路由
func (self OrderController) RegisterRoute(e *echo.Group) {
	e.Post("/order/prepare", echo.HandlerFunc(self.PrepareOrder))
}

// 准备信息页 form-> goods_list:[{"goods_id":"3","selected":"1","goods_num":"2"}]
func (OrderController) PrepareOrder(ctx echo.Context) error {
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
			goodsNum, err := strconv.ParseInt(goodsInfo["goods_num"], 10, 0)
			if "1" == goodsInfo["selected"] {
				if nil == err && 0 < goodsNum {
					goodsIdList = append(goodsIdList, goodsId)
				}
			}
		}
	}

	// 获取 goodsId 信息
	goodsException, goodsIdMap, err := logic.GetCartInfo(goodsIdList)
	if nil != err {
		return err
	}

	// 获取 goodsId 信息
	shipTimeList := []string{"XX", "XX"}
	// shipTimeList := logic.GetShipTime()
	if nil != err {
		return err
	}

	// 获取地址信息
	myAddress, err := logic.GetDefaultAddress(10)
	if nil != err {
		return err
	}

	// 读入配置信息
	orderConf, err := daoConf.OrderConf()
	if nil != err {
		// log
		return err
	}

	// 拼装接口数据
	orderData := apiIndex.Order{
		Address:      myAddress,
		ShipTimeList: shipTimeList,
	}
	if 0 < len(goodsException) {
		orderData.Alert = fmt.Sprintf(orderConf.Alert, goodsException)
	} else {
		orderData.Alert = ""
	}

	// orderData
	for _, goodsInfo := range goodsList {
		goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 0)
		if nil != err {
			// log
			continue
		}

		goodsNum, _ := strconv.ParseInt(goodsInfo["goods_num"], 10, 0)
		if v, ok := goodsIdMap[goodsId]; ok {
			if goodsNum > int64(v.Storage) {
				goodsNum = int64(v.Storage)
			}

			orderGoodsInfo, _ := logic.GenOrderGoods(goodsIdMap[goodsId], goodsNum)
			orderData.GoodsList = append(orderData.GoodsList, orderGoodsInfo)
		} else {
			// log
			continue
		}
	}

	orderData.Order, _ = logic.GenOrder(orderData.GoodsList)

	return util.Success(ctx, orderData)
}
