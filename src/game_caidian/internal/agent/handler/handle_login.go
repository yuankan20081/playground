package handler

import (
	"github.com/gin-gonic/gin"
	"grpc/service/agent"
	"log"
	"net/http"
)

func (h *Handler) OnLogin(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var request = &agent.LoginRequest{
		Token: token,
	}

	iReply, err := h.p.Call(h.ctx, "Login", request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if reply, ok := iReply.(*agent.LoginReply); ok {
		if reply.Sucess {
			c.JSON(http.StatusOK, gin.H{
				"code":    0, //success
				"message": "",
				"token":   "some random token",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    1, //fail
				"message": "",
				"token":   "bad token",
			})
		}
		return
	} else {
		log.Println("wrong type, not a loginreply")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
