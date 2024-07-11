package router

import (
	"admin-template-go/common/config"
	"admin-template-go/middleware"
	"github.com/gin-gonic/gin"
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

}
