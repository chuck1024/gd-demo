/**
 * Copyright 2020 gd-demo Author. All rights reserved.
 * Author: Chuck1024
 */

package boot

import (
	"gd-demo/app/route"
	"github.com/chuck1024/gd"
)

func Run() {
	// init gd
	d := gd.Default()

	// init inject
	Inject()

	// route register
	route.Register(d)

	// gd run
	if err := d.Run(); err != nil {
		gd.Crashf("gd-demo run occur err:%v", err)
		return
	}
}
