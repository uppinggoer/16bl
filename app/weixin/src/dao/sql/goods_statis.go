package sql

import
// "fmt"

(
	. "global"
	"time"
)

type GoodsStatis struct {
	Id           uint64 `gorm:"primary_key"`
	GoodsId      uint64
	Month        uint64
	StoreId      uint64
	MonthSalenum uint64
	WeekSalenum  uint64
}

func (this *GoodsStatis) TableName() string {
	return "goods_statis"
}

// 购物车详情
type GoodsOrderInfo struct {
	GoodsId  uint64 `json:"goods_id"`
	GoodsNum uint16 `json:"goods_num"`
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func UpdateGoodsStatis(goodsList []*GoodsOrderInfo) error {
	if 0 >= len(goodsList) {
		// log
		return RecordEmpty
	}

	month := time.Now().Format("20060102")
	sqlStr := "INSERT INTO goods_statis SET goods_id=(?),month=(?),month_salenum=(?) " +
		" ON DUPLICATE KEY UPDATE month_salenum=month_salenum+(?)"
	for _, v := range goodsList {
		rest := DB.Exec(sqlStr, v.GoodsId, month, v.GoodsNum, v.GoodsNum)
		if nil != rest.Error || 0 >= rest.RowsAffected {
			// log
		}
	}

	return nil
}
