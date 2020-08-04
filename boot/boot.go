/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"github.com/chuck1024/dlog"
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd-demo/route"
)

func Run() {
	d := gd.Default()

	route.Register(d)

	if err := d.Run(); err != nil {
		dlog.Crashf("gd-demo run occur err:%v", err)
		return
	}
}
