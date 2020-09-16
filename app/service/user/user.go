/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package user

import (
	"github.com/chuck1024/gd-demo/app/model"
	"github.com/chuck1024/gd-demo/app/service/sp"
	"github.com/chuck1024/gd/derror"
	"github.com/chuck1024/gd/dlog"
	"github.com/chuck1024/gd/utls"
	"net/http"
)

type DemoTestReq struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type DemoTestResp struct {
	Msg string `json:"msg"`
}

type RegisterOrUpdateReq struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
	Nickname string `json:"nickname"`
}

func RegisterOrUpdate(passport, nickname string, password uint64) error {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		dlog.Error("RegisterOrUpdate UserModel.Query occur err:%v", err)
		return err
	}

	if userInfo != nil {
		err = sp.Get().UserModel.Update(passport, password, nickname)
		if err != nil {
			dlog.Error("RegisterOrUpdate UserModel.Update occur er:%v", err)
			return err
		}
		return nil
	}

	if err = sp.Get().UserModel.Insert(passport, password, nickname); err != nil {
		dlog.Error("RegisterOrUpdate UserModel.Insert occur er:%v", err)
		return err
	}
	return nil
}

type LoginReq struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
}

type LoginRes struct {
	SessionId string `json:"sessionId"`
}

func Login(passport string, password uint64) (*LoginRes, error) {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		dlog.Error("Login UserModel.Query occur err:%v", err)
		return nil, err
	}

	if userInfo == nil {
		dlog.Info("passport[%v] null", passport)
		return nil, derror.NewCodeError(http.StatusBadRequest, "no data")
	}

	if password != userInfo.Password {
		return nil, derror.NewCodeError(http.StatusBadRequest, "password error")
	}

	sessionId := passport + "__" + utls.RandString(20)
	sess := &model.Session{
		Passport: userInfo.Passport,
		Password: userInfo.Password,
		Nickname: userInfo.Nickname,
	}

	if err = sp.Get().SessionCache.Set(sessionId, sess); err != nil {
		dlog.Error("Login SessionCache.Set occur err:%v", err)
		return nil, derror.MakeCodeError(http.StatusInternalServerError, err)
	}

	return &LoginRes{SessionId: sessionId}, nil
}

type GetUserInfoReq struct {
	Passport string `json:"passport"`
}

type GetUserInfoRes struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
	Nickname string `json:"nickname"`
}

func GetUserInfo(passport string) (*GetUserInfoRes, error) {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		dlog.Error("GetUserInfo UserModel.Query occur err:%v", err)
		return nil, err
	}

	if userInfo == nil {
		return nil, derror.NewCodeError(http.StatusBadRequest, "no data")
	}

	return &GetUserInfoRes{
		Passport: userInfo.Passport,
		Password: userInfo.Password,
		Nickname: userInfo.Nickname,
	}, nil
}
