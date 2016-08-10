package api

import (
	daoConf "dao/conf"
)

// 购物车接口
type Order struct {
	Alert        string       `json:"alert"` // 提示文案
	Address      *AddressType `json:"address"`
	ShipTimeList []string     `json:"shipTimeList"`
	OrderInfo    `json:"orderInfo"`
	Cancel       *Cancel
	AddressList  []*AddressType `json:"addressList"`
	WeChatPay    *WeChatPay
}

func (self *Order) Format() {
	self.OrderInfo.Format()
}

type OrderList struct {
	Base    string   `json:"base"`
	HasMore bool     `json:"hasMore"`
	List    []*Order `json:"list"`
}

func (self *OrderList) Format() {
	for _, v := range self.List {
		v.Format()
	}
}

type Cancel struct {
	CanCancel bool
	CancelTip struct {
		Tel string
		Tip string
	}
	CancelReason []*CancelReasonType
}
type CancelReasonType struct {
	Flag    string
	Context string
}

type WeChatPay struct {
	daoConf.WeChatJs
}
