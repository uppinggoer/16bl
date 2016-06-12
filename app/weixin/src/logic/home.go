package logic

import (
	_ "fmt"

	daoConf "dao/conf"
	daoSql "dao/sql"

	"github.com/labstack/echo"
)

/**
 * @abstract 返回首页主框架信息，与goodsList 信息
 * @param ctx
 * @return
 *   confInfo  目前直接返回 dao/conf/homeConf中字段
 *   goodsList 目前直接返回 dao/sql/goods中字段
 */
func GetHomeData(ctx echo.Context) (daoConf.Home, map[int64]*daoSql.Goods, error) {
	homeConf, err := daoConf.NewHome()
	if nil != err {
		// log
		return daoConf.Home{}, map[int64]*daoSql.Goods{}, err
	}

	// all the goods
	goodsIdList := []int64{}
	for _, classInfo := range homeConf.Class {
		goodsIdList = append(goodsIdList, classInfo.GoodsIdList...)
	}

	goodsIdMap, err := daoSql.GetGoodsListById(goodsIdList)
	if nil != err {
		return daoConf.Home{}, map[int64]*daoSql.Goods{}, err
	}

	return homeConf, goodsIdMap, nil
}
