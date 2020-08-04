/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package api

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/gd-demo/app/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DemoTest(c *gin.Context, req *user.DemoTestReq)(code int, message string, err error, ret *user.DemoTestResp) {
	dlog.Debug("DemoTest req:%v", req)

	ret = &user.DemoTestResp{
		Msg: "test success ok!!!",
	}

	return http.StatusOK, "ok", nil, ret
}