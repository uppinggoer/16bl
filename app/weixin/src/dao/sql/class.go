package sql

import (
	// "fmt"
	. "global"
)

type GoodsClass struct {
	Id        int64 `gorm:"primary_key"`
	Name      string
	ParentId  int64
	Sort      int8
	OpUser    string `json:"-"`
	UpdatedAt string `json:"-"`
}

func (this *GoodsClass) TableName() string {
	return "goods_class"
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param classIdList
 * @return map[int64]Goods
 */
func GetClassListById(classIdList []int64) (map[int64]*GoodsClass, error) {
	if 0 >= len(classIdList) {
		// log
		return nil, RecordEmpty
	}

	classList := []*GoodsClass{}
	sqlRet := DB.Where("id in (?)", classIdList).Find(&classList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	classMap := map[int64]*GoodsClass{}
	for _, classInfo := range classList {
		classMap[classInfo.Id] = classInfo
	}

	return classMap, nil
}
