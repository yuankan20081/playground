package main

import (
	"flag"
	"game_caidian/internal/agent/handler"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func main() {
	var debugMode = flag.Bool("debug", false, "-debug=true to debug")

	flag.Parse()

	if !*debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	agentHandler := handler.New(context.Background())

	// TODO: add handlers
	agentGroup := engine.Group("/api_agent")
	{
		agentGroup.POST("/login", agentHandler.OnLogin)
	}

	engine.Run(":9090")
}
