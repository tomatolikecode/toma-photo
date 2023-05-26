package service

import "github.com/toma-photo/internal/service/basic"

/*
	服务层, 实现具体的服务
*/

type ServiceGroup struct {
	BasicServiceGroup basic.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
