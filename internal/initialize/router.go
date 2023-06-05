package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toma-photo/internal/docs"
	"github.com/toma-photo/internal/middleware"

	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// 初始化 swagger 服务
	initSwaggerServer(Router)

	// Router Server
	basicRouters := router.RouterGroupApp.Basic
	systemRouters := router.RouterGroupApp.System

	publicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		basicRouters.InitAPiRouter(publicGroup) // 健康监测
		systemRouters.InitBasicRouter(publicGroup)
		systemRouters.InitUserRouterWithOutJwt(publicGroup)
	}

	privateGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth())
	{
		systemRouters.InitUserRouter(privateGroup)
	}

	global.ZAP_LOG.Info("router register success")
	return Router
}

// 初始化 swagger 服务, 只有在非发版服务才开启swagger
func initSwaggerServer(Router *gin.Engine) {
	switch gin.Mode() {
	case gin.ReleaseMode:
		return
	default:
		// swagger 初始化
		docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
		Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.ZAP_LOG.Info("register swagger handler")
	}
}
