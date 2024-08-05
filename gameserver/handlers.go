package gameserver

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.Engine, h *Hub) {
	apiRouteGroup := rg.Group("/api")

	roomApis := apiRouteGroup.Group("/rooms")
	roomApis.GET("/", h.ListRooms)
	roomApis.POST("/", h.CreateRoom)
	roomApis.GET("/:roomId", h.JoinRoom)
}
