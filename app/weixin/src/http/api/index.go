package api

import (
	daoSql "dao/sql"
)

// 首页接口
type Home struct {
	Banner []button `json:"banner"`
	Nav    []button `json:"nav"`
	Func   []button `json:"func"`
	Class  []class  `json:"class"`
}

// 界面上区块按钮  如 banner/nav/func
type button struct {
	Icon    string `json:"icon"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Trigger string `json:"trigger"`
}

// 下方列表
type class struct {
	Name      string      `json:"name"`
	Img       string      `json:"img"`
	GoodsList []sql.Goods `json:"goodsList"`
}
