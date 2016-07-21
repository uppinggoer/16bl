package api

import daoConf "dao/conf"

// 首页接口
type Home struct {
	Banner []daoConf.Button `json:"banner"`
	Nav    []daoConf.Button `json:"nav"`
	Class  []*HomeClass     `json:"class"`
}

func (self *Home) Format() {
	for _, v := range self.Class {
		v.Format()
	}
}

// 下方列表
type HomeClass struct {
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	Img       string   `json:"img"`
	GoodsList []*Goods `json:"goodsList"`
}

func (self *HomeClass) Format() {
	for _, goodsItem := range self.GoodsList {
		goodsItem.Format()
	}
}
