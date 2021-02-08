/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package api

import (
	"gd-demo/app/domain"
	"gd-demo/app/model"
	"gd-demo/app/service/middleware"
	"gd-demo/app/service/user"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/derror"
	"github.com/chuck1024/gd/runtime/gl"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 测试一下
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body domain.DemoTestReq true "请求参数" required
// @Router /demo/v1/test [get]
// @Success 200 {object} domain.DemoTestRsp
func DemoTest(c *gin.Context, req *domain.DemoTestReq) (code int, message string, err error, ret *domain.DemoTestRsp) {
	ret = &domain.DemoTestRsp{
		Msg: "test success ok!!!",
	}

	return http.StatusOK, "ok", nil, ret
}

// @Summary 获取用户信息
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param cookie header string true "cookie" required
// @Router /demo/v1/getUserInfo [get]
// @Success 200 {object} domain.GetUserInfoRsp
func GetUserInfo(c *gin.Context, req interface{}) (code int, message string, err error, ret *domain.GetUserInfoRsp) {
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

		gd.Error("user.GetUserInfo occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	return http.StatusOK, "ok", nil, ret
}

// @Summary 注册或更新用户信息
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body domain.RegisterOrUpdateReq true "请求参数" required
// @Router /demo/v1/register [post]
// @Success 200
func RegisterOrUpdate(c *gin.Context, req *domain.RegisterOrUpdateReq) (code int, message string, err error, ret interface{}) {
	if req.Passport == "" || req.Password == 0 || req.Nickname == "" {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	err = user.RegisterOrUpdate(req.Passport, req.Nickname, req.Password)
	if err != nil {
		gd.Error("RegisterOrUpdate user occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	return http.StatusOK, "ok", nil, nil
}

// @Summary 用户登录
// @Tags 用户
// @Accept application/json
// @Produce json
// @Param entity body domain.LoginReq true "请求参数" required
// @Router /demo/v1/login [post]
// @Success 200 {object} domain.LoginRsp
func Login(c *gin.Context, req *domain.LoginReq) (code int, message string, err error, ret *domain.LoginRsp) {
	if req.Passport == "" || req.Password == 0 {
		return http.StatusBadRequest, "passport null", nil, nil
	}

	ret, err = user.Login(req.Passport, req.Password)
	if err != nil {
		v, ok := err.(*derror.CodeError)
		if ok {
			return v.Code(), v.Error(), nil, nil
		}

		gd.Error("Login user occur err:%v", err)
		return http.StatusInternalServerError, err.Error(), err, nil
	}

	c.SetCookie(middleware.SessionIdCookie, ret.SessionId, 1, "/", "localhost", false, true)
	return http.StatusOK, "ok", nil, ret
}
