package conf

// 订单配置
type Env struct {
	ServiceTel string
}

// 读取 env.yaml的数据
func EnvConf() (Env, error) {
	// var confName = "env"

	return Env{ServiceTel: "18211121906"}, nil
}
