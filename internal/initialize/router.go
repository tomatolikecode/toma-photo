package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toma-photo/internal/global"
)

func Routers() *gin.Engine {
	router := gin.Default()
	// swagger
	router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.ZAP_LOG.Info("register swagger handler")
	publicGroup := router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// privateGroup := router.Group(global.CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	global.ZAP_LOG.Info("router register success")
	return router
}
