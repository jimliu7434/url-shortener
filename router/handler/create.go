package handler

import "github.com/gin-gonic/gin"

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: generate new uid
		// TODO: insert into db
	}
}
