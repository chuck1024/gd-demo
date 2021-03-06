/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package route

import (
	"gd-demo/app/api"
	"gd-demo/app/service/middleware"
	_ "gd-demo/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/net/dhttp"
	"github.com/chuck1024/gd/runtime/inject"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"sync"
)

var (
	initOnce sync.Once
)

func Register(e *gd.Engine) {
	inject.RegisterOrFail("httpServerInit",func(g *gin.Engine) error {
		r := g.Group("")
		// swagger
		ok := gd.Config("Swagger", "swagger").MustBool()
		if ok {
			r.GET("/demo/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		r.Use(
			dhttp.GlFilter(),
			dhttp.StatFilter(),
			dhttp.GroupFilter(),
			dhttp.Logger(gd.Config("Server", "serverName").String()),
		)
		return route(e, r)
	})
}

func route(e *gd.Engine, r *gin.RouterGroup) error {
	var ret error
	initOnce.Do(func() {
		g := r.Group("demo/v1")
		g.Use(middleware.Cors())

		e.HttpServer.GET(g, "test", api.DemoTest)
		e.HttpServer.POST(g, "register", api.RegisterOrUpdate)
		e.HttpServer.POST(g, "update", api.RegisterOrUpdate)
		e.HttpServer.POST(g, "login", api.Login)

		g.Use(middleware.Auth())
		e.HttpServer.GET(g, "getUserInfo", api.GetUserInfo)

		if ret = e.HttpServer.CheckHandle(); ret != nil {
			return
		}
	})
	return ret
}
