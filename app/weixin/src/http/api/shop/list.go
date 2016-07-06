package index

import (
	daoSql "dao/sql"
)

// 购物车接口
type Shop struct {
	ClassList []Class `json:"classList"`
}

type Class struct {
	ClassName string          `json:"className"`
	ClassId   int64           `json:"-"`
	GoodsList []*daoSql.Goods `json:"goodsList"`
}
