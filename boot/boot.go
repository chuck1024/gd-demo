/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/app/model"
	"github.com/chuck1024/gd-demo/route"
	"github.com/chuck1024/inject"
)

func Run() {
	// init gd
	d := gd.Default()

	// init inject
	inject.InitDefault()
	defer inject.Close()

	// inject userDao
	inject.Reg("UserDao", (*model.UserDao)(&model.UserDao{MysqlClient: }))

	// route register
	route.Register(d)

	// gd run
	if err := d.Run(); err != nil {
		dlog.Crashf("gd-demo run occur err:%v", err)
		return
	}
}
