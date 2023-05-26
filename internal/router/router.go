package router

import "github.com/toma-photo/internal/router/basic"

/*
	路由模块
*/

type RouterGroup struct {
	Basic basic.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
