package sql

import (
	// "fmt"
	. "global"
)

type Goods struct {
	Id           int64 `gorm:"primary_key"`
	Name         string
	StoreId      int64
	StoreName    string
	Unit         string
	MinOrdernum  int64
	MaxOrdernum  int64
	Marketprice  int64
	Price        int64
	Costprice    int64
	Salenum      int64
	MonthSalenum int64
	WeekSalenum  int64
	State        int8
	Image        string
	Desc         string
	StcId        int64
	Sort         int64
	Barcode      string
	Storage      int32
	Addtime      int64
	Edittime     int64  `json:"-"`
	OpUser       string `json:"-"`
	Mtime        string `json:"-"`
}

func (this *Goods) TableName() string {
	return "goods"
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func GetGoodsListById(goodsIdList []int64) (map[int64]*Goods, error) {
	if 0 >= len(goodsIdList) {
		// log
		return nil, RecordEmpty
	}

	goodsList := []*Goods{}
	sqlRet := DB.Where("id in (?)", goodsIdList).Find(&goodsList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	goodsIdMap := map[int64]*Goods{}
	for _, goodsInfo := range goodsList {
		goodsIdMap[goodsInfo.Id] = goodsInfo
	}

	return goodsIdMap, nil
}
