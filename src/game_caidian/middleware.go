package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireLoginMiddleware(c *gin.Context) {
	//debug only
	c.AbortWithStatus(http.StatusForbidden)
}
