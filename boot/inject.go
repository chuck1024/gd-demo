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
	"github.com/chuck1024/inject"
	"github.com/chuck1024/mysqldb"
	"github.com/chuck1024/redisdb"
)

func Inject(d *gd.Engine) {
	// inject UserDao
	inject.Reg("UserDao", (*model.UserDao)(&model.UserDao{MysqlClient: &mysqldb.MysqlClient{
		DbConfPath: "conf/db.ini",
	}}))

	// inject SessionCache
	inject.Reg("SessionCache", (*model.SessionCache)(&model.SessionCache{RedisConfig: &redisdb.RedisConfig{
		Addrs:   d.Config("Redis","addr").Strings(","),
	}}))

	// inject dependency
	inject.RegisterOrFail("serviceProvider", (*sp.ServiceProvider)(nil))
	err := sp.Init()
	if err != nil {
		dlog.Crashf("init package sp fail,err=%v", err)
	}
}