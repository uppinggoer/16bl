package sql

import (
	// "fmt"
	. "global"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderLog struct {
	Id            uint64 `gorm:"primary_key"`
	OrderId       uint64
	StoreId       uint64
	LogAmount     uint64
	LogMsg        string
	LogUserId     uint64
	LogUser       string
	LogOrderstate int16
	LogTime       string
	Ext           string
}

func (this *OrderLog) TableName() string {
	return "order_log"
}

func InsertOrderLog(db *gorm.DB, orderId uint64, msg string, logTime uint64, orderState int16, orderAmount uint64) error {
	if 0 >= orderId {
		// log
		return RecordEmpty
	}

	var logTimeStr string
	if 0 >= logTime {
		logTimeStr = time.Now().Format("2006-01-02 15:04:05")
	} else {
		logTimeStr = time.Unix(int64(logTime), 0).Format("2006-01-02 15:04:05")
	}

	log := OrderLog{
		OrderId:       orderId,
		LogAmount:     orderAmount,
		LogMsg:        msg,
		LogTime:       logTimeStr,
		LogUserId:     0,
		LogUser:       "系统",
		LogOrderstate: orderState,
	}

	db.Create(&log)
	if 0 >= log.Id {
		return InsertError
	}
	return nil
}
