package logic

import (
	"fmt"
	"testing"
)

func TestGetMyOrderList(t *testing.T) {
	// 商品列表基础信息
	classList, _ := GetMyOrderList(1, -1, 10)

	fmt.Printf("%+v", classList[0])
	panic(0)
}
