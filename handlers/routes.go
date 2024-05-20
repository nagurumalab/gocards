package handlers

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	sessionRoutes := rg.Group("/session")
	addSessionRoutes(sessionRoutes)
}
