package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/danielthank/exchat-server/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	authHandler := &AuthHandler{
		authUsecase: authUsecase,
	}
	return authHandler
}

func (handler *AuthHandler) Login(c *gin.Context) {
	b := make([]byte, 1024)
	sha256 := sha256.Sum256(b)
	state := hex.EncodeToString(sha256[:])

	err := handler.authUsecase.SaveState(state, c.Request, c.Writer)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	url, err := handler.authUsecase.GetAuthCodeURL(state)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.Redirect(http.StatusFound, url)
}

func (handler *AuthHandler) Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	if err := handler.authUsecase.CheckState(state, c.Request); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
	}
	accessToken, err := handler.authUsecase.GetAccessToken(state, code)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
	}
	log.Println(accessToken)
}

func (handler *AuthHandler) Logout(c *gin.Context) {

}
