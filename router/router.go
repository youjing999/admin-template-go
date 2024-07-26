package router

import (
	"admin-template-go/api/controller"
	"admin-template-go/common/config"
	"admin-template-go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter  初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	router.Use(middleware.Logger())
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)

	jwt := router.Group("/api", middleware.AuthMiddleware())
	{
		jwt.POST("/post/add", controller.CreateSysPost)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/info", controller.GetSysPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)
		jwt.GET("/dept/list", controller.GetSysDeptList)
		jwt.POST("/dept/add", controller.CreateSysDept)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)
	}

}
