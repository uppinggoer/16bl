package logic

import (
	"strings"

	daoSql "dao/sql"
	. "global"
)

/**
 * @abstract 返回购物车信息
 * @param goodsList
 * @return
 *   exception 购物车提示信息
 *   goodsList 目前直接返回 dao/sql/goods中字段
 *   err
 */
func GetCartInfo(goodsIdList []int64) (string, map[int64]*daoSql.Goods, error) {
	if 0 >= len(goodsIdList) {
		return "", map[int64]*daoSql.Goods{}, CartEmpty
	}

	goodsIdMap, err := daoSql.GetGoodsListById(goodsIdList)
	if nil != err {
		return "", map[int64]*daoSql.Goods{}, err
	}

	goodsNotValid := []string{}
	for goodsId, goodsInfo := range goodsIdMap {
		if goodsInfo.State != daoSql.GOODS_ON {
			goodsNotValid = append(goodsNotValid, goodsInfo.Name+"-"+goodsInfo.Norms) // 商品名-商品规格
			delete(goodsIdMap, goodsId)                                               //不会影响遍历 因为range用的是复本
		}
	}

	goodsException := ""
	if 0 < len(goodsNotValid) {
		goodsException += strings.Join(goodsNotValid, ", ")
	}
	return goodsException, goodsIdMap, nil
}
