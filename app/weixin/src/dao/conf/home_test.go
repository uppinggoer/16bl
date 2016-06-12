package conf_test

import (
	daoConf "dao/conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestNewHome(t *testing.T) {
	file, _ := os.Open("/Users/yanghongzhi/work/16bl/app/weixin/conf/app/home.yaml")
	confData, _ := ioutil.ReadAll(file)

	homeD := &daoConf.Home{}
	yaml.Unmarshal(confData, homeD)
	// fmt.Println(t)
	b, _ := json.Marshal(homeD)
	fmt.Println(string(b))
	t.Fatal("err:", "fail")
}
