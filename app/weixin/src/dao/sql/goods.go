package sql

import
// "fmt"

. "global"

type Goods struct {
	Id           uint64 `gorm:"primary_key"`
	Name         string
	Norms        string
	StoreId      uint64
	StoreName    string
	Unit         string
	MinOrdernum  uint64
	MaxOrdernum  uint64
	Marketprice  uint16 `json:"-"`
	Price        uint16 `json:"-"`
	Costprice    uint16 `json:"-"`
	Storage      uint16
	GoodsSalenum uint64
	State        uint8
	Image        string
	Desc         string
	ClassId      uint64
	Sort         uint8
	Barcode      string
}

const (
	GOODS_OFF = 0
	GOODS_ON  = 1
	GOODS_DEL = 254
)

func (this *Goods) TableName() string {
	return "goods"
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func GetGoodsListById(goodsIdList []uint64) (map[uint64]*Goods, error) {
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

	goodsIdMap := map[uint64]*Goods{}
	for _, goodsInfo := range goodsList {
		goodsIdMap[goodsInfo.Id] = goodsInfo
	}
	return goodsIdMap, nil
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param cond map[string]string 条件
 * @return map[int64]Goods
 */
func GetAllGoods(cond *map[string]string) ([]*Goods, error) {
	goodsList := []*Goods{}

	dbBuild := DB
	// 未做校验
	for k, v := range *cond {
		dbBuild = dbBuild.Where(k+"=?", v)
	}
	sqlRet := dbBuild.Find(&goodsList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}
	return goodsList, nil
}
