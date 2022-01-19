package handler

import (
	"net/http"
	"url-shortener/common/log"
	model "url-shortener/model/redis"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Redirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: record some metrics

		uid := c.Param("uid")

		if uid == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// get data from db
		val, err := model.URL.Get(uid)
		if err != nil {
			if err == redis.Nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			} else {
				log.TraceLog.Infof("[Redirect] redis err: %s when uid= %s", err.Error(), uid)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		// 302 redirect
		c.Redirect(http.StatusFound, val)
	}
}
