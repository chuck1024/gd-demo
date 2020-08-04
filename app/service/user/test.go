/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package user

type DemoTestReq struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type DemoTestResp struct {
	Msg  string `json:"msg"`
}
