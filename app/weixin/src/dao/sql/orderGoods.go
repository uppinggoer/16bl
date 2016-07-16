package sql

import . "global"

type OrderGoods struct {
	Id               int64 `gorm:"primary_key" json:"-"`
	OrderId          int64 `json:"-"`
	BuyerId          int64 `json:"-"`
	GoodsId          int64
	GoodsName        string
	GoodsNorms       string
	GoodsImage       string
	GoodsUnit        string
	GoodsNum         int64
	GoodsPrice       int
	GoodsMarketprice int
	GoodsCostprice   int `json:"-"`
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func GetOrderGoodsMap(orderIdList []int64) (map[int64][]*OrderGoods, error) {
	if 0 >= len(orderIdList) {
		// log
		return nil, RecordEmpty
	}

	goodsList := []*OrderGoods{}
	sqlRet := DB.Where("order_id in (?)", orderIdList).Find(&goodsList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	goodsIdMap := map[int64][]*OrderGoods{}
	for _, goodsInfo := range goodsList {
		goodsIdMap[goodsInfo.OrderId] = append(goodsIdMap[goodsInfo.OrderId], goodsInfo)
	}

	return goodsIdMap, nil
}
