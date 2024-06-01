package handlers

import (
	"github.com/gin-gonic/gin"
	//"github.com/nagurumalab/gocards/gocards"
)

func addSessionRoutes(rg *gin.RouterGroup) {
	rg.POST("/", createSession)
}

func createSession(c *gin.Context) {
	//gocards.Session{}
}
