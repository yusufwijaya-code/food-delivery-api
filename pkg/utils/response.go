package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

func JSONSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"data": data})
}
