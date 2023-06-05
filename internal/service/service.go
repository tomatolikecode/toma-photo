package service

import (
	"github.com/toma-photo/internal/service/basic"
	"github.com/toma-photo/internal/service/system"
)

/*
	服务层, 实现具体的服务
*/

type ServiceGroup struct {
	BasicServiceGroup  basic.ServiceGroup
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
