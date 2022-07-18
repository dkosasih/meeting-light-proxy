package middleware

import (
	"github.com/gin-gonic/gin"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestedApiVersion := c.Param("ver")
		if header := c.GetHeader("x-version"); len(header) > 0 {
			requestedApiVersion = header
		}

		if requestedApiVersion == "1" || requestedApiVersion == "1.0" {
			c.Set("version", "1")
		}
		c.Next()
	}
}
