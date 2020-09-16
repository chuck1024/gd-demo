/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package api

import (
	"github.com/chuck1024/gd-demo/app/model"
	"github.com/chuck1024/gd-demo/app/service/middleware"
	"github.com/chuck1024/gd-demo/app/service/user"
	"github.com/chuck1024/gd/derror"
	"github.com/chuck1024/gd/dlog"
	"github.com/chuck1024/gd/runtime/gl"
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
// @Param cookie header string true "cookie" required
// @Router /v1/getUserInfo [get]
// @Success 200 {object} user.GetUserInfoRes
func GetUserInfo(c *gin.Context, req interface{}) (code int, message string, err error, ret *user.GetUserInfoRes) {
	var passport string
	v, ok := gl.Get(middleware.SessionIdCookie)
	if ok {
		passport = v.(*model.Session).Passport
	}

	if passport == "" {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	ret, err = user.GetUserInfo(passport)
	if err != nil {
		v, ok := err.(*derror.CodeError)
		if ok {
			return v.Code(), v.Error(), nil, nil
		}

		dlog.Error("user.GetUserInfo occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	return http.StatusOK, "ok", nil, ret
}

// @Summary 注册或更新用户信息
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body user.RegisterOrUpdateReq true "请求参数" required
// @Router /v1/register [post]
// @Success 200
func RegisterOrUpdate(c *gin.Context, req *user.RegisterOrUpdateReq) (code int, message string, err error, ret interface{}) {
	if req.Passport == "" || req.Password == 0 || req.Nickname == "" {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	err = user.RegisterOrUpdate(req.Passport, req.Nickname, req.Password)
	if err != nil {
		dlog.Error("RegisterOrUpdate user occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	return http.StatusOK, "ok", nil, nil
}

// @Summary 用户登录
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body user.LoginReq true "请求参数" required
// @Router /v1/login [post]
// @Success 200 {object} user.LoginRes
func Login(c *gin.Context, req *user.LoginReq) (code int, message string, err error, ret *user.LoginRes) {
	if req.Passport == "" || req.Password == 0 {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	ret, err = user.Login(req.Passport, req.Password)
	if err != nil {
		v, ok := err.(*derror.CodeError)
		if ok {
			return v.Code(), v.Error(), nil, nil
		}

		dlog.Error("Login user occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	c.SetCookie(middleware.SessionIdCookie, ret.SessionId, 1, "/", "localhost", false, true)
	return http.StatusOK, "ok", nil, ret
}
