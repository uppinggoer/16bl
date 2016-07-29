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
	CancelFlag        uint16
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

	OrderNotCancel  = 0 // 未取消
	OrderHasCancel  = 1 // 已经取消
	OrderFailCancel = 2 // 取消失败
)

var CancelReason = map[uint16]string{
	100: "不符合活动参与条件",
	101: "订单超出配送范围",
	102: "用户联系运营退单",
	103: "暂时无法配送",
	104: "商品售罄",
	105: "商品价格/库存问题",
	106: "订单地址不准确",
	107: "无法联系您",

	200: "商品选错了，重新选择",
	201: "配送信息填写错误",
	202: "改变主意不想要了",
	203: "无法享受优惠",
	204: "无法及时配送",
	205: "商品价格变动/缺货",
	206: "其他",
}

func (this *Order) TableName() string {
	return "order"
}

// 根据 订单编号获取订单信息
func GetOrderByOrderSn(orderSn string) (orderInfo *Order, err error) {
	orderInfo = &Order{}
	if 0 >= len(orderSn) {
		return
	}

	sqlRet := DB.Where("order_sn = ?", orderSn).Find(orderInfo)
	if nil != sqlRet.Error {
		// log sqlRet.Error
		return nil, RecordError
	}
	if 0 >= sqlRet.RowsAffected {
		// log
		return nil, RecordEmpty
	}

	orderInfo.Filter()

	return
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
