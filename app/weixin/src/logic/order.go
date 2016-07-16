package logic

import (
	daoSql "dao/sql"
	"strconv"
)

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func genOrderGoods(goodsInfo *daoSql.Goods, goodsNum int64) (*daoSql.OrderGoods, error) {
	orderGoods := daoSql.OrderGoods{
		GoodsId:          goodsInfo.Id,
		GoodsName:        goodsInfo.Name,
		GoodsNorms:       goodsInfo.Norms,
		GoodsImage:       goodsInfo.Image,
		GoodsUnit:        goodsInfo.Unit,
		GoodsNum:         goodsNum,
		GoodsPrice:       int(goodsInfo.Price * goodsNum),
		GoodsMarketprice: int(goodsInfo.Marketprice * goodsNum),
		GoodsCostprice:   int(goodsInfo.Costprice * goodsNum),
	}
	return &orderGoods, nil
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GenOrder(uid int64, goodsList []map[string]string) (*daoSql.Order, error) {
	// 生成订单商品信息
	orderGoodsList := []*daoSql.OrderGoods{}
	for _, goodsInfo := range goodsList {
		goodsId, err := strconv.ParseInt(goodsInfo["goods_id"], 10, 64)
		if nil != err {
			// log
			continue
		}
		goodsNum, _ := strconv.ParseInt(goodsInfo["goods_num"], 10, 64)
		if v, ok := goodsIdMap[goodsId]; ok {
			if goodsNum > int64(v.Storage) {
				goodsNum = int64(v.Storage)
			}
			orderGoodsInfo, _ := genOrderGoods(goodsIdMap[goodsId], goodsNum)
			orderGoodsList = append(orderGoodsList, orderGoodsInfo)
		} else {
			// log
			continue
		}
	}

	// 生成订单主体信息
	var amount = 0
	for _, item := range orderGoodsList {
		amount += item.GoodsPrice
	}
	order := &daoSql.Order{
		CostAmount:      amount,
		Amount:          amount,
		CostOrderAmount: 0,
		OrderAmount:     amount,
		OrderState:      1,
	}

	// 过滤参数
	order.Filter()
	return order, nil
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func SubmitOrder(uid int64, goodsList []map[string]string) (*daoSql.Order, error) {
	var amount, marketAmount, costAmount = 0
	for _, item := range goodsInfo {
		amount += item.GoodsPrice
		costAmount += item.Costprice
		marketAmount += item.Marketprice
	}
	order := &daoSql.Order{
		MemberId:        uid,
		Amount:          amount,
		CostOrderAmount: 0,
		OrderAmount:     amount,
		OrderState:      1,
	}

	// UserOrderTimes  int64  `json:"-"`
	// AddressId       int32  `json:"-"`
	// OrderTime       int64  `json:"-"`
	// OrderTimeStr    string `gorm:"-"`
	// ExpectTime      int64  `json:"-"`
	// ExpectTimeStr   string `gorm:"-"`
	// ConfirmTime     int64  `json:"-"`
	// ConfirmTimeStr  string `gorm:"-"`
	// FinishedTime    int64  `json:"-"`
	// FinishedTimeStr string `gorm:"-"`
	// PaySn           int64  `json:"-"`
	// PaymentTime     int64  `json:"-"`
	// PaymentTimeStr  string `gorm:"-"`
	// CostAmount      int
	// Amount          int
	// CostOrderAmount int
	// OrderAmount     int
	// RefundAmount    int64
	// OrderState      int16
	// CancelFlag      int8
	// Ext             string `json:"-"`
	// ExtInfo         string `gorm:"-"`

	// 加事务
	tx := model.DB.Begin()
	// 插入订单

	// 减少 goodsStorage
	// 插入 order_goods
	// 更新 goodsStatis userStatis
	// 提交事务 / 回滚事务   tx.Rollback()

	// 过滤订单参数
	order.Filter()
	return order, nil
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param uid
 * @param baseId
 * @param rn
 * @return
 * 		[]map{
 * 			"order"
 *			"goodsList"
 *		 	"addressInfo"
 *		}
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GetMyOrderList(uid, baseId, rn int) (orderDetailList []map[string]interface{}, hasMore bool, err error) {
	if 0 >= rn {
		rn = 20
	}
	orderDetailList = []map[string]interface{}{}
	hasMore = true

	// uid 订单列表
	orderList, err := daoSql.GetListById(uid, baseId, rn+1)
	if nil != err {
		return
	}
	// orderIdList
	orderIdList := []int64{}
	for _, v := range orderList {
		orderIdList = append(orderIdList, v.OrderId)
	}

	// orderGoodsMap
	orderGoodsMap, err := daoSql.GetOrderGoodsMap(orderIdList)
	if nil != err {
		return
	}

	// addressIdList
	addressIdList := []int32{}
	for _, v := range orderList {
		addressIdList = append(addressIdList, v.AddressId)
	}
	// addressIdMap
	addressIdMap, err := daoSql.GetAddressListById(addressIdList)
	if nil != err {
		return
	}

	for _, v := range orderList {
		orderDetailList = append(orderDetailList, map[string]interface{}{
			"order":       v,
			"goodsList":   orderGoodsMap[v.OrderId],
			"addressInfo": addressIdMap[v.AddressId],
		})
	}

	return
}
