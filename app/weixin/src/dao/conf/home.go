package conf

import (
	"util"
)

// 界面上区块按钮  如 banner/nav/func
type Button struct {
	Icon    string `json:"icon"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Trigger string `json:"trigger"`
}

// 下方列表
type Class struct {
	Name        string   `json:"name"`
	Img         string   `json:"img"`
	Color       string   `json:"color"`
	GoodsIdList []uint64 `yaml:"goodsIdList" json:"goodsIdList"`
}

// 下方列表
type Home struct {
	Banner []Button `json:"banner"`
	Nav    []Button `json:"nav"`
	Func   []Button `json:"func"`
	Class  []Class  `json:"class"`
}

// 读取 home.yaml的数据
func HomeConf() (Home, error) {
	var confName = "home"

	home := new(Home)
	err := util.AppData(confName, home)
	if nil != err {
		return *home, err
	}

	return *home, err
}
