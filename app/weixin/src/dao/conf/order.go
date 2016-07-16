package conf

import (
	"util"
)

// 订单配置
type Order struct {
	Tips         string
	Alert        string
	StorageAlert string
}

// 读取 cart.yaml的数据
func OrderConf() (Order, error) {
	var confName = "order"

	order := new(Order)
	err := util.AppData(confName, order)
	if nil != err {
		return *order, err
	}

	return *order, err
}
