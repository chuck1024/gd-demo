/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"gd-demo/app/model"
	"gd-demo/app/service/sp"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/databases/redisdb"
	"github.com/chuck1024/gd/runtime/inject"
)

func Inject() {
	// inject UserDao
	inject.Reg("UserDao", (*model.UserDao)(nil))

	// inject SessionCache
	inject.Reg("SessionCache", (*model.SessionCache)(&model.SessionCache{RedisConfig: &redisdb.RedisConfig{
		Addrs: gd.Config("Redis", "addr").Strings(","),
	}}))

	// inject dependency
	inject.RegisterOrFail("serviceProvider", (*sp.ServiceProvider)(nil))
	err := sp.Init()
	if err != nil {
		gd.Crashf("init package sp fail,err=%v", err)
	}
}
