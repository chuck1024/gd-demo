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

// @Summary 测试一下
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body user.DemoTestReq true "请求参数" required
// @Router /v1/test [post]
// @Success 200 {object} user.DemoTestResp
func DemoTest(c *gin.Context, req *user.DemoTestReq) (code int, message string, err error, ret *user.DemoTestResp) {
	ret = &user.DemoTestResp{
		Msg: "test success ok!!!",
	}

	return http.StatusOK, "ok", nil, ret
}

// @Summary 获取用户信息
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body user.GetUserInfoReq true "请求参数" required
// @Router /v1/getUserInfo [post]
// @Success 200 {object} user.GetUserInfoRes
func GetUserInfo(c *gin.Context, req *user.GetUserInfoReq) (code int, message string, err error, ret *user.GetUserInfoRes) {
	if req.Passport == "" {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	ret, err = user.GetUserInfo(req.Passport)
	if err != nil {
		dlog.Error("user.GetUserInfo occur err:%v", err)
		return
	}

	return http.StatusOK, "ok", nil, ret
}
