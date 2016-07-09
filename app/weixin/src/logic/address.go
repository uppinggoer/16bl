package logic

import (
	daoSql "dao/sql"
)

/**
 * @abstract 返回购物车信息
 * @param uid
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GetDefaultAddress(uid int64) (*daoSql.Address, error) {
	addressList, err := daoSql.GetAddressListByUid([]int64{uid}, true)
	if nil != err {
		return &daoSql.Address{}, err
	}

	return addressList[0], nil
}
