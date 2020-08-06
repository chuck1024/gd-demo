/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/app/model"
	"github.com/chuck1024/gd-demo/app/service/sp"
	"github.com/chuck1024/gd-demo/route"
	"github.com/chuck1024/inject"
	"github.com/chuck1024/mysqldb"
)

func Run() {
	// init gd
	d := gd.Default()

	// init inject
	inject.InitDefault()
	inject.SetLogger(dlog.Global)
	defer inject.Close()

	// inject UserDao
	inject.Reg("UserDao", (*model.UserDao)(&model.UserDao{MysqlClient: &mysqldb.MysqlClient{
		DbConfPath: "conf/db.ini",
	}}))

	// inject dependency
	inject.RegisterOrFail("serviceProvider", (*sp.ServiceProvider)(nil))
	err := sp.Init()
	if err != nil {
		dlog.Crash("init package sp fail,err=%v", err)
	}

	// route register
	route.Register(d)

	// gd run
	if err := d.Run(); err != nil {
		dlog.Crashf("gd-demo run occur err:%v", err)
		return
	}
}
