package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AgentLoginFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "testing agent login"})
}

func AgentQiangZhuangFunc(c *gin.Context) {

}

func AgentYaDaFunc(c *gin.Context) {

}

func AgentYaXiaoFunc(c *gin.Context) {

}
