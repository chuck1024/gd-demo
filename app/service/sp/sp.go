/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package sp

import (
	"gd-demo/app/model"
	"github.com/chuck1024/gd/runtime/inject"
	"github.com/go-errors/errors"
)

var p *ServiceProvider

type ServiceProvider struct {
	UserModel    *model.UserDao      `inject:"UserDao"`
	SessionCache *model.SessionCache `inject:"SessionCache"`
}

func Init() error {
	f, ok := inject.Find("serviceProvider")
	if !ok {
		return errors.New("serviceProvider not found")
	}

	ff, ok := f.(*ServiceProvider)
	if !ok {
		return errors.New("serviceProvider not valid")
	}
	p = ff
	return nil
}

func Get() *ServiceProvider {
	return p
}
