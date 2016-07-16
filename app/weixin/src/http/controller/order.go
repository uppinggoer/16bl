package controller

import (
	// "fmt"

	daoConf "dao/conf"
	daoSql "dao/sql"
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
	e.Get("/order/list", echo.HandlerFunc(self.MyOrderList))
	e.Post("/order/prepare", echo.HandlerFunc(self.PrepareOrder))
	e.Post("/order/do_order", echo.HandlerFunc(self.DoOrder))
}

// 确认信息页 form-> goods_list:[{"goods_id":"3","selected":"1","goods_num":"2"}]
func (OrderController) PrepareOrder(ctx echo.Context) error {
	// fmt.Printf("%#v,%#v", ctx.Request().FormParams(), ctx.Request().Header())
	// write();

	// 获取购物车商品列表
	goodsList, err = getRequestGoodsList()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 收集 goodsId
	goodsIdList := []int64{}
	for _, goodsInfo := range goodsList {
		goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 64)
		if nil != err {
			// log
		} else {
			goodsNum, err := strconv.ParseInt(goodsInfo["goods_num"], 10, 64)
			if "1" == goodsInfo["selected"] {
				if nil == err && 0 < goodsNum {
					goodsIdList = append(goodsIdList, goodsId)
				}
			}
		}
	}

	// 获取 购物车商品 信息
	goodsException, goodsIdMap, err := logic.GetCartInfo(goodsIdList)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// // 遍历商品 验证库存不足信息
	// for _, goodsInfo := range goodsList {
	// 	goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 64)
	// 	if nil != err {
	// 		// log
	// 		continue
	// 	}

	// 	var selected = goodsInfo["selected"]
	// 	goodsNum, _ := strconv.ParseInt(goodsInfo["goods_num"], 10, 64)
	// 	if v, ok := goodsIdMap[goodsId]; ok {
	// 		if goodsNum > int64(v.Storage) {
	// 			goodsNum = int64(v.Storage)
	// 		}
	// 		goodsTmp := apiIndex.Goods{GoodsInfo: *v, Selected: selected, GoodsNum: strconv.FormatInt(goodsNum, 10)}
	// 		cartData.GoodsList = append(cartData.GoodsList, goodsTmp)
	// 	} else {
	// 		// log
	// 		continue
	// 	}
	// }
	// 验证并修正库存信息  如果只有3个，购买5个，会强制改为3个
	goodsException, goodsIdMap, err = logic.VerifyGoodsNum(goodsIdList)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 获取 goodsId 信息
	shipTimeList := []string{"XX", "XX"}
	// shipTimeList := logic.GetShipTime()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 获取地址信息
	myAddressList, err := daoSql.GetAddressListByUid(10, true)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	myAddress := daoSql.Address{}
	if 0 < len(myAddressList) {
		myAddress := myAddressList[0]
	}

	// 生成预处理订单
	orderInfo, orderGoodsList, err := logic.GenOrder(goodsList)
	if nil != err {
		// log
	}

	// 拼装接口数据
	orderData := apiIndex.Order{
		Address:      myAddress,
		ShipTimeList: shipTimeList,
		GoodsList:    orderGoodsList,
		Order:        orderInfo,
	}

	// 读入配置信息
	orderConf, err := daoConf.OrderConf()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	if 0 < len(goodsException) {
		orderData.Alert = fmt.Sprintf(orderConf.Alert, goodsException)
	} else {
		orderData.Alert = ""
	}

	return util.Success(ctx, orderData)
}

// 确认信息页 form-> goods_list:[{"goods_id":"3","selected":"1","goods_num":"2"}]
func (OrderController) DoOrder(ctx echo.Context) error {
	// // 避免同一个订单重复提交
	// $curOrderMd5 = md5(json_encode($_REQUEST));
	// $preOrderMd5 = Redis::getValue('do_order', $this->curUser['uid']);
	// if ($curOrderMd5 === $preOrderMd5) {
	// 	Logger::error(__METHOD__, "订单重复提交");
	// 	return $this->fail('请勿重复提交！');
	// }
	// // 3秒内不允许重复提交同一订单
	// Redis::setValue('do_order', $this->curUser['uid'], $curOrderMd5, 3);

	// 获取购物车商品列表
	goodsList, err = getRequestGoodsList()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	uid = 10
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

	// 获取 商品详情
	goodsException, goodsIdMap, err := logic.GetCartInfo(goodsIdList)
	if nil != err {
		return util.Fail(ctx, 10, err.Error())
	}

	// 获取地址信息
	address, err = fetchAddress()
	if nil != err {
		return util.Fail(ctx, 10, "地址信息无效")
	}

	// 读入配置信息
	orderConf, err := daoConf.OrderConf()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	if 0 < len(goodsException) {
		return util.Fail(ctx, 10, fmt.Sprintf(orderConf.Alert, goodsException))
	}

	// 生成预处理订单
	orderInfo, orderGoodsList, err := logic.SubmitOrder(goodsList, address)
	if nil != err {
		// log
	}
	// 拼装接口数据
	orderData := apiIndex.Order{
		Address:      myAddress,
		ShipTimeList: shipTimeList,
		GoodsList:    orderGoodsList,
		Order:        orderInfo,
	}
	return util.Success(ctx, orderData)
}

// 订单列表页
func (OrderController) MyOrderList(ctx echo.Context) error {
	// 获取订单列表信息
	// uid, base_id, rn
	myOrderMapList, hasMore, err := logic.GetMyOrderList(10, 0, 20)
	if nil != err {
		return err
	}
	// 拼装接口数据
	orderData := &apiIndex.OrderList{
		HasMore: hasMore,
	}

	// orderData
	for _, v := range myOrderMapList {
		orderData.List = append(orderData.List, &apiIndex.Order{
			Address: v["addressInfo"].(*daoSql.Address),
			OrderBase: apiIndex.OrderBase{
				Order:     v["order"].(*daoSql.Order),
				GoodsList: v["goodsList"].([]*daoSql.OrderGoods),
			},
		})
	}
	var v = myOrderMapList[0]
	var a = &apiIndex.Order{
		Address: v["addressInfo"].(*daoSql.Address),
		OrderBase: apiIndex.OrderBase{
			Order:     v["order"].(*daoSql.Order),
			GoodsList: v["goodsList"].([]*daoSql.OrderGoods),
		},
	}
	for i := 1; i < 20; i++ {
		orderData.List = append(orderData.List, a)
	}

	return util.Success(ctx, orderData)
}

// 生成订单列表页页的 html 不提供外部接口
func (OrderController) GenOrderListHtml(ctx echo.Context) error {
	return util.Render(ctx, "order/list", "订单列表", map[string]interface{}{})
}

// 根据 请求参数 goods_list 获取购物车中商品信息
func getRequestGoodsList() ([]map[string]string, error) {
	goodsListStr := ctx.FormValue("goods_list")

	goodsList := []map[string]string{}
	err := json.Unmarshal([]byte(goodsListStr), &goodsList)
	if err != nil {
		// log
		return goodsList, err.Error()
	}
	if 0 >= len(goodsList) {
		return goodsList, CartEmpty.Error()
	}
	return goodsList, nil
}

// 提交订单时新地址(address_id<=0) 会先插入地址表
func fetchAddress() (daoSql.Address, err) {
	uid := 10

	// 获取提交订单时指定的地址
	addressId := ctx.FormValue("address_id")
	if 0 < addressId {
		myAddressMap, err := daoSql.GetAddressListById([]int64{addressId})
		if nil != err {
			// log
			return daoSql.Address{}, err
		}
		myAddress, ok := myAddressMap[addressId]
		if !ok || uid != myAddress.MemberId {
			// log
			return daoSql.Address{}, RecordEmpty
		}
		return myAddress, nil
	}

	// 插入新的地址信息
	trueName := ctx.FormValue("true_name")
	gender := ctx.FormValue("gender")
	liveArea := ctx.FormValue("live_area")
	address := ctx.FormValue("address")
	mobile := ctx.FormValue("mobile")
	// 显式提取地址信息
	addressInfo := map[string]string{
		"true_name": trueName,
		"gender":    gender,
		"live_area": liveArea,
		"address":   address,
		"mobile":    mobile,
	}
	address, err = daoSql.SaveMyAddress(uid, addressInfo)
	if nil != err {
		// log
		return daoSql.Address{}, err
	}
	return address, err
}
