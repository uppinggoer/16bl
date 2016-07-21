package logic

import (
	"util"

	daoSql "dao/sql"

	"github.com/labstack/echo"
)

/**
 * @abstract 返回商品列表信息
 * @param ctx
 * @return
 *  []*daoSql.GoodsClass  class_id list 已经排序
 *   map[int64][]daoSql.Goods  map<class_id,goodsList已经排序>
 */
func GetShopData(ctx echo.Context) (classList []*daoSql.GoodsClass, classIdMap map[uint64][]*daoSql.Goods, err error) {
	classList = make([]*daoSql.GoodsClass, 0)
	classIdMap = make(map[uint64][]*daoSql.Goods)

	// 获取 所有商品 信息
	cond := map[string]string{
		"state": "1",
	}
	goodsList, err := daoSql.GetAllGoods(&cond)
	if nil != err {
		return
	}

	// 收集 map(class_id, []goods)
	for _, goodsInfo := range goodsList {
		classIdMap[goodsInfo.ClassId] = append(classIdMap[goodsInfo.ClassId], goodsInfo)
	}

	// 收集分类 id_list
	classIdList := []uint64{}
	for k, _ := range classIdMap {
		classIdList = append(classIdList, k)
		classIdMap[k] = util.SortList(classIdMap[k], "sort", true).([]*daoSql.Goods)
	}

	// 获取所能分类的信息
	classIdInfoMap, err := daoSql.GetClassListById(classIdList)
	if nil != err {
		return
	}
	// 收集分类 id_list
	for _, v := range classIdInfoMap {
		classList = append(classList, v)
	}
	classList = util.SortList(classList, "sort", true).([]*daoSql.GoodsClass)

	return
}
