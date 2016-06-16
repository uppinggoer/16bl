// 使用 fake_context 进行测试，可以在不重新编译情况下测试代码功能，类比php测试

package controller

import (
	"testing"

	"util"
)

func TestIndex(t *testing.T) {
	context := util.NewContext("/")
	index := IndexController{}
	err := index.Index(context)
	if err != nil {
		t.Fatal("err:", err)
	}
}

// func TestCart(t *testing.T) {
// 	context := util.NewContext("/cart/list?goods_ids=1,2,3,4&request_id=eet")
// 	index := CartController{}
// 	err := index.CartList(context)
// 	if err != nil {
// 		t.Fatal("err:", err)
// 	}
// }
