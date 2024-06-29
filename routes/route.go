package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"ops_tool/conf"
	"ops_tool/controller"
	// 千万不要忘了导入把你上一步生成的docs
	"ops_tool/public"
)

func SetupRoute() *gin.Engine {
	gin.SetMode(conf.Conf.App.Mode)
	r := gin.Default()

	// r.Use(middleware.GinLogger(), middleware.GinRecovery(true)) 添加gin自定义日志中间件和获取Panic中间件
	opsTool := r.Group("/jenkins")
	{
		// jenkins构建状态钉钉通知api
		opsTool.POST("/jkDingTalk", controller.JksDingTalk)
	}

	// 建立swagger文档
	r.GET("/swagger/*any", func(c *gin.Context) {
		gs.WrapHandler(swaggerFiles.Handler)
		public.Log.Info(c.Request.URL.Path + " 加载成功")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": controller.CodeInvalidPath,
			"msg":  controller.CodeInvalidPath.Msg(),
		})
		public.Log.Error(c.Request.URL.Path + " 路径不存在")
	})

	public.Log.Info("初始化路由完成！")
	return r
}
