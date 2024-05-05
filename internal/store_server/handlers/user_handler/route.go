package user_handler

import "github.com/gin-gonic/gin"

func RegisterAuthRoute(router *gin.RouterGroup, handlerClient *Client) {
	handler := handlerClient.AuthHandler

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
}
