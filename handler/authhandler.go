package handler

import (
	"github.com/danielthank/exchat-server/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	authHandler := &AuthHandler{
		authUsecase: authUsecase,
	}
	return authHandler
}

func (authHandler *AuthHandler) Login(c *gin.Context) {

}

func (authHandler *AuthHandler) Callback(c *gin.Context) {

}

func (authHandler *AuthHandler) Logout(c *gin.Context) {

}
