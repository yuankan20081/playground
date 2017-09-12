package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	var debugMode = flag.Bool("debug", false, "-debug=true to debug")

	flag.Parse()

	if !*debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	// TODO: add handlers
	agentGroup := engine.Group("/api_agent")
	{
		agentGroup.POST("/login", AgentLoginFunc)

		agentRequireLoginGroup := agentGroup.Group("/", RequireLoginMiddleware)
		{
			agentRequireLoginGroup.POST("/qiangzhuang", AgentQiangZhuangFunc)
			agentRequireLoginGroup.POST("/yada", AgentYaDaFunc)
			agentRequireLoginGroup.POST("/yaxiao", AgentYaXiaoFunc)
		}
	}

	engine.Run(":9090")
}
