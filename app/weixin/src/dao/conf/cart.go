package conf

import (
	"util"
)

// 购物车配置
type Cart struct {
	Tips  string `json:"banner"`
	Alert string `json:"nav"`
}

// 读取 cart.yaml的数据
func CartConf() (Cart, error) {
	var confName = "cart"

	cart := new(Cart)
	err := util.AppData(confName, cart)
	if nil != err {
		return *cart, err
	}

	return *cart, err
}
