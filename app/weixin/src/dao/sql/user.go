package sql

import
// "fmt"
. "global"

type Member struct {
	MemberId     uint64 `gorm:"primary_key"`
	Name         string
	Passwd       string
	Truename     string
	Avatar       string
	Mobile       string
	Sex          uint8
	Birthday     uint64
	RegisterTime string
	Ext          string
}

func (this *Member) TableName() string {
	return "member"
}

// 根据 用户 id 获取订单列表
// @param baseId -1 表示最大
func GetInfoByUid(uid uint64) (*Member, error) {
	// 订单列表
	if 0 >= uid {
		return &Member{}, RecordEmpty
	}

	userInfo := Member{}
	sqlRet := DB.Where("member_id = ?", uid).First(&userInfo)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}

	return &userInfo, nil
}

// 根据 用户 id 获取订单列表
// @param baseId -1 表示最大
func AddNewUser(uid uint64) (*Member, error) {
	// 订单列表
	if 0 >= uid {
		return &Member{}, RecordEmpty
	}

	userInfo := Member{}
	sqlRet := DB.Where("member_id = ?", uid).First(&userInfo)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}

	return &userInfo, nil
}
