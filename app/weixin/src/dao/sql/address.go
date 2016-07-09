package sql

import
// "fmt"

. "global"

type Address struct {
	Id        int32 `gorm:"primary_key"`
	MemberId  int32
	TrueName  string
	Gender    int8
	LiveArea  string
	Address   string
	MobPhone  string
	IsDefault int8
	UpdatedAt string `json:"-"`
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param goodsIdList
 * @return map[int64]Goods
 */
func GetAddressListByUid(uid []int64, onlyDefault bool) ([]*Address, error) {
	if 0 >= len(uid) {
		// log
		return nil, RecordEmpty
	}

	addressList := []*Address{}
	// db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	dbBuild := DB
	if onlyDefault {
		dbBuild = dbBuild.Where("is_default = ?", onlyDefault)
	}
	sqlRet := dbBuild.Where("member_id in (?)", uid).Find(&addressList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	return addressList, nil
}
