/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package route

import (
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/app/api"
	_ "github.com/chuck1024/gd-demo/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/chuck1024/gd-demo/library/middleware"
	"github.com/chuck1024/gd/net/dhttp"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"sync"
)

var (
	initOnce sync.Once
)

func Register(e *gd.Engine) {
	e.HttpServer.SetInit(func(g *gin.Engine) error {
		r := g.Group("")
		r.Use(
			dhttp.GlFilter(),
			dhttp.GroupFilter(),
			dhttp.Logger(),
		)

		return route(e, r)
	})
}

func route(e *gd.Engine, r *gin.RouterGroup) error {
	var ret error
	initOnce.Do(func() {
		// swagger
		ok := e.Config("Swagger", "swagger").MustBool()
		if ok {
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		g := r.Group("v1")
		g.Use(middleware.Cors())

		e.HttpServer.DefaultAddHandler("test", api.DemoTest)
		e.HttpServer.DefaultAddHandler("getUserInfo", api.GetUserInfo)

		for path, fun := range e.HttpServer.DefaultHandlerMap {
			f, err := dhttp.Wrap(fun)
			if err != nil {
				ret = err
				return
			}

			g.GET(path, f)
			g.POST(path, f)
		}
	})
	return ret
}
