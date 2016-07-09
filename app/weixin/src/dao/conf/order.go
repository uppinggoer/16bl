package conf

import (
	"util"
)

// 读取 cart.yaml的数据
func OrderConf() (Cart, error) {
	var confName = "cart"

	cart := new(Cart)
	err := util.AppData(confName, cart)
	if nil != err {
		return *cart, err
	}

	return *cart, err
}
