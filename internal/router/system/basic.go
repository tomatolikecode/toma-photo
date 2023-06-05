package system

import (
	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/api"
)

type BasicRouter struct{}

func (u *BasicRouter) InitBasicRouter(Router *gin.RouterGroup) {
	basicRouter := Router.Group("basic")
	basicApi := api.ApiGroupApp.SystemApiGroup.BasicApi
	{
		basicRouter.POST("smscode", basicApi.Sms)
	}
}
