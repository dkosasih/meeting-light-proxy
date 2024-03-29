package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestedApiVersion := c.Param("ver")
		if header := c.GetHeader("x-version"); len(header) > 0 {
			requestedApiVersion = header
		}

		if requestedApiVersion != "" {
			c.Set("version", requestedApiVersion)
		}
		c.Next()
	}
}
