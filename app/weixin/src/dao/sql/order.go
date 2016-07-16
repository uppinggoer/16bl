package sql

import (
	. "global"
	"time"
)

type Order struct {
	OrderId           int64 `gorm:"primary_key"`
	OrderSn           string
	StoreId           int64
	MemberId          int64  `json:"-"`
	AddressId         int32  `json:"-"`
	AddTime           int64  `json:"-"`
	AddTimeStr        string `gorm:"-"`
	ExpectTime        int64  `json:"-"`
	ExpectTimeStr     string `gorm:"-"`
	ShippingTime      int64  `json:"-"`
	ShippingTimeStr   string `gorm:"-"`
	FinishedTime      int64  `json:"-"`
	FinishedTimeStr   string `gorm:"-"`
	EvaluationTime    int64  `json:"-"`
	EvaluationTimeStr string `gorm:"-"`
	PaySn             int64  `json:"-"`
	PaymentTime       int64  `json:"-"`
	PaymentTimeStr    string `gorm:"-"`
	ReciverName       string
	ReciverMobile     string
	OrderMessage      string
	DeliveryScore     int8
	CostAmount        int
	Amount            int64
	MarketOrderAmount int64
	CostOrderAmount   int64
	OrderAmount       int64
	OrderState        int32
	CancelFlag        int8
	Ext               string `json:"-"`
	ExtInfo           string `gorm:"-"`
}

func (this *Order) TableName() string {
	return "order"
}

// 根据 用户 id 获取订单列表
// @param baseId -1 表示最大
func GetListById(uid, baseId, rn int) ([]*Order, error) {
	// 默认取 10 条
	if 0 >= rn {
		rn = 20
	}

	// 订单列表
	if 0 >= uid {
		return []*Order{}, nil
	}

	orderList := []*Order{}

	dbBuild := DB
	if 0 < baseId {
		dbBuild = dbBuild.Where("order_id < ?", baseId)
	}
	sqlRet := dbBuild.Where("buyer_id = ?", uid).Order("order_id desc").Limit(rn).Find(&orderList)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	for _, v := range orderList {
		v.Filter()
	}

	return orderList, nil
}

func (self *Order) Filter() {
	self.OrderTimeStr = time.Unix(self.OrderTime, 0).Format("2006-01-02 15:04:05")
	self.ExpectTimeStr = time.Unix(self.ExpectTime, 0).Format("2006-01-02 15:04:05")
	self.ConfirmTimeStr = time.Unix(self.ConfirmTime, 0).Format("2006-01-02 15:04:05")
	self.FinishedTimeStr = time.Unix(self.FinishedTime, 0).Format("2006-01-02 15:04:05")
	self.PaymentTimeStr = time.Unix(self.PaymentTime, 0).Format("2006-01-02 15:04:05")
}
