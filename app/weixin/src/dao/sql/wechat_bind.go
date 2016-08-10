package sql

import (
	"fmt"
	"strings"

	. "global"
	"util"
)

type WechatBind struct {
	Id         uint64 `gorm:"primary_key"`
	MemberId   uint64
	Openid     string
	Nickname   string
	Sex        uint8
	Province   string
	City       string
	Headimgurl string
	UpdatedAt  string
	Ext        string
}

func (this *WechatBind) TableName() string {
	return "wechat_bind"
}

// 用户统计信息
type WeChatInfo struct {
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        uint8  `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Headimgurl string `json:"headimgurl"`
}

func GetUidByOpenId(openId string) (uint64, error) {
	if 0 >= len(openId) {
		return 0, nil
	}

	wechatInfo := WechatBind{}
	sqlRet := DB.Where("openid = ?", openId).First(&wechatInfo)

	if nil != sqlRet.Error {
		// log sqlRet.Error
		return 0, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return 0, RecordEmpty
	}

	return wechatInfo.MemberId, nil
}

func UpdateWeChatBind(uid uint64, weChatInfo *WeChatInfo) error {
	if 0 >= uid {
		// log
		return RecordEmpty
	}

	sqlStr := "INSERT INTO wechat_bind SET openid=" + weChatInfo.OpenId + ", %s ON DUPLICATE KEY UPDATE %s"
	args := make([]interface{}, 8)
	updateSql := make([]string, 8)

	updateSqlId := 0
	// uid
	updateSql[updateSqlId] = "member_id=(?)"
	args[updateSqlId] = uid
	updateSqlId++
	// nick
	updateSql[updateSqlId] = "nickname=(?)"
	args[updateSqlId] = weChatInfo.Nickname
	updateSqlId++
	// sex
	updateSql[updateSqlId] = "province=(?)"
	args[updateSqlId] = weChatInfo.Province
	updateSqlId++
	// city
	updateSql[updateSqlId] = "city=(?)"
	args[updateSqlId] = weChatInfo.City
	updateSqlId++
	// nick
	updateSql[updateSqlId] = "headimgurl=(?)"
	args[updateSqlId] = weChatInfo.Headimgurl
	updateSqlId++

	updateStr := strings.Join(updateSql, ",")
	updateStr = fmt.Sprint(sqlStr, updateStr, updateStr)

	args = util.Merge(args, args)
	rest := DB.Exec(updateStr, args...)
	if nil != rest.Error || 0 >= rest.RowsAffected {
		// log
	}

	return nil
}
