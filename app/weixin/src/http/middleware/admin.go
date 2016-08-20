package middleware

import (
	"logic"
	"net/http"
	"time"

	daoSql "dao/sql"

	"github.com/labstack/echo"
)

// UserInfo 用于 echo 框架的获取用户信息
func UserInfo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(ctx echo.Context) error {
			// 执行实体
			if err := userInfo(ctx); nil != err {
				return err
			}
			if err := next(ctx); err != nil {
				return err
			}
			return nil
		})
	}
}

func userInfo(ctx echo.Context) error {
	var curUser = &daoSql.Member{}
	var token string
	var openId string

	// 需要登录的  路由
	var needLogin = map[string]bool{
		"/":               true,
		"/order/prepare":  true,
		"/order/do_order": true,
		"/user":           true,
	}
	// 强制登录
	forceLogin := ctx.QueryParam("force_login") // 参数要求必须登录
	if "1" == forceLogin {
		// 重定向到首页
		return ctx.Redirect(http.StatusFound, logic.GenAuthUrl("/"))
	}
	// 强制刷新
	code := ctx.QueryParam("code")
	if 0 < len(code) {
		// 注册OR刷新 openId,uid
		var err error
		openId, curUser, err = logic.AuthorizeAndUserInfo(code, true)
		if nil != err {
			token = logic.GenToken(curUser.MemberId, openId, time.Now().Unix())
		}
		goto END
	}

	// 是否有 openid cookie
	if openIdCookie, err := ctx.Cookie("openid"); nil == err {
		openId = openIdCookie.Value()
	}

	// openId 及 token
	if "" != openId {
		tokenCookie, err := ctx.Cookie("token")
		if nil == err {
			token = tokenCookie.Value()
		}

		// 存在 token cookie 计算出 uid
		if "" != token {
			curUser = logic.CheckToken(token, openId)
		} else {
			// 没有 token cookie 生成 token   使用openId登录
			curUser = logic.GetUserByOpenId(openId)
			// 根据 openId 计算新的 uid
			token = logic.GenToken(curUser.MemberId, openId, time.Now().Unix())
			goto END
		}
	}

	// 重定向到 微信授权页面
	if needUserInfo, ok := needLogin[ctx.Path()]; ok && needUserInfo {
		// 强制获取信息
		if "" == openId {
			// 重定向到首页
			return ctx.Redirect(http.StatusFound, logic.GenAuthUrl("/"))
		}
	}
END:

	// 种植 token cookie
	cookie := new(echo.Cookie)
	cookie.SetName("token")
	cookie.SetValue(token)
	cookie.SetExpires(time.Now().Add(720 * time.Hour))
	ctx.SetCookie(cookie)

	// 种植 token openid
	cookie = new(echo.Cookie)
	cookie.SetName("openid")
	cookie.SetValue(openId)
	cookie.SetExpires(time.Now().Add(720 * time.Hour))
	ctx.SetCookie(cookie)

	ctx.Set("uid", curUser.MemberId)
	ctx.Set("curUser", curUser)

	return nil
}
