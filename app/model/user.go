/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package model

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/mysqldb"
	"time"
)

type User struct {
	Id       uint64 `json:"id" mysqlField:"id"`
	Passport string `json:"passport" mysqlField:"passport"`
	Password uint64 `json:"password" mysqlField:"password"`
	Nickname string `json:"nickname" mysqlField:"nickname"`
	CreateTs int64  `json:"create_time" mysqlField:"create_time"`
}

type UserDao struct {
	MysqlClient *mysqldb.MysqlClient `inject:"mysqlClient"`
}

func (u *UserDao) Start() error {
	return u.MysqlClient.Start()
}

func (u *UserDao) Insert(passport string, password uint64, nickName string) error {
	insert := &User{
		Passport: passport,
		Password: password,
		Nickname: nickName,
		CreateTs: time.Now().Unix(),
	}

	if err := u.MysqlClient.Add("user", insert, true); err != nil {
		dlog.Error("user Insert occur err:%s", err)
		return err
	}

	return nil
}

func (u *UserDao) Query(passport string) (*User, error) {
	query := "select ? from user where passport = ?"

	data, err := u.MysqlClient.Query((*User)(nil), query, passport)
	if err != nil {
		dlog.Error("user Query occur err:%s", err)
		return nil, err
	}

	if data == nil {
		dlog.Error("user Query occur err:%s", err)
		return nil, err
	}

	dlog.Debug("user Query value:%v", data.(*User))
	return data.(*User), nil
}
