package logic

import (
	"fmt"
	"testing"
)

func TestGetJsTicket(t *testing.T) {
	// 商品列表基础信息
	fmt.Println(getJsTicket())
	panic(0)
}

func TestGenAuthUrl(t *testing.T) {
	t.Fatal(GenAuthUrl("www.baidu.com"))
}

func TestAuthJsInfo(t *testing.T) {
	// 商品列表基础信息
	timestamp, nonceStr, signature, err := AuthJsInfo("www.baidu.com")
	t.Fatal(timestamp, nonceStr, signature, err)
}
