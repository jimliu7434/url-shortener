package router

import (
	"url-shortener/common/log"

	"github.com/gin-gonic/gin"

	handler "url-shortener/router/handler"
	middleware "url-shortener/router/middleware"
)

// SetupRouter :
func SetupRouter(isDebugMode bool) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// log
	r.Use(middleware.Logger(log.AccLog))

	r.Use(func(c *gin.Context) {
		c.Set("IsDebugMode", isDebugMode)
	})

	rAPI := r.Group("/api/v1")
	{
		rAPI.POST("/urls", handler.Create())
	}

	r.GET("/:uid", handler.Redirect())

	return r
}
