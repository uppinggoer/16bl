// 测试框架配合 util.fake_context 使用  正式环境中不可过滤掉，目前未做处理

package controller

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type TestController struct{}

// 注册路由
func (self TestController) RegisterRoute(e *echo.Group) {
	e.Get("/test", echo.HandlerFunc(self.Test))
}

// 测试框架入口
func (TestController) Test(ctx echo.Context) error {
	fileName := "/tmp/" + ctx.QueryParam("file")
	file, _ := os.Open(fileName)
	data, _ := ioutil.ReadAll(file)

	return ctx.JSONBlob(http.StatusOK, data)
}
