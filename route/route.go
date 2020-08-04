/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package route

import (
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/app/api"
	"github.com/chuck1024/gd-demo/library/middleware"
	"github.com/chuck1024/gd/net/dhttp"
	"github.com/gin-gonic/gin"
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
		g := r.Group("v1")
		g.Use(middleware.Cors())

		e.HttpServer.DefaultAddHandler("test", api.DemoTest)

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
