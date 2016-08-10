package middleware

import (
	"logic"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// UserInfo 用于 echo 框架的获取用户信息
func UserInfo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(ctx echo.Context) error {
			var uid uint64
			var token string
			var openId string

			// 需要登录的  路由
			var needLogin = map[string]bool{
				"/":               true,
				"/order/do_order": true,
				"/user":           true,
			}
			for {
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
					openId, uid, err = logic.AuthorizeAndUserInfo(code)
					if nil != err {
						token = logic.GenToken(uid, openId, time.Now().Unix())
					}
					break
				}

				// 是否有 openid cookie
				openIdCookie, err := ctx.Cookie("openid")
				openId := ""
				if nil == err {
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
						uid = logic.CheckToken(token, openId)
					}
					// 没有 token cookie 生成 token
					if 0 >= uid {
						// 根据 openId 计算新的 uid
						// uid = daoSql.GetUidByOpenId(openId)
						token = logic.GenToken(uid, openId, time.Now().Unix())
						break
					}
				}

				// 重定向到 微信授权页面
				needUserInfo, ok := needLogin[ctx.Path()]
				if ok && needUserInfo {
					// 强制获取信息
					if "" != openId {
						// 重定向到首页
						return ctx.Redirect(http.StatusFound, logic.GenAuthUrl("/"))
					}
				}
			}

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

			ctx.Set("uid", uid)
			if err := next(ctx); err != nil {
				return err
			}
			return nil
		})
	}
}
