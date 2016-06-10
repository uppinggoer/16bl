package util

import (
	"encoding/json"
	"net/http"

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

// render html 输出
func Render(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	contentTpl = TPL_PATH + contentTpl + ".tpl"

	// // 如果没有定义css和js模板，则定义之
	// if jsTpl := tpl.Lookup("js"); jsTpl == nil {
	// 	tpl.Parse(`{{define "js"}}{{end}}`)
	// }
	// if jsTpl := tpl.Lookup("css"); jsTpl == nil {
	// 	tpl.Parse(`{{define "css"}}{{end}}`)
	// }

	// buf := new(bytes.Buffer)
	// err := tpl.Execute(buf, data)
	// if err != nil {
	// 	// objLog.Errorln("excute template error:", err)
	// 	return err
	// }

	// return ctx.HTML(http.StatusOK, buf.String())
	return ctx.HTML(http.StatusOK, contentTpl)
	// ctx.HTML(http.StatusOK, "<html><body>XXX</body></html>")
}
