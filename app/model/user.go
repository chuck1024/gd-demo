/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package model

import (
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/databases/mysqldb"
	"time"
)

var (
	useTableName = "user"
)

type User struct {
	Id       uint64 `json:"id" mysqlField:"id"`
	Passport string `json:"passport" mysqlField:"passport"`
	Password uint64 `json:"password" mysqlField:"password"`
	Nickname string `json:"nickname" mysqlField:"nickname"`
	CreateTs int64  `json:"create_time" mysqlField:"create_time"`
	UpdateTs int64  `json:"update_time" mysqlField:"update_time"`
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
		UpdateTs: time.Now().Unix(),
	}

	if err := u.MysqlClient.Add(useTableName, insert, true); err != nil {
		gd.Error("user Insert occur err:%s", err)
		return err
	}

	return nil
}

func (u *UserDao) Update(passport string, password uint64, nickName string) error {
	where := make(map[string]interface{}, 0)
	where["passport"] = passport

	updateData := &User{
		Password: password,
		Nickname: nickName,
		UpdateTs: time.Now().Unix(),
	}

	updateFields := []string{"password", "nickName", "update_time"}

	err := u.MysqlClient.Update(useTableName, updateData, where, updateFields)
	if err != nil {
		gd.Error("Update occur err:%s", err)
		return err
	}
	return nil
}

func (u *UserDao) Query(passport string) (*User, error) {
	query := "select ? from user where passport = ?"

	data, err := u.MysqlClient.Query((*User)(nil), query, passport)
	if err != nil {
		gd.Error("user Query occur err:%s", err)
		return nil, err
	}

	if data == nil {
		gd.Info("user Query passport[%v] is nil", passport)
		return nil, nil
	}

	gd.Debug("user Query value:%v", data.(*User))
	return data.(*User), nil
}
