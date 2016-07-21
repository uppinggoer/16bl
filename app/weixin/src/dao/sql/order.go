package sql

import (
	"encoding/json"
	"fmt"
	. "global"
	"math/rand"
	"strconv"
	"time"
)

type Order struct {
	OrderId           uint64 `gorm:"primary_key"`
	OrderSn           string
	StoreId           uint64
	MemberId          uint64
	AddressId         uint64
	ReciverName       string
	ReciverMobile     string
	OrderMessage      string
	Score             uint8 `json:"-"`
	PaySn             string
	CostAmount        uint64 `json:"-"`
	Amount            uint64
	MarketOrderAmount uint64
	CostOrderAmount   uint64 `json:"-"`
	OrderAmount       uint64
	OrderState        uint16
	CancelFlag        uint8
	AddTime           uint64   `json:"-"`
	ExpectTime        uint64   `json:"-"`
	ShippingTime      uint64   `json:"-"`
	FinishedTime      uint64   `json:"-"`
	EvaluationTime    uint64   `json:"-"`
	PaymentTime       uint64   `json:"-"`
	Ext               string   `json:"-"`
	ExtInfo           OrderExt `gorm:"-"`
}
type OrderExt struct {
	GoodsList       []uint64 `json:"goodsList"`
	CancelReason    string   `json:"cancelReason"`
	NotCancelReason string   `json:"notCancelReason"`
	Evaluate        string   `json:"evaluate"`
	AddressInfo     *Address `json:"addressInfo"`
}

const (
	OrderStateNew      = 0  // 新订单，未付款
	OrderStatePay      = 10 // 新订单，已支付
	OrderStateOrder    = 20 // 已接单
	OrderStateSend     = 30 // 已发货
	OrderStateSuccess  = 40 // 已收货，交易成功；已送达
	OrderStateEvaluate = 50 // 已评价
)

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
	sqlRet := dbBuild.Where("member_id = ?", uid).Order("order_id desc").Limit(rn).Find(&orderList)
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
	// 格式化
	json.Unmarshal([]byte(self.Ext), &self.ExtInfo)
}

// 获取订单 sn
func GenOrderSn(storeId, uid uint64) string {
	str := strconv.Itoa(rand.Intn(89) + 10)
	str += fmt.Sprintf("%d", (time.Now().Unix()-1467302400)/60)
	str += fmt.Sprintf("%02d", rand.Intn(90)+9)
	str += fmt.Sprintf("%02d", uid/100)

	return str
}

// 获取支付订单 sn
func GenPaySn(storeId, uid uint64) string {
	// 前面部分主要为了跟踪订单及排序
	str := strconv.Itoa(rand.Intn(89) + 10)
	str += fmt.Sprintf("%010d", time.Now().Unix())
	str += fmt.Sprintf("%03d", time.Now().Nanosecond())
	str += fmt.Sprintf("%03d", uid/1000)

	return str
}
