package logic

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"

	daoConf "dao/conf"
	daoRedis "dao/redis"
	daoSql "dao/sql"
	. "global"
	"util"
)

/**
 * @abstract 获得JS授权信息
 * @param goodsList
 * @return
 */
func AuthJsInfo(authUrl string) (timestamp int64, nonceStr, signature string, err error) {
	timestamp = time.Now().Unix()
	nonceStr = getNonceStr()
	signature = ""
	err = nil

	var jsTicket = ""
	accessToken := getAccessToken()
	if 0 >= len(accessToken) {
		accessToken = touchAccessToken(true)
	} else {
		jsTicket = getJsTicket()
	}

	if 0 >= len(jsTicket) {
		jsTicket = touchJsTicket(accessToken, true)
	}
	if 0 >= len(jsTicket) {
		// log
		err = JsAuthError
		return
	}
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, nonceStr, timestamp, authUrl)
	bs := sha1.Sum([]byte(str))
	signature = fmt.Sprintf("%x", bs)
	return
}

func GetUserByOpenId(openId string) *daoSql.Member {
	uid, _ := daoSql.GetUidByOpenId(openId)
	if 0 < uid {
		userInfo, err := daoSql.GetInfoByUid(uid)
		if nil != err {
			return userInfo
		}
	}

	return &daoSql.Member{}
}

// 使用用户授权 code 获取用户 openid
// 注册or刷新用户信息
// 返回 openId, member, err
func AuthorizeAndUserInfo(code string, repeat bool) (openId string, userInfo *daoSql.Member, err error) {
	openId = ""
	userInfo = &daoSql.Member{}
	err = nil

	if 0 < len(code) {
		// log
		return
	}

	// get access_token
	userInfoUrl := "https://api.weixin.qq.com/sns/oauth2/access_token"
	params := url.Values{
		"appid":      {daoConf.WeChat_AppId},
		"secret":     {daoConf.WeChat_Secret},
		"code":       {code},
		"grant_type": {"authorization_code"},
	}
	for intTry := 2; intTry > 0 && repeat; intTry-- {
		var resBody []byte
		resBody, err = util.CallHttp(util.HTTP_GET, userInfoUrl, params)
		if nil != err {
			// log error
			continue
		}

		userInfoRes := &struct {
			ErrCode     int64  `json:"errcode"`
			ErrMsg      string `json:"errmsg"`
			AccessToken string `json:"access_token"`
			OpenId      string `json:"openid"`
		}{}
		err = json.Unmarshal(resBody, userInfoRes)
		if nil != err || 0 != userInfoRes.ErrCode {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		// openId
		openId = userInfoRes.OpenId
		// 获取 weChatInfo
		var weChatInfo *daoSql.WeChatInfo
		weChatInfo, err = getUserFromWechat(userInfoRes.AccessToken, userInfoRes.OpenId, false)
		if err != nil {
			var uid uint64
			uid, err = daoSql.GetUidByOpenId(weChatInfo.OpenId)
			if 0 >= uid {
				// 注册 生成用户主体信息
				userInfo.Name = weChatInfo.Nickname
				userInfo.Avatar = weChatInfo.Headimgurl
				userInfo.Sex = weChatInfo.Sex
				daoSql.DB.Create(userInfo)

				uid = userInfo.MemberId
				// 注册 生成 微信用户 绑定信息
				weChatBind := &daoSql.WechatBind{
					MemberId:   uid,
					Openid:     weChatInfo.OpenId,
					Nickname:   weChatInfo.Nickname,
					Sex:        weChatInfo.Sex,
					City:       weChatInfo.City,
					Headimgurl: weChatInfo.Headimgurl,
					Province:   weChatInfo.Province,
				}
				daoSql.DB.Create(weChatBind)
				return
			} else {
				userInfo, err = daoSql.GetInfoByUid(uid)
				return
			}
		}
	}

	return
}

// 生成用户授权信息的 url
func GenAuthUrl(redirectUrl string) string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo#wechat_redirect",
		daoConf.WeChat_AppId, redirectUrl)
}

// 返回 uid
func CheckToken(token, openId string) *daoSql.Member {
	// token md5(openId+uid+time+"xZCFwetwq4")
	curTime := time.Now().Unix()
	arrStr := strings.Split(token, "$")
	if 3 > len(arrStr) {
		return &daoSql.Member{}
	}
	lastTime := util.Atoi(arrStr[1], 64, true).(int64)
	// token 有效为15天(也没什么蛋用！！)  如此一个微信号只能登录一个用户
	if 86400*15 < (curTime - lastTime) {
		return &daoSql.Member{}
	}

	uid := util.Atoi(arrStr[2], 64, false).(uint64)
	if token == GenToken(uid, openId, lastTime) && 0 < uid {
		userInfo, err := daoSql.GetInfoByUid(uid)
		if nil == err {
			return userInfo
		}
	}
	return &daoSql.Member{}
}

func GenToken(uid uint64, openId string, curTime int64) string {
	// token md5(openId+uid+time+"xZCFwetwq4")&time&uid
	str := fmt.Sprintf("%s%d%dxZCFwetwq4", openId, uid, curTime)
	str = fmt.Sprintf("%x$%d$%d", md5.Sum([]byte(str)), curTime, uid)
	return str
}

// 返回 open_id,nick,imgUrl
// accessToken 微信用户授权 token
func getUserFromWechat(accessToken, openid string, repeat bool) (*daoSql.WeChatInfo, error) {
	// get userinfo
	userInfoUrl := "https://api.weixin.qq.com/sns/userinfo"
	params := url.Values{
		"lang":         {"zh_CN"},
		"access_token": {accessToken},
		"openid":       {openid},
	}

	for intTry := 2; intTry > 0 && repeat; intTry-- {
		resBody, err := util.CallHttp(util.HTTP_GET, userInfoUrl, params)
		if nil != err {
			// log error
			continue
		}

		userInfoRes := &struct {
			ErrCode int64  `json:"errcode"`
			ErrMsg  string `json:"errmsg"`
			daoSql.WeChatInfo
		}{}
		err = json.Unmarshal(resBody, userInfoRes)
		if nil != err || 0 != userInfoRes.ErrCode {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		return &userInfoRes.WeChatInfo, err
	}

	// return userInfoRes.WeChatInfo
	return &daoSql.WeChatInfo{}, JsAuthError
}

func getAccessToken() string {
	return daoRedis.NewRedisClient().Key(daoRedis.KeyWeiXin, daoRedis.KeyWeiXinAccessToken).GET("")
}

func touchAccessToken(repeat bool) string {
	// 设置 redis key
	redis := daoRedis.NewRedisClient().Key(daoRedis.KeyWeiXin, daoRedis.KeyWeiXinAccessToken)

	// 重新取 accessToken
	accessUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token")

	params := url.Values{
		"grant_type": {"client_credential"},
		"appid":      {daoConf.WeChat_AppId},
		"secret":     {daoConf.WeChat_Secret},
	}

	var err error
	var accessToken string
	accessFail := true
	for intTry := 2; intTry > 0 && accessFail && repeat; intTry-- {
		resBody, err := util.CallHttp(util.HTTP_GET, accessUrl, params)
		if nil != err {
			// log error
			continue
		}

		accessTokenRes := &struct {
			AccessToken string `json:"access_token"`
			ErrCode     int64  `json:"errcode"`
			ErrMsg      string `json:"errmsg"`
		}{}
		err = json.Unmarshal(resBody, accessTokenRes)
		if nil != err || 0 != accessTokenRes.ErrCode {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		accessToken = accessTokenRes.AccessToken
		err = redis.SET("", accessToken, int64(1.8*3600))
		if nil != err {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		accessFail = false
	}

	if accessFail || nil != err {
		return ""
	}
	return accessToken
}

func getJsTicket() string {
	return daoRedis.NewRedisClient().Key(daoRedis.KeyWeiXin, daoRedis.KeyWeiXinJsToken).GET("")
}

func touchJsTicket(accessToken string, repeat bool) string {
	// 设置 redis key
	redis := daoRedis.NewRedisClient().Key(daoRedis.KeyWeiXin, daoRedis.KeyWeiXinJsToken)

	// 重新取 jsTicket
	accessUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket")

	params := url.Values{
		"type":         {"jsapi"},
		"access_token": {accessToken},
	}

	var err error
	var jsTicket string
	accessFail := true
	for intTry := 2; intTry > 0 && accessFail && repeat; intTry-- {
		resBody, err := util.CallHttp(util.HTTP_GET, accessUrl, params)
		if nil != err {
			// log error
			continue
		}

		accessTokenRes := &struct {
			Ticket  string `json:"ticket"`
			ErrCode int64  `json:"errcode"`
			ErrMsg  string `json:"errmsg"`
		}{}
		err = json.Unmarshal(resBody, accessTokenRes)
		if nil != err || 0 != accessTokenRes.ErrCode {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		jsTicket = accessTokenRes.Ticket
		err = redis.SET("", jsTicket, int64(1.8*3600))

		if nil != err {
			// objLog.Errorln("add elemeid,barcode map error:", newFood.Data.FoodId, goodInfo.AppFoodCode)
			continue
		}

		accessFail = false
	}

	if accessFail || nil != err {
		return ""
	}
	return jsTicket
}

func getNonceStr() string {
	var arrStr = strings.Split("abcdefghijklmnopqrstuvwsyzABCEFGHIJKLMNOPQRSTUVWSYXYZ1234567890", "")
	var strList = make([]string, 8)
	for i := 7; i >= 0; i-- {
		strList[i] = arrStr[rand.Intn(62)]
	}

	return strings.Join(strList, "")
}

// outputXml
func outputXml(outputMap map[string]interface{}) string {
	var buf bytes.Buffer

	buf.WriteString("<xml>")
	for key, value := range outputMap {
		switch value.(type) {
		case int8, uint8, int16, uint16, int, int32, uint32, int64, uint64:
			buf.WriteString(fmt.Sprint("<%s>%d</%s>", key, value, key))
		case float32, float64:
			buf.WriteString(fmt.Sprint("<%s>%f</%s>", key, value, key))
		case string:
			buf.WriteString(fmt.Sprint("<%s><![CDATA[%s]]</%s>", key, value, key))
		default:
			valueBuf, err := json.Marshal(value)
			if nil != err {
				// log
			} else {
				buf.WriteString(fmt.Sprint("<%s><![CDATA[%s]]</%s>", key, string(valueBuf), key))
			}
		}
	}
	buf.WriteString("</xml>")

	return string(buf.Bytes())
}

// parseXml
func parseXml(data []byte, xmlStruct interface{}) error {
	if 0 >= len(data) {
		return XmlParseEmpty
	}

	err := xml.Unmarshal(data, xmlStruct)
	if err != nil {
		return err
	}
	return nil
}
