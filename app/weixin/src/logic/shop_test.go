package logic

import (
	"fmt"
	"testing"
	"time"
)

func TestGetShopData(t *testing.T) {
	// 商品列表基础信息
	classList, goodsIdMap, _ := GetShopData(nil)

	fmt.Printf("%+v", classList)
	fmt.Println(goodsIdMap)
	panic(0)
	t.Fatal("%#v,%#v", classList, goodsIdMap)
}

func TestGenToken(t *testing.T) {
	// 商品列表基础信息
	token := GenToken(2, "XXXXX", time.Now().Unix())
	fmt.Println(token)
	panic(0)
}
