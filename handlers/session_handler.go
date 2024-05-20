package handlers

import "github.com/gin-gonic/gin"

func addSessionRoutes(rg *gin.RouterGroup) {
	rg.POST("/", createSession)
}

func createSession(c *gin.Context) {

}
