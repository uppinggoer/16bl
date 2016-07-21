package sql

import . "global"

type OrderGoods struct {
	// GoodsCostprice   int32
	Id               uint64 `gorm:"primary_key" json:"-"`
	OrderId          uint64 `json:"-"`
	MemberId         uint64 `json:"-"`
	StoreId          uint64
	GoodsId          uint64
	ClassId          uint64
	GoodsName        string
	GoodsNorms       string
	GoodsImage       string
	GoodsUnit        string
	GoodsNum         uint16
	GoodsPrice       uint64
	GoodsMarketprice uint64
	GoodsCostprice   uint64 `json:"-"`
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func GetOrderGoodsMap(orderIdList []uint64) (map[uint64][]*OrderGoods, error) {
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

	goodsIdMap := map[uint64][]*OrderGoods{}
	for _, goodsInfo := range goodsList {
		goodsIdMap[goodsInfo.OrderId] = append(goodsIdMap[goodsInfo.OrderId], goodsInfo)
	}

	return goodsIdMap, nil
}
