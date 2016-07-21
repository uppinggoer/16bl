package sql

import (
	// "fmt"
	. "global"
	"strings"
	"time"
)

type UserStatis struct {
	MemberId     uint64 `gorm:"primary_key"`
	FirstOrderId uint64
	LastOrderId  uint64
	LoginNum     uint64
	LoginIp      string
	LoginTime    string
	OrderNum     uint64
	Ext          string
}

func (this *UserStatis) TableName() string {
	return "member_statis"
}

// 用户统计信息
type UserStatisInfo struct {
	OrderId   uint64
	LoginIp   string
	LoginTime int64
}

/**
 * @abstract 根据id 列表获取商品信息
 * @param uid
 * @return
 */
func UpdateUserStatis(uid uint64, userStatis *UserStatisInfo) error {
	if 0 >= uid {
		// log
		return RecordEmpty
	}

	sqlStr := "INSERT INTO member_statis SET member_id=(?),last_order_id=(?)" +
		",login_num=1,login_ip=(?),order_num=(?),login_time=(?)" +
		" ON DUPLICATE KEY UPDATE "
	args := []interface{}{uid, userStatis.OrderId, userStatis.LoginIp}
	// 如果是订单则插入 order_num = 1
	if 0 < userStatis.OrderId {
		args = append(args, 1)
	}

	updateSql := []string{}
	// 格式化时间 拼装 sql 及 args
	var logTime string
	if 0 < userStatis.LoginTime {
		tm := time.Unix(userStatis.LoginTime, 0)
		logTime = tm.Format("2006-01-02 15:04:05")
		args = append(args, logTime)

		updateSql = append(updateSql, "login_time=(?)")
		args = append(args, logTime)
		updateSql = append(updateSql, "login_num=login_num+(?)")
		args = append(args, 1)
	} else {
		logTime = time.Now().Format("2006-01-02 15:04:05")
		args = append(args, logTime)
	}

	// 拼装 IP sql 及 args
	if 0 < len(userStatis.LoginIp) {
		sqlStr += "login_ip=(?) "
		args = append(args, userStatis.LoginIp)
	}

	// 拼装 IP sql 及 args
	if 0 < userStatis.OrderId {
		updateSql = append(updateSql, "last_order_id=(?)")
		args = append(args, userStatis.OrderId)
	}

	// 拼装 IP sql 及 args
	if 0 < userStatis.OrderId {
		updateSql = append(updateSql, "last_order_id=(?)")
		args = append(args, userStatis.OrderId)
		updateSql = append(updateSql, "order_num=order_num+(?)")
		args = append(args, 1)
	}

	sqlStr = sqlStr + strings.Join(updateSql, ",")

	rest := DB.Exec(sqlStr, args...)
	if nil != rest.Error || 0 >= rest.RowsAffected {
		// log
	}

	// // 更新 first_order_id
	// if 0 < userStatis.OrderId {
	// 	sqlStr = "UPDATE member_statis SET first_order_id=(?) WHERE member_id=(?) and first_order_id<=0"

	// 	rest := DB.Exec(sqlStr, userStatis.OrderId, uid)
	// 	if nil != rest.Error || 0 >= rest.RowsAffected {
	// 		// log
	// 	}
	// }

	return nil
}
