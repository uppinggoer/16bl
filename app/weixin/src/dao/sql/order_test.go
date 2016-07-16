package sql

import (
	"fmt"
	"testing"
)

func TestGetListById(t *testing.T) {
	// 商品列表基础信息
	orderList, _ := GetListById(1, -1, 20)

	fmt.Printf("%#v", orderList[0])
	panic(0)
}
