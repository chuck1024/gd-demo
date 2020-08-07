/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package model

import (
	"encoding/json"
	"github.com/chuck1024/redisdb"
)

const (
	sessionExpiredTs = 30 * 60
)

type Session struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
	Nickname string `json:"nickname"`
}

type SessionCache struct {
	RedisConfig *redisdb.RedisConfig
	redisPool   *redisdb.RedisPool
}

func (c *SessionCache) Start() (err error) {
	c.redisPool, err = redisdb.NewRedisPools(c.RedisConfig)
	if err != nil {
		return
	}
	return nil
}

func (c *SessionCache) Set(sessionId string, v *Session) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = c.redisPool.SetNX(sessionId, string(data), sessionExpiredTs)
	if err != nil {
		return err
	}

	return nil
}

func (c *SessionCache) Get(sessionId string) (v *Session, err error) {
	value, err := c.redisPool.Get(sessionId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *SessionCache) Del(sessionId string) error {
	err := c.redisPool.Del(sessionId)
	if err != nil {
		return err
	}
	return nil
}
