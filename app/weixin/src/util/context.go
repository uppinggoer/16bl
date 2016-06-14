package util

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"html/template"
	"net/http"
	"strings"

	. "global"

	"github.com/labstack/echo"
)

func Success(ctx echo.Context, data interface{}) error {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "操作成功",
		"data": data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		// 可以提前捕获一步
		return err
	}

	if ctx.Response().Committed() {
		// getLogger(ctx).Flush()
		return nil
	}

	return ctx.JSONBlob(http.StatusOK, b)
}

func Fail(ctx echo.Context, code int, msg string) error {
	if ctx.Response().Committed() {
		// getLogger(ctx).Flush()
		return nil
	}

	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	// getLogger(ctx).Errorln("operate fail:", result)
	return ctx.JSON(http.StatusOK, result) // 最终调用 JSONBlob
}

type footMentEntiy struct {
	Activite bool
	Url      string
	Name     string
	Icon     string
}
type renderEntity struct {
	Content  interface{}
	FootMenu []footMentEntiy
}

// render html 输出
// contentTpl 模板名  多个模板用,分割
func Render(ctx echo.Context, contentTpl string, data interface{}) error {
	tplInfo := renderEntity{}

	// Content 元素
	tplInfo.Content = data

	// 填写下方导航
	foot := []footMentEntiy{}
	foot = append(foot, footMentEntiy{
		Activite: true,
		Icon:     "icon-home",
		Url:      "www.baidu.com",
		Name:     "首页",
	})
	foot = append(foot, footMentEntiy{
		Activite: false,
		Icon:     "icon-list",
		Url:      "www.baidu.com",
		Name:     "超市",
	})
	foot = append(foot, footMentEntiy{
		Activite: false,
		Icon:     "icon-shopping-cart",
		Url:      "www.baidu.com",
		Name:     "购物车",
	})
	foot = append(foot, footMentEntiy{
		Activite: false,
		Icon:     "icon-user",
		Url:      "www.baidu.com",
		Name:     "我的",
	})
	tplInfo.FootMenu = foot

	// 所有模板
	contentTpl = "layout" + "," + contentTpl + "," + "home/bottom-nav"
	// 为了使用自定义的模板函数，首先New一个以第一个模板文件名为模板名。
	// 这样，在ParseFiles时，新返回的*Template便还是原来的模板实例
	htmlFiles := strings.Split(contentTpl, ",")
	for i, contentTpl := range htmlFiles {
		htmlFiles[i] = TPL_PATH + contentTpl + ".tpl"
	}
	tpl, err := template.New("layout.tpl").ParseFiles(htmlFiles...)
	if err != nil {
		// objLog.Errorf("解析模板出错（ParseFiles）：[%q] %s\n", Request(ctx).RequestURI, err)
		return err
	}
	return executeTpl(ctx, tpl, tplInfo)
}

// 真正渲染模板
func executeTpl(ctx echo.Context, tpl *template.Template, data interface{}) error {
	// objLog := logic.GetLogger(ctx)

	// 如果没有定义css和js模板，则定义之
	if jsTpl := tpl.Lookup("js"); jsTpl == nil {
		tpl.Parse(`{{define "js"}}{{end}}`)
	}
	if jsTpl := tpl.Lookup("css"); jsTpl == nil {
		tpl.Parse(`{{define "css"}}{{end}}`)
	}

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, data)
	if err != nil {
		// objLog.Errorln("excute template error:", err)
		return err
	}

	return ctx.HTML(http.StatusOK, buf.String())
}
