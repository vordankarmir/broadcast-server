package controllers

import (
	"github.com/gin-gonic/gin"
)

func RegisterCacheRoutes(r *gin.Engine, origin string) {
	api := r.Group("/")

	api.Any("/*action", func(c *gin.Context) {
	})
}
