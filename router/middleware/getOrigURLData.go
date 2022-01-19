package middleware

import "github.com/gin-gonic/gin"

func GetOrigURLData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get data from db
	}
}
