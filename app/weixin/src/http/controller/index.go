// Copyright 2016 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	polaris@studygolang.com

package controller

import (
	daoConf "dao/conf"
	_ "fmt"
	"util"

	"github.com/labstack/echo"
)

type IndexController struct{}

// 注册路由
func (self IndexController) RegisterRoute(e *echo.Group) {
	e.Get("/", echo.HandlerFunc(self.Index))
}

// Index 首页
func (IndexController) Index(ctx echo.Context) error {
	homeConf, err := daoConf.NewHome()

	if nil != err {
		// log
		return util.Fail(ctx, 10, "XXX")
	}
	// fmt.Println("/n/n%v/n", m)

	// time.Sleep(3 * time.Second)

	return util.Success(ctx, homeConf)
	// return util.Render(ctx, "", homeConf)
}
