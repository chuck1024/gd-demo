/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package middleware

import (
	"github.com/chuck1024/gd-demo/app/service/sp"
	"github.com/chuck1024/gd/net/dhttp"
	"github.com/chuck1024/gl"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SessionIdCookie = "sess"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(SessionIdCookie)
		if err != nil || sessionId == ""{
			c.Abort()
			dhttp.Return(c, http.StatusForbidden,"no auth", err, nil)
			return
		}

		v, err := sp.Get().SessionCache.Get(sessionId)
		if err != nil {
			dhttp.Return(c, http.StatusBadRequest,"cookie expire", err, nil)
		}

		gl.Set(SessionIdCookie, v)
		c.Next()
	}
}
