package index

import (
	daoSql "dao/sql"
)

// 购物车接口
type Cart struct {
	Tips  string `json:"tips"`  // 广告方案
	Alert string `json:"alert"` // 提示文案
	// Cost      int64          `json:"cost"` // 花费多少钱, 前端计算
	GoodsList []Goods `json:"goodsList"`
}

type Goods struct {
	GoodsInfo daoSql.Goods `json:"goodsInfo"`
	GoodsNum  string       `json:"goodsNum"` // 购买的数量
	Selected  string       `json:"selected"` // 是否选中
}
