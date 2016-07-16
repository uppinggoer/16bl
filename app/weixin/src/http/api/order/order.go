package index

import (
	daoSql "dao/sql"
)

// 购物车接口
type Order struct {
	Alert        string          `json:"alert"` // 提示文案
	Address      *daoSql.Address `json:"address"`
	ShipTimeList []string        `json:"shipTimeList"`
	OrderBase    `json:"orderInfo"`
}

type OrderList struct {
	Base    string   `json:"base"`
	HasMore bool     `json:"hasMore"`
	List    []*Order `json:"list"`
}

type OrderBase struct {
	Order     *daoSql.Order        `json:"order"` // 订单信息
	GoodsList []*daoSql.OrderGoods `json:"goodsList"`
}
