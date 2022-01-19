package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger : 將 gin log 記錄到文件上
func Logger(logger *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()

		// 處理 request
		c.Next()

		// 結束時間
		endTime := time.Now()

		// 執行時間
		latencyTime := endTime.Sub(startTime)

		// request 方式
		reqMethod := c.Request.Method

		// request router
		reqURI := c.Request.RequestURI

		// request IP
		clientIP := c.ClientIP()

		// response 狀態碼
		statusCode := c.Writer.Status()

		// log 格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
