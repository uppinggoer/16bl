package conf

import (
	_ "encoding/json"
	_ "fmt"
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
	Name  string   `json:"name"`
	Img   string   `json:"img"`
	Goods []Button `json:"goods"`
}

// 下方列表
type Home struct {
	Banner []Button `json:"banner"`
	Nav    []Button `json:"nav"`
	Func   []Button `json:"func"`
	Class  []Class  `json:"class"`
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
