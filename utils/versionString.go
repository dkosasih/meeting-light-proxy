package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetVersionString(c *gin.Context) string {
	version, exists := c.Get("version")

	if !exists {
		return ""
	}

	return strings.Split(version.(string), ".")[0]
}
