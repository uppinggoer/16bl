package controller

import (
	// "fmt"

	daoConf "dao/conf"
	daoSql "dao/sql"
	"encoding/json"
	"fmt"
	. "global"
	apiIndex "http/api"
	"logic"
	"time"
	"util"

	"github.com/labstack/echo"
)

type OrderController struct{}

// 注册路由
func (self OrderController) RegisterRoute(e *echo.Group) {
	e.Get("/order/list", echo.HandlerFunc(self.MyOrderList))
	e.Get("/order/detail", echo.HandlerFunc(self.Detail))
	e.Post("/order/prepare", echo.HandlerFunc(self.PrepareOrder))
	e.Post("/order/do_order", echo.HandlerFunc(self.DoOrder))
	e.Post("/order/cancel_rder", echo.HandlerFunc(self.CancelOrder))
	e.Post("/order/eval_order", echo.HandlerFunc(self.EvalOrder))
}

// 确认信息页 form-> goods_list:[{"goods_id":"3","selected":"1","goods_num":"2"}]
func (OrderController) PrepareOrder(ctx echo.Context) error {
	uid := ctx.Get("uid").(uint64)
	fmt.Printf("%#v", uid)

	goodsList, err := getCartGoodsList(ctx)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 收集 goodsId
	goodsIdList := []uint64{}
	for _, goodsInfo := range goodsList {
		if 1 == goodsInfo.Selected {
			if 0 < goodsInfo.GoodsNum {
				goodsIdList = append(goodsIdList, goodsInfo.GoodsId)
			}
		}
	}

	// 获取 goodsId 信息
	goodsException, goodsIdMap, _ := logic.GetCartInfo(goodsIdList)
	// 验证并修正库存信息  如果只有3个，购买5个，会强制改为3个
	goodsNoStorageException, _ := logic.VerifyGoodsNum(goodsIdMap, goodsList)

	// 获取 goodsId 信息
	shipTimeList := []string{"XX", "XX"}
	// shipTimeList := logic.GetShipTime()

	// 获取用户所有地址
	myAddressList, err := daoSql.GetAddressListByUid(10, false)
	if nil != err && RecordEmpty != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}
	var myAddress *daoSql.Address
	for idx, addressItem := range myAddressList {
		// 默认取第一个
		if 0 == idx {
			myAddress = addressItem
		}
		// 取默认地址
		if uint8(1) == addressItem.IsDefault {
			myAddress = addressItem
		}
	}
	if nil == myAddress {
		myAddress = &daoSql.Address{}
	} else {
		myAddress.IsDefault = 1
	}

	// 读入配置信息
	orderConf, err := daoConf.OrderConf()
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 生成预处理订单
	orderMap := logic.OrderMap{
		Address:    myAddress,
		GoodsIdMap: goodsIdMap,
		GoodsList:  goodsList,
	}
	orderInfo, orderGoodsList, err := logic.GenOrder(uid, orderMap)
	if nil != err {
		// log
	}
	// 过滤订单参数
	orderInfo.Filter()

	arrApiOrderGoods := make([]*apiIndex.OrderGoods, len(orderGoodsList))
	for idx, item := range orderGoodsList {
		arrApiOrderGoods[idx] = &apiIndex.OrderGoods{OrderGoods: item}
	}
	// 拼装接口数据
	orderData := apiIndex.Order{
		Address:      (*apiIndex.AddressType)(myAddress),
		ShipTimeList: shipTimeList,
		OrderInfo: apiIndex.OrderInfo{
			GoodsList: arrApiOrderGoods,
			Order:     &apiIndex.OrderBase{Order: orderInfo},
		},
	}
	if 0 < len(goodsException) {
		orderData.Alert = fmt.Sprintf(orderConf.Alert, goodsException)
	} else if 0 < len(goodsNoStorageException) {
		orderData.Alert = fmt.Sprintf(orderConf.StorageAlert, goodsNoStorageException)
	} else {
		orderData.Alert = ""
	}

	// 格式化地址列表
	for _, addressItem := range myAddressList {
		orderData.AddressList = append(orderData.AddressList, (*apiIndex.AddressType)(addressItem))
	}

	orderData.Format()
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

	var uid uint64
	uid = 10

	// 获取购物车商品列表
	goodsList, err := getCartGoodsList(ctx)
	if nil != err {
		// log
		return util.Fail(ctx, 10, err.Error())
	}

	// 收集 goodsId
	goodsIdList := []uint64{}
	for _, goodsInfo := range goodsList {
		if 1 == goodsInfo.Selected {
			if 0 < goodsInfo.GoodsNum {
				goodsIdList = append(goodsIdList, goodsInfo.GoodsId)
			}
		}
	}

	// 获取 商品详情
	goodsException, goodsIdMap, _ := logic.GetCartInfo(goodsIdList)
	// 验证并修正库存信息  如果只有3个，购买5个，会强制改为3个
	goodsNoStorageException, _ := logic.VerifyGoodsNum(goodsIdMap, goodsList)

	// 获取地址信息
	address, err := fetchAddress(ctx)
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
	if 0 < len(goodsNoStorageException) {
		return util.Fail(ctx, 10, fmt.Sprintf(orderConf.StorageAlert, goodsNoStorageException))
	}

	// 提交订单
	orderMap := logic.OrderMap{
		Address:      address,
		GoodsIdMap:   goodsIdMap,
		GoodsList:    goodsList,
		ExceptTime:   time.Now().Unix(),
		OrderMessage: ctx.FormValue("order_message"),
	}
	orderInfo, orderGoodsList, err := logic.SubmitOrder(uid, orderMap)
	if nil != err {
		// log
	}

	arrApiOrderGoods := make([]*apiIndex.OrderGoods, len(orderGoodsList))
	for idx, item := range orderGoodsList {
		arrApiOrderGoods[idx] = &apiIndex.OrderGoods{OrderGoods: item}
	}
	// 拼装接口数据
	orderData := apiIndex.Order{
		Address:      (*apiIndex.AddressType)(address),
		ShipTimeList: []string{},
		OrderInfo: apiIndex.OrderInfo{
			GoodsList: arrApiOrderGoods,
			Order:     &apiIndex.OrderBase{Order: orderInfo},
		},
		Cancel: genApiCancel(orderInfo),
	}

	orderData.Format()
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
		arrApiOrderGoods := make([]*apiIndex.OrderGoods, len(v["goodsList"].([]*daoSql.OrderGoods)))
		for idx, item := range v["goodsList"].([]*daoSql.OrderGoods) {
			arrApiOrderGoods[idx] = &apiIndex.OrderGoods{OrderGoods: item}
		}

		orderData.List = append(orderData.List, &apiIndex.Order{
			Address: (*apiIndex.AddressType)(v["addressInfo"].(*daoSql.Address)),
			OrderInfo: apiIndex.OrderInfo{
				Order:     &apiIndex.OrderBase{Order: v["order"].(*daoSql.Order)},
				GoodsList: arrApiOrderGoods,
			},
		})
	}

	orderData.Format()
	return util.Success(ctx, orderData)
}

// 订单列表页
func (OrderController) CancelOrder(ctx echo.Context) error {
	// 获取订单列表信息
	uid := uint64(10)
	orderSn := ctx.FormValue("order_sn")
	cancelFlag := util.Atoi(ctx.FormValue("cancel_flag"), 16, false).(uint16)

	err := logic.CancelOrder(uid, orderSn, cancelFlag)
	if nil != err {
		return util.Fail(ctx, 10, err.Error())
	}
	return util.Success(ctx, nil)
}

// 订单列表页
func (OrderController) EvalOrder(ctx echo.Context) error {
	// 获取订单列表信息
	uid := uint64(10)
	orderSn := ctx.FormValue("order_sn")
	stars := util.Atoi(ctx.FormValue("stars"), 8, false).(uint8)
	feedback := ctx.FormValue("feedback")

	err := logic.EvalOrder(uid, orderSn, stars, feedback)
	if nil != err {
		return util.Fail(ctx, 10, err.Error())
	}
	return util.Success(ctx, nil)
}

// 订单列表页
func (OrderController) Detail(ctx echo.Context) error {
	ordeSn := ctx.QueryParam("order_sn")

	// 获取订单列表信息
	// uid, orderSn
	myOrderMap, err := logic.GetOrderDetail(10, ordeSn)
	if nil != err {
		return err
	}
	// 拼装接口数据
	arrApiOrderGoods := make([]*apiIndex.OrderGoods, len(myOrderMap["goodsList"].([]*daoSql.OrderGoods)))
	for idx, item := range myOrderMap["goodsList"].([]*daoSql.OrderGoods) {
		arrApiOrderGoods[idx] = &apiIndex.OrderGoods{OrderGoods: item}
	}
	orderData := &apiIndex.Order{
		Alert:   "",
		Address: (*apiIndex.AddressType)(myOrderMap["addressInfo"].(*daoSql.Address)),
		OrderInfo: apiIndex.OrderInfo{
			Order:     &apiIndex.OrderBase{Order: myOrderMap["order"].(*daoSql.Order)},
			GoodsList: arrApiOrderGoods,
		},
		Cancel: genApiCancel(myOrderMap["order"].(*daoSql.Order)),
	}

	orderData.Format()
	return util.Render(ctx, "order/info", "订单详情", orderData)
}

// 生成订单列表页页的 html 不提供外部接口
func (OrderController) GenOrderListHtml(ctx echo.Context) error {
	return util.Render(ctx, "order/list", "订单列表", map[string]interface{}{})
}

// 获取购物车中商品列表
func getCartGoodsList(ctx echo.Context) (goodsList []*logic.CartInfo, err error) {
	goodsList = []*logic.CartInfo{}

	goodsListStr := ctx.FormValue("goods_list")
	goodsListMap := []map[string]interface{}{}
	err = json.Unmarshal([]byte(goodsListStr), &goodsListMap)

	for _, item := range goodsListMap {
		tmpInfo := &logic.CartInfo{
			GoodsId:  util.MustNum(item["goods_id"], 64, false).(uint64),
			Selected: util.MustNum(item["selected"], 8, false).(uint8),
			GoodsNum: util.MustNum(item["goods_num"], 16, false).(uint16),
		}
		goodsList = append(goodsList, tmpInfo)
	}

	if err != nil {
		// log
		return
	}
	if 0 >= len(goodsList) {
		return goodsList, CartEmpty
	}

	return
}

// 提交订单时新地址(address_id<=0) 会先插入地址表
func fetchAddress(ctx echo.Context) (*daoSql.Address, error) {
	var uid uint64
	uid = 10

	// 获取提交订单时指定的地址
	addressId := util.Atoi(ctx.FormValue("address_id"), 64, false).(uint64)
	if 0 < addressId {
		myAddressMap, err := daoSql.GetAddressListById([]uint64{addressId})
		if nil != err {
			// log
			return &daoSql.Address{}, err
		}
		myAddress, ok := myAddressMap[addressId]
		if !ok || uid != myAddress.MemberId {
			// log
			return &daoSql.Address{}, RecordEmpty
		}
		return myAddress, nil
	}

	// 插入新的地址信息
	trueName := ctx.FormValue("true_name")
	liveArea := ctx.FormValue("live_area")
	address := ctx.FormValue("address")
	mobile := ctx.FormValue("mobile")
	// 显式提取地址信息
	addressInfo := daoSql.UserAddressInfo{
		TrueName: trueName,
		LiveArea: liveArea,
		Address:  address,
		Mobile:   mobile,
	}
	myAddress, err := daoSql.SaveMyAddress(uid, &addressInfo)
	if nil != err {
		// log
		return &daoSql.Address{}, err
	}
	return myAddress, err
}

// 根据订单
func genApiCancel(orderInfo *daoSql.Order) (cancel *apiIndex.Cancel) {
	// 读入配置信息
	envConf, _ := daoConf.EnvConf()

	// 订单取消信息
	cancelInfo := logic.GetCancelInfo(orderInfo)

	cancel = &apiIndex.Cancel{
		CanCancel: cancelInfo.CanCancel,
	}

	if !cancelInfo.CanCancel {
		if 0 < len(envConf.ServiceTel) {
			cancel.CancelTip.Tel = envConf.ServiceTel
		} else {
			cancel.CancelTip.Tel = cancelInfo.CancelTip.Tel
		}
		cancel.CancelTip.Tip = cancelInfo.CancelTip.Tip

	} else {
		cancelReasonList := []*apiIndex.CancelReasonType{}
		for k, v := range cancelInfo.CancelReason {
			tmp := &apiIndex.CancelReasonType{
				Flag:    util.Itoa(k),
				Context: v,
			}
			cancelReasonList = append(cancelReasonList, tmp)
		}
		cancel.CancelReason = cancelReasonList
	}
	return cancel
}
