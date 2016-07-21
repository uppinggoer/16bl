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
func GetCartInfo(goodsIdList []uint64) (string, map[uint64]*daoSql.Goods, error) {
	if 0 >= len(goodsIdList) {
		return "", map[uint64]*daoSql.Goods{}, CartEmpty
	}

	goodsIdMap, err := daoSql.GetGoodsListById(goodsIdList)
	if nil != err {
		return "", map[uint64]*daoSql.Goods{}, err
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

// 购物车详情
type CartInfo struct {
	GoodsId  uint64 `json:"goods_id"`
	Selected uint8  `json:"selected"`
	GoodsNum uint16 `json:"goods_num"`
}

/**
 * @abstract 验证并修正库存信息  如果只有3个，购买5个，会强制改为3个
 * @param goodsList
 * @return
 *   exception 购物车提示信息
 *   goodsList 目前直接返回 dao/sql/goods中字段
 *   err
 */
func VerifyGoodsNum(cartGoodsIdMap map[uint64]*daoSql.Goods, goodsList []*CartInfo) (string, error) {
	// 遍历商品 验证库存不足信息
	goodsNoStorage := []string{}
	for _, goodsInfo := range goodsList {
		if v, ok := cartGoodsIdMap[goodsInfo.GoodsId]; ok {
			if goodsInfo.GoodsNum > v.Storage {
				goodsInfo.GoodsNum = v.Storage
				// 校正购买数量
				goodsNoStorage = append(goodsNoStorage, cartGoodsIdMap[goodsInfo.GoodsId].Name+"-"+cartGoodsIdMap[goodsInfo.GoodsId].Norms)
			}
		} else {
			// log
			continue
		}
	}

	goodsException := ""
	if 0 < len(goodsNoStorage) {
		goodsException += strings.Join(goodsNoStorage, ", ")
	}
	return goodsException, nil
}
