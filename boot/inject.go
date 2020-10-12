/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/app/model"
	"github.com/chuck1024/gd-demo/app/service/sp"
	"github.com/chuck1024/gd/databases/mysqldb"
	"github.com/chuck1024/gd/databases/redisdb"
	"github.com/chuck1024/gd/dlog"
	"github.com/chuck1024/gd/runtime/inject"
)

func Inject() {
	// inject UserDao
	inject.Reg("UserDao", (*model.UserDao)(&model.UserDao{MysqlClient: &mysqldb.MysqlClient{
		DbConfPath: "conf/db.ini",
	}}))

	// inject SessionCache
	inject.Reg("SessionCache", (*model.SessionCache)(&model.SessionCache{RedisConfig: &redisdb.RedisConfig{
		Addrs: gd.Config("Redis", "addr").Strings(","),
	}}))

	// inject dependency
	inject.RegisterOrFail("serviceProvider", (*sp.ServiceProvider)(nil))
	err := sp.Init()
	if err != nil {
		dlog.Crashf("init package sp fail,err=%v", err)
	}
}
