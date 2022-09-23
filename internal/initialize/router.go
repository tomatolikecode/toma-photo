package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/global"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOGGER.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")
	// systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("v1")
	{
		// 心跳检测 健康检测
		PublicGroup.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		//systemRouter.InitUserBaseRouter(PublicGroup) // 用户基本接口,注册&登录

	}
	PrivateGroup := Router.Group("v2")
	//PrivateGroup.Use(middleware.JWTAuth())
	{
		PrivateGroup.GET("private", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		//	systemRouter.InitDemo(PrivateGroup)         // demo
		//	systemRouter.InitSystemRouter(PrivateGroup) // system相关路由
		//	systemRouter.InitUserRouter(PrivateGroup)   // 用户操作接口
	}
	global.LOGGER.Info("router register success")
	return Router
}
