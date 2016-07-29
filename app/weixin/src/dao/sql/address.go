package sql

import . "global"

type Address struct {
	Id        uint64 `gorm:"primary_key"`
	MemberId  uint64
	TrueName  string
	Gender    uint8
	LiveArea  string
	Address   string
	Mobile    string
	IsDefault uint8
}

func (this *Address) TableName() string {
	return "address"
}

/**
 * @abstract 根据id 列表获取地址信息
 * @param addressIdList
 * @return map[int64]Address
 */
func GetAddressListById(addressIdList []uint64) (map[uint64]*Address, error) {
	if 0 >= len(addressIdList) {
		// log
		return nil, RecordEmpty
	}

	addressList := []*Address{}
	sqlRet := DB.Where("id in (?)", addressIdList).Find(&addressList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	addressIdMap := map[uint64]*Address{}
	for _, addressInfo := range addressList {
		addressIdMap[addressInfo.Id] = addressInfo
	}

	return addressIdMap, nil
}

/**
 * @abstract 根据uid 列表获取地址信息
 * @param uid
 * @return map[int64]Goods
 */
func GetAddressListByUid(uid uint64, onlyDefault bool) ([]*Address, error) {
	if 0 >= uid {
		// log
		return nil, RecordEmpty
	}

	addressList := []*Address{}
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

// 地址详情
type UserAddressInfo struct {
	TrueName  string
	Gender    uint8
	LiveArea  string
	Address   string
	Mobile    string
	IsDefault uint8
}

/**
 * @abstract 根据uid 列表获取地址信息
 * @param uid
 * @param addressInfo 需要插入的信息， key值见 insertFields
 * @return address
 */
func SaveMyAddress(uid uint64, addressInfo *UserAddressInfo) (*Address, error) {
	// insertFields = []string{"true_name", "gender", "live_area", "address", "mobile"}
	if 0 >= uid {
		// log
		return &Address{}, RecordEmpty
	}

	address := Address{
		MemberId:  uid,
		TrueName:  addressInfo.TrueName,
		Gender:    addressInfo.Gender,
		LiveArea:  addressInfo.LiveArea,
		Address:   addressInfo.Address,
		Mobile:    addressInfo.Mobile,
		IsDefault: 1,
	}
	// 插入地址信息
	DB.Create(&address)

	if 0 >= address.Id {
		// log
		return &address, RecordEmpty
	}

	return &address, nil
}
