package logic

import (
	daoSql "dao/sql"
	"errors"
	. "global"

	"encoding/json"
	"time"
)

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func genOrderGoods(uid uint64, goodsInfo *daoSql.Goods, goodsNum uint16) (*daoSql.OrderGoods, error) {
	orderGoods := daoSql.OrderGoods{
		GoodsId:          goodsInfo.Id,
		MemberId:         uid,
		GoodsName:        goodsInfo.Name,
		GoodsNorms:       goodsInfo.Norms,
		GoodsImage:       goodsInfo.Image,
		GoodsUnit:        goodsInfo.Unit,
		GoodsNum:         goodsNum,
		GoodsPrice:       uint64(goodsInfo.Price * goodsNum),
		GoodsMarketprice: uint64(goodsInfo.Marketprice * goodsNum),
		GoodsCostprice:   uint64(goodsInfo.Costprice * goodsNum),
	}
	return &orderGoods, nil
}

type OrderMap struct {
	Address      *daoSql.Address
	GoodsIdMap   map[uint64]*daoSql.Goods
	GoodsList    []*CartInfo
	ExceptTime   int64
	OrderMessage string
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param uid
 * @param orderMap 具体字段参见
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GenOrder(uid uint64, orderMap OrderMap) (*daoSql.Order, []*daoSql.OrderGoods, error) {
	// 生成订单商品信息
	var amount, marketAmount, costAmount uint64
	orderGoodsList := []*daoSql.OrderGoods{}
	orderGoodsIdList := []uint64{}
	for _, goodsInfo := range orderMap.GoodsList {
		if _, ok := orderMap.GoodsIdMap[goodsInfo.GoodsId]; !ok {
			// log
			break
		}

		// 生成 order_goods
		orderGoodsInfo, _ := genOrderGoods(uid, orderMap.GoodsIdMap[goodsInfo.GoodsId], goodsInfo.GoodsNum)
		// 生成 order_goods 列表
		orderGoodsList = append(orderGoodsList, orderGoodsInfo)
		// 收集订单中 goods_id
		orderGoodsIdList = append(orderGoodsIdList, goodsInfo.GoodsId)

		amount += orderGoodsInfo.GoodsPrice
		costAmount += orderGoodsInfo.GoodsCostprice
		marketAmount += orderGoodsInfo.GoodsMarketprice
	}
	// 生成订单主体信息
	order := &daoSql.Order{
		MemberId:          uid,
		AddressId:         orderMap.Address.Id,
		AddTime:           uint64(time.Now().Unix()),
		ExpectTime:        uint64(orderMap.ExceptTime),
		ReciverName:       orderMap.Address.TrueName,
		ReciverMobile:     orderMap.Address.Mobile,
		CostAmount:        costAmount,
		Amount:            amount,
		CostOrderAmount:   costAmount,
		MarketOrderAmount: marketAmount,
		OrderAmount:       amount,
		OrderState:        1,
		CancelFlag:        0,
		OrderMessage:      orderMap.OrderMessage,
		ExtInfo: daoSql.OrderExt{
			GoodsList:   orderGoodsIdList,
			AddressInfo: orderMap.Address,
		},
	}

	return order, orderGoodsList, nil
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func SubmitOrder(uid uint64, orderMap OrderMap) (orderInfo *daoSql.Order, orderGoodsList []*daoSql.OrderGoods, err error) {
	orderInfo, orderGoodsList, err = GenOrder(uid, orderMap)
	if nil != err {
		return
	}

	// 补充其它字段
	orderInfo.OrderSn = daoSql.GenOrderSn(0, uid)
	orderInfo.PaySn = daoSql.GenPaySn(0, uid)
	// GenPaySn
	ext, _ := json.Marshal(orderInfo.ExtInfo)
	orderInfo.Ext = string(ext)

SUBMIT:
	for i := 0; i < 3; i++ {
		// 加事务
		tx := daoSql.DB.Begin()
		// 插入订单
		tx.Create(orderInfo)
		if 0 >= orderInfo.OrderId {
			tx.Rollback()
			// 重新提交流程
			continue SUBMIT
		}

		for _, v := range orderGoodsList {
			v.OrderId = orderInfo.OrderId
			// 插入order_goods
			tx.Create(v)
			if 0 >= v.Id {
				tx.Rollback()
				// 重新提交流程
				continue SUBMIT
			}

			// 减少 goodsStorage
			// DB.where("id = ?", v.GoodsId).
			rest := tx.Exec("UPDATE goods SET storage=storage-(?),goods_salenum=goods_salenum+(?) WHERE id = (?)", v.GoodsNum, v.GoodsNum, v.GoodsId)
			if nil != rest.Error || 0 >= rest.RowsAffected {
				tx.Rollback()
				// 重新提交流程
				continue SUBMIT
			}
		}

		// 插入 log
		err = daoSql.InsertOrderLog(tx, orderInfo.OrderId, "新订单", orderInfo.AddTime, daoSql.OrderStateNew, orderInfo.OrderAmount)
		if nil != err {
			tx.Rollback()
			continue SUBMIT
		}

		// 提交事务 / 回滚事务
		tx.Commit()

		// 另起线程更新统计信息 goodsStatis userStatis
		go func(orderMap OrderMap, orderInfo *daoSql.Order) {
			goodsInfo := make([]*daoSql.GoodsOrderInfo, len(orderMap.GoodsList))
			for i, v := range orderMap.GoodsList {
				if 1 == v.Selected && 0 < v.GoodsNum {
					goodsInfo[i] = &daoSql.GoodsOrderInfo{
						GoodsId:  v.GoodsId,
						GoodsNum: v.GoodsNum,
					}
				}
			}

			daoSql.UpdateGoodsStatis(goodsInfo)
			userStatis := daoSql.UserStatisInfo{
				OrderId: orderInfo.OrderId,
			}
			daoSql.UpdateUserStatis(orderInfo.MemberId, &userStatis)
		}(orderMap, orderInfo)

		// 过滤订单参数  准备输出
		orderInfo.Filter()
		return
	}

	return
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
	orderIdList := []uint64{}
	for _, v := range orderList {
		orderIdList = append(orderIdList, v.OrderId)
	}

	// orderGoodsMap
	orderGoodsMap, err := daoSql.GetOrderGoodsMap(orderIdList)
	if nil != err {
		return
	}

	// addressIdList
	addressIdList := []uint64{}
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

/**
 * @abstract 获取订单详情
 * @param uid
 * @param baseId
 * @return
 * 		[]map{
 * 			"order"
 *			"goodsList"
 *		 	"addressInfo"
 *		}
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GetOrderDetail(uid uint64, orderSn string) (orderDetail map[string]interface{}, err error) {
	orderDetail = map[string]interface{}{}

	// uid 订单列表
	orderInfo, err := daoSql.GetOrderByOrderSn(orderSn)
	if nil != err {
		return
	}
	if orderInfo.MemberId != uid {
		return nil, NotYourOrder
	}

	// orderGoodsMap
	orderGoodsMap, err := daoSql.GetOrderGoodsMap([]uint64{orderInfo.OrderId})

	// addressIdMap
	addressIdMap, err := daoSql.GetAddressListById([]uint64{orderInfo.AddressId})
	if nil != err {
		addressIdMap = map[uint64]*daoSql.Address{
			orderInfo.AddressId: &daoSql.Address{},
		}
		err = nil
	}

	// 订单详情
	orderDetail = map[string]interface{}{
		"order":       orderInfo,
		"goodsList":   orderGoodsMap[orderInfo.OrderId],
		"addressInfo": addressIdMap[orderInfo.AddressId],
	}

	return
}

/**
 * @abstract 评介订单
 * @param uid
 * @param orderSn
 * @param score
 * @param feedback 反馈的内容
 * @return
 *   err
 */
func EvalOrder(uid uint64, orderSn string, score uint8, feedback string) (err error) {
	// 取出 uid,orderSn
	orderInfo, err := daoSql.GetOrderByOrderSn(orderSn)
	if nil != err {
		return
	}
	if orderInfo.MemberId != uid {
		return NotYourOrder
	}

	if orderInfo.OrderState > daoSql.OrderStateSuccess {
		return errors.New("订单未完成，不能评价")
	}
	if orderInfo.OrderState >= daoSql.OrderStateEvaluate {
		return errors.New("订单不能重复评价")
	}

	// update
	orderInfo.Score = score
	orderInfo.OrderState = daoSql.OrderStateEvaluate
	orderInfo.ExtInfo.Evaluate = feedback
	err = orderInfo.Save()
	if nil != err {
		return
	}

	// insert order_log
	err = daoSql.InsertOrderLog(daoSql.DB, orderInfo.OrderId, "订单评论", 0, daoSql.OrderStateEvaluate, orderInfo.OrderAmount)
	if nil != err {
		// log
		return err
	}

	return nil
}

/**
 * @abstract 取消订单
 * @param uid
 * @param orderSn
 * @param cancelFlag 取消原因
 * @return
 *   err
 */
func CancelOrder(uid uint64, orderSn string, cancelFlag uint16) (err error) {
	// 取出 uid,orderSn
	orderInfo, err := daoSql.GetOrderByOrderSn(orderSn)
	if nil != err {
		return
	}
	if orderInfo.MemberId != uid {
		return NotYourOrder
	}

	if !orderInfo.CancelOrder() {
		return errors.New("订单未完成，不能评价")
	}

	// update
	orderInfo.OrderState = cancelFlag
	orderInfo.ExtInfo.CancelReason = daoSql.CancelReason[cancelFlag]
	err = orderInfo.Save()
	if nil != err {
		return
	}

	// insert order_log
	err = daoSql.InsertOrderLog(daoSql.DB, orderInfo.OrderId, "订单取消", 0, daoSql.OrderStateEvaluate, orderInfo.OrderAmount)
	if nil != err {
		// log
		return err
	}

	return nil
}

type OrderCancel struct {
	CanCancel bool
	CancelTip struct {
		Tel string
		Tip string
	}
	CancelReason map[uint16]string
}

/**
 * @abstract 判断订单是否被用户取消
 * @param orderInfo
 * @return
 * 	 OrderCancel
 *   err
 */
func GetCancelInfo(orderInfo *daoSql.Order) (cancelInfo *OrderCancel) {
	cancelInfo = &OrderCancel{
		CanCancel:    true,
		CancelReason: map[uint16]string{},
	}

	cancelInfo.CanCancel = orderInfo.CancelOrder()
	if cancelInfo.CanCancel {
		for k, v := range daoSql.CancelReason {
			if 200 < k {
				cancelInfo.CancelReason[k] = v
			}
		}
	} else {
		cancelInfo.CancelTip.Tel = "18211121906"
		cancelInfo.CancelTip.Tip = "订单已经确认，请联系客服"
	}
	return
}
