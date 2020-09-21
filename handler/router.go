package handler

import "github.com/gin-gonic/gin"

func AddWSRoutes(router *gin.Engine, wsHandler *WSHandler) {
	router.GET("/ws", wsHandler.Handle)
}

func AddAuthRoutes(router *gin.Engine, authHandler *AuthHandler) {
	group := router.Group("/auth")
	{
		group.GET("/login", authHandler.Login)
		group.GET("/callback", authHandler.Callback)
		group.GET("/logout", authHandler.Logout)
	}
}
