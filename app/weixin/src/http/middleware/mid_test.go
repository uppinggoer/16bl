package middleware

import (
	"testing"

	"util"
)

func TestUserInfo(t *testing.T) {
	// context := util.NewContext("", "/a?code=XXX", "force_login=1")
	context := util.NewContext("", "/a", "force_login=1")
	// context.Request().Header().Set("cookie", "openid=1")
	context.Request().Header().Set("cookie", "token=412b56b6b0085f40dd5e0b1de3bfd4d2$1471362113$1;openid=XXXX;")
	err := userInfo(context)
	if err != nil {
		t.Fatal("err:", err)
	}
	panic(0)
}
