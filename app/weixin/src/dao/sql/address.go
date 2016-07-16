package sql

import . "global"

type Address struct {
	Id        int32 `gorm:"primary_key"`
	MemberId  int32
	TrueName  string
	Gender    int8
	LiveArea  string
	Address   string
	Mobile    string
	IsDefault int8
}

/**
 * @abstract 根据id 列表获取地址信息
 * @param addressIdList
 * @return map[int64]Address
 */
func GetAddressListById(addressIdList []int32) (map[int32]*Address, error) {
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

	addressIdMap := map[int32]*Address{}
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
func GetAddressListByUid(uid int64, onlyDefault bool) ([]*Address, error) {
	if 0 >= len(uid) {
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

/**
 * @abstract 根据uid 列表获取地址信息
 * @param uid
 * @param addressInfo 需要插入的信息， key值见 insertFields
 * @return address
 */
func SaveMyAddress(uid int64, addressInfo map[string]string) (*Address, error) {
	// insertFields = []string{"true_name", "gender", "live_area", "address", "mobile"}
	if 0 >= uid {
		// log
		return &Address{}, RecordEmpty
	}

	trueName, _ = addressInfo["true_name"]
	gender, _ = addressInfo["gender"]
	liveArea, _ = addressInfo["live_area"]
	address, _ = addressInfo["address"]
	mobile, _ = addressInfo["mobile"]

	address := Address{
		MemberId:  uid,
		TrueName:  trueName,
		Gender:    gender,
		LiveArea:  liveArea,
		Address:   address,
		Mobile:    mobile,
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
