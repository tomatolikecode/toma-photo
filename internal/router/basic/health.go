package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/api"
)

type HealthRouter struct{}

func (h *HealthRouter) InitAPiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("")
	baseApi := api.ApiGroupApp.BasicApiGroup.HealthApi
	{
		apiRouter.GET("health", baseApi.Health) // 创建Api
	}
}
