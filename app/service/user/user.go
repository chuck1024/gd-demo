/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package user

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/gd-demo/app/service/sp"
)

type DemoTestReq struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type DemoTestResp struct {
	Msg string `json:"msg"`
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

	return &GetUserInfoRes{
		Passport: userInfo.Passport,
		Password: userInfo.Password,
		Nickname: userInfo.Nickname,
	}, nil
}
