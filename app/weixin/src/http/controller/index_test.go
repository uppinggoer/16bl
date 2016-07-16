// 使用 fake_context 进行测试，可以在不重新编译情况下测试代码功能，类比php测试

package controller

import (
	"testing"

	. "global"
	"util"
)

func TestIndex(t *testing.T) {
	context := util.NewContext("", "/", "")
	index := IndexController{}
	err := index.Index(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestGenCartHtml(t *testing.T) {
	context := util.NewContext(STATIC_PATH+"cart.html", "/cart/list", "")
	index := CartController{}
	err := index.GenCartIndexHtml(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestCart(t *testing.T) {
	index := CartController{}
	context := util.NewContext("", "/cart/list", `goods_list=%5B%7B%22goods_id%22%3A%221%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%229%22%7D%2C%7B%22goods_id%22%3A%222%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%226%22%7D%2C%7B%22goods_id%22%3A%223%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%2210%22%7D%2C%7B%22goods_id%22%3A%224%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%221%22%7D%5D`)
	err := index.CartList(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestPrepareOrder(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/prepare", `goods_list=%5B%7B%22goods_id%22%3A%221%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%222%22%7D%2C%7B%22goods_id%22%3A%222%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%224%22%7D%2C%7B%22goods_id%22%3A%223%22%2C%22selected%22%3A%221%22%2C%22goods_num%22%3A%222%22%7D%5D`)
	err := index.PrepareOrder(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestShopList(t *testing.T) {
	index := ShopController{}
	context := util.NewContext("", "/shop/list", ``)
	err := index.ShopList(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestMyOrderList(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/list", ``)
	err := index.MyOrderList(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestTest(t *testing.T) {
	index := TestController{}
	context := util.NewContext("", "/test?file=text", ``)
	err := index.Test(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestGenOrderListHtml(t *testing.T) {
	context := util.NewContext(STATIC_PATH+"orderList.html", "/order/list", "")
	index := OrderController{}
	err := index.GenOrderListHtml(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}
