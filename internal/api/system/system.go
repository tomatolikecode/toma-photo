package system

import "github.com/toma-photo/internal/service"

/*
	系统功能模块
*/

type ApiGroup struct {
	UserApi
	BasicApi
}

var (
	userServce = service.ServiceGroupApp.SystemServiceGroup.UserService
)
