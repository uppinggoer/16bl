// 伪造一个 context 用于go test时测试函数
// 将 response 输出到文件 /tmp/test
// 注册路由  ip:port/test?file=test 即可以在 chrome 等查阅此文件

package util

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
)

func NewContext(fileName, url, queryString string) echo.Context {
	e := echo.New()
	req := test.NewRequest(echo.POST, url, strings.NewReader(queryString))
	req.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec).(echo.Context)

	if 0 >= len(fileName) {
		fileName = "/tmp/text"
	}
	if Exist(fileName) {
		os.Truncate(fileName, 0)
	}
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// 将 http.response 输出到文件
	c.Response().SetWriter(file)
	return c
}
