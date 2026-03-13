package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/opencodeai/srv/api-getaway/middleware"
	"github.com/yuhang-jieke/orderai/srv/api-getaway/handler"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 全局限流中间件
	r.Use(middleware.GlobalRateLimitMiddleware())

	// 指标收集中间件
	r.Use(middleware.MetricsMiddleware())
	r.POST("orders", middleware.RateLimitMiddleware(), handler.OrderAdd)
	r.GET("orders/:id", middleware.RateLimitMiddleware(), handler.GetId)
	r.DELETE("orders/:id", middleware.RateLimitMiddleware(), handler.DelOrder)
	r.PUT("orders/:id", middleware.RateLimitMiddleware(), handler.UpdateId)
	return r
}
