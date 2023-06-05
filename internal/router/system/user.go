package system

import (
	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/api"
)

type UserRouter struct{}

// 用户路由组, 有鉴权
func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("info", userApi.GetUserInfo)
		userRouter.POST("info", userApi.UpdateUserInfo)
		userRouter.PUT("changepwd", userApi.ChangeUserPassword)
		userRouter.PUT("changephone", userApi.ChangeUserPhone)
		userRouter.PUT("recoverpwd", userApi.RecoverUserPassword)
		userRouter.DELETE("logout", userApi.UserLogout)
		userRouter.DELETE("logoff", userApi.UserLogoff)
	}
}

// 用户路由组, 没有鉴权
func (u *UserRouter) InitUserRouterWithOutJwt(Router *gin.RouterGroup) {
	userRouter := Router.Group("")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", userApi.UserLogin)
		userRouter.POST("register", userApi.UserRegister)
	}
}
