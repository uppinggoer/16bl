package index

import (
	daoConf "dao/conf"
	daoSql "dao/sql"
)

// 首页接口
type Home struct {
	Banner []daoConf.Button `json:"banner"`
	Nav    []daoConf.Button `json:"nav"`
	Class  []Class          `json:"class"`
}

// 下方列表
type Class struct {
	Name      string         `json:"name"`
	Color     string         `json:"color"`
	Img       string         `json:"img"`
	GoodsList []daoSql.Goods `json:"goodsList"`
}
