package api

import (
	daoSql "dao/sql"
	"time"
	"util"
)

// 商品接口
type Goods struct {
	*daoSql.Goods

	Marketprice string
	Price       string
}

func (self *Goods) Format() {
	self.Marketprice = util.GetMoneyStr(int64(self.Goods.Marketprice))
	self.Price = util.GetMoneyStr(int64(self.Goods.Price))
}

// 订单商品接口
type OrderGoods struct {
	*daoSql.OrderGoods

	GoodsMarketprice string
	GoodsPrice       string
}

func (self *OrderGoods) Format() {
	self.GoodsMarketprice = util.GetMoneyStr(int64(self.OrderGoods.GoodsMarketprice))
	self.GoodsPrice = util.GetMoneyStr(int64(self.OrderGoods.GoodsPrice))
}

// 订单商品接口
type OrderBase struct {
	*daoSql.Order

	Amount            string
	MarketOrderAmount string
	OrderAmount       string
	AddTime           string
	ExpectTime        string
	FinishedTime      string
	ShippingTime      string
}

func (self *OrderBase) Format() {
	self.Amount = util.GetMoneyStr(int64(self.Order.Amount))
	self.MarketOrderAmount = util.GetMoneyStr(int64(self.Order.MarketOrderAmount))
	self.OrderAmount = util.GetMoneyStr(int64(self.Order.OrderAmount))

	self.AddTime = time.Unix(int64(self.Order.AddTime), 0).Format("2006-01-02 15:04:05")
	self.ExpectTime = time.Unix(int64(self.Order.ExpectTime), 0).Format("2006-01-02 15:04:05")
	self.FinishedTime = time.Unix(int64(self.Order.FinishedTime), 0).Format("2006-01-02 15:04:05")
	self.ShippingTime = time.Unix(int64(self.Order.ShippingTime), 0).Format("2006-01-02 15:04:05")
}

// 分类信息
type Class struct {
	ClassName string   `json:"className"`
	ClassId   uint64   `json:"-"`
	GoodsList []*Goods `json:"goodsList"`
}

func (self *Class) Format() {
	for _, goodsItem := range self.GoodsList {
		goodsItem.Format()
	}
}

// 订单基本信息
type OrderInfo struct {
	Order     *OrderBase    `json:"order"` // 订单信息
	GoodsList []*OrderGoods `json:"goodsList"`
}

func (self *OrderInfo) Format() {
	self.Order.Format()
	for _, goodsItem := range self.GoodsList {
		goodsItem.Format()
	}
}

// 地址类型
type AddressType daoSql.Address

func (self *AddressType) Format() {
}
