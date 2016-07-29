package sql

import (
	// "fmt"
	. "global"
)

type GoodsClass struct {
	ClassId   uint64 `gorm:"primary_key"`
	ClassName string
	Level     uint8
	ParentId  int64
	StcState  uint8
	StoreId   int64
	Sort      uint8
	StcImage  string
}

func (this *GoodsClass) TableName() string {
	return "goods_class"
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param classIdList
 * @return map[int64]Goods
 */
func GetClassListById(classIdList []uint64) (map[uint64]*GoodsClass, error) {
	if 0 >= len(classIdList) {
		// log
		return nil, RecordEmpty
	}

	classList := []*GoodsClass{}
	sqlRet := DB.Where("class_id in (?)", classIdList).Find(&classList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	classMap := map[uint64]*GoodsClass{}
	for _, classInfo := range classList {
		classMap[classInfo.ClassId] = classInfo
	}

	return classMap, nil
}
