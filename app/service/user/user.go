/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package user

import (
	"gd-demo/app/domain"
	"gd-demo/app/model"
	"gd-demo/app/service/sp"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/derror"
	"github.com/chuck1024/gd/utls"
	"net/http"
)

func RegisterOrUpdate(passport, nickname string, password uint64) error {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		gd.Error("RegisterOrUpdate UserModel.Query occur err:%v", err)
		return err
	}

	if userInfo != nil {
		err = sp.Get().UserModel.Update(passport, password, nickname)
		if err != nil {
			gd.Error("RegisterOrUpdate UserModel.Update occur er:%v", err)
			return err
		}
		return nil
	}

	if err = sp.Get().UserModel.Insert(passport, password, nickname); err != nil {
		gd.Error("RegisterOrUpdate UserModel.Insert occur er:%v", err)
		return err
	}
	return nil
}

func Login(passport string, password uint64) (*domain.LoginRsp, error) {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		gd.Error("Login UserModel.Query occur err:%v", err)
		return nil, err
	}

	if userInfo == nil {
		gd.Info("passport[%v] null", passport)
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
		gd.Error("Login SessionCache.Set occur err:%v", err)
		return nil, derror.MakeCodeError(http.StatusInternalServerError, err)
	}

	return &domain.LoginRsp{SessionId: sessionId}, nil
}

func GetUserInfo(passport string) (*domain.GetUserInfoRsp, error) {
	userInfo, err := sp.Get().UserModel.Query(passport)
	if err != nil {
		gd.Error("GetUserInfo UserModel.Query occur err:%v", err)
		return nil, err
	}

	if userInfo == nil {
		return nil, derror.NewCodeError(http.StatusBadRequest, "no data")
	}

	return &domain.GetUserInfoRsp{
		Passport: userInfo.Passport,
		Password: userInfo.Password,
		Nickname: userInfo.Nickname,
	}, nil
}
