// 使用 fake_context 进行测试，可以在不重新编译情况下测试代码功能，类比php测试

package controller

import (
	"fmt"
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
	context := util.NewContext("", "/cart/list", `goods_list=[{"goods_id":"102","selected":"1","goods_num":"2"},{"goods_id":"104","selected":"1","goods_num":"2"}]`)
	err := index.CartList(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestPrepareOrder(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/prepare?uid=10", `goods_list=[{"goods_id":"102","selected":"1","goods_num":"20000"},{"goods_id":"103","selected":"1","goods_num":"4"},{"goods_id":"104","selected":"1","goods_num":"2"}]`)
	fmt.Println(context.Path())
	panic(0)
	err := index.PrepareOrder(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}
func TestDoOrder(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/do_order", `goods_list=[{"goods_id":"101","selected":"1","goods_num":"1"},{"goods_id":"102","selected":"1","goods_num":"2"}]`)
	context.Request().Header().Set("cookie", "zhima_debug=1")
	err := index.DoOrder(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}
func TestDetail(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/do_order?order_sn=17288449600", ``)
	// context.Request().Header().Set("cookie", "zhima_debug=1")

	err := index.Detail(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestCancelOrder(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/eval_order", `order_sn=17288079600&stars=5&feedback=XDEFGF`)
	// context.Request().Header().Set("cookie", "zhima_debug=1")

	err := index.EvalOrder(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}
func TestEvalOrder(t *testing.T) {
	index := OrderController{}
	context := util.NewContext("", "/order/do_order", `order_sn=17288079600&cancel_flag=11`)
	// context.Request().Header().Set("cookie", "zhima_debug=1")

	err := index.CancelOrder(context)
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

func TestAddressList(t *testing.T) {
	index := AddressController{}
	context := util.NewContext("", "/address/list?from=order", ``)
	// context.Request().Header().Set("cookie", "zhima_debug=1")
	err := index.AddressList(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestAddressModify(t *testing.T) {
	index := AddressController{}
	context := util.NewContext("", "/address/modify", `true_name=%E6%9D%A8%E5%AE%8F%E5%BF%97&mobile=18511280986&live_area=%E8%A5%BF%E5%B7%A5%E5%A4%A7&address=%E6%98%9FF+632&addressId=2&uid=10`)
	// context.Request().Header().Set("cookie", "zhima_debug=1")

	err := index.AddressModify(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestAddressDel(t *testing.T) {
	index := AddressController{}
	context := util.NewContext("", "/address/modify", `addressId=2&uid=10`)
	// context.Request().Header().Set("cookie", "zhima_debug=1")

	err := index.AddressDel(context)
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
