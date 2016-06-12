package conf

import (
	"util"
)

// 界面上区块按钮  如 banner/nav/func
type button struct {
	Icon    string `json:"icon"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Trigger string `json:"trigger"`
}

// 下方列表
type class struct {
	Name        string `json:"name"`
	Img         string `json:"img"`
	GoodsIdList []int  `yaml:"goodsIdList"`
}

// 下方列表
type Home struct {
	Banner []button `json:"banner"`
	Nav    []button `json:"nav"`
	Func   []button `json:"func"`
	Class  []class  `json:"class"`
}

var confName = "home"

// 读取 home.yaml的数据
func NewHome() (Home, error) {
	home := new(Home)
	err := util.AppData(confName, home)
	if nil != err {
		return *home, err
	}

	return *home, err
}
