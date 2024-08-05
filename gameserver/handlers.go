package gameserver

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.Engine, h *Hub) {
	apiRouteGroup := rg.Group("/api")

	apiRouteGroup.Group("/rooms")
	apiRouteGroup.GET("/", h.ListRooms)
	apiRouteGroup.POST("/", h.CreateRoom)
	apiRouteGroup.GET("/:roomId", h.JoinRoom
}
