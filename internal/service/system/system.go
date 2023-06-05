package system

import "github.com/toma-photo/internal/global"

/*
	系统功能服务
*/

var Db = global.DB()

type ServiceGroup struct {
	UserService
}
