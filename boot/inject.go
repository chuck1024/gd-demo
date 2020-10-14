/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"gd-demo/app/service/sp"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/databases/mysqldb"
	"github.com/chuck1024/gd/databases/redisdb"
	"github.com/chuck1024/gd/runtime/inject"
)

func Inject() {
	// inject demoMysqlClient and init mysql client
	inject.RegisterOrFail("demoMysqlClient",(*mysqldb.MysqlClient)(&mysqldb.MysqlClient{
		DataBases: "demo",
	}))

	// inject demoRedisClient and init redis pool client
	inject.RegisterOrFail("demoRedisClient", (*redisdb.RedisPoolClient)(&redisdb.RedisPoolClient{
		PoolName: "demo",
	}))

	// inject dependency
	inject.RegisterOrFail("serviceProvider", (*sp.ServiceProvider)(nil))
	err := sp.Init()
	if err != nil {
		gd.Crashf("init package sp fail,err=%v", err)
	}
}
