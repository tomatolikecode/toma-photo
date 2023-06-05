package router

import (
	"github.com/toma-photo/internal/router/basic"
	"github.com/toma-photo/internal/router/system"
)

/*
	路由模块
*/

type RouterGroup struct {
	Basic  basic.RouterGroup
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
