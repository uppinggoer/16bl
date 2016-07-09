package logic

import (
	daoSql "dao/sql"
)

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GenOrderGoods(goodsInfo *daoSql.Goods, goodsNum int64) (*daoSql.OrderGoods, error) {
	orderGoods := daoSql.OrderGoods{
		GoodsId:          goodsInfo.Id,
		GoodsName:        goodsInfo.Name,
		GoodsNorms:       goodsInfo.Norms,
		GoodsImage:       goodsInfo.Image,
		GoodsUnit:        goodsInfo.Unit,
		GoodsNum:         goodsNum,
		GoodsPrice:       int(goodsInfo.Price * goodsNum),
		GoodsMarketprice: int(goodsInfo.Marketprice * goodsNum),
		GoodsCostprice:   0,
	}
	return &orderGoods, nil
}

/**
 * @abstract 根据商品信息+购买数量生成 orderGoods 信息
 * @param goodsInfo *daoSql.Goods
 * @param goodsNum int64
 * @return
 *   addressInfo 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GenOrder(goodsInfo []*daoSql.OrderGoods) (*daoSql.Order, error) {
	var amount = 0
	for _, item := range goodsInfo {
		amount += item.GoodsPrice
	}
	order := daoSql.Order{
		CostAmount:      amount,
		Amount:          amount,
		CostOrderAmount: 0,
		OrderAmount:     amount,
	}
	return &order, nil
}
