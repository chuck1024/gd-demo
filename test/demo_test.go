/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package test

import (
	"fmt"
	"github.com/chuck1024/gd"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	host := "http://127.0.0.1:10240"
	client := gd.NewHttpClient(3*time.Second, host)
	if err := client.Start(); err != nil {
		return
	}

	Convey("test ping", t, func() {
		response, body, err := client.Method("POST", "/demo/v1/test", nil, "")
		So(err, ShouldBeNil)
		So(response.StatusCode, ShouldEqual, 200)
		fmt.Println(body)
	})

	Convey("test login", t, func() {
		response, body, err := client.Method("POST", "/demo/v1/login", nil, "{\"passport\":\"cc\",\"password\":1233456}")
		So(err, ShouldBeNil)
		So(response.StatusCode, ShouldEqual, 200)
		fmt.Println(body)
	})
}
