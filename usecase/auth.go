package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/danielthank/exchat-server/domain/repository"
	"github.com/danielthank/exchat-server/domain/service"
	"golang.org/x/oauth2"
)

type AuthUsecase interface {
	GetAuthCodeURL(state string) (string, error)
	GetAccessToken(state string, code string) (string, error)
	SaveState(state string, r *http.Request, w http.ResponseWriter) error
	CheckState(state string, r *http.Request) error
}

type authUsecase struct {
	profileRepository repository.ProfileRepository
	sessionService    service.SessionService
	lineOAuthConfig   *oauth2.Config
	profileEndpoint   string
	sessionID         string
}

func NewAuthUsecase(profileRepository repository.ProfileRepository, sessionService service.SessionService) AuthUsecase {
	authUsecase := &authUsecase{
		profileRepository: profileRepository,
		sessionService:    sessionService,
		lineOAuthConfig: &oauth2.Config{
			ClientID:     os.Getenv("LINE_CLIENT_ID"),
			ClientSecret: os.Getenv("LINE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("LINE_CALLBACK_URL"),
			Scopes:       []string{"profile", "openid"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
				TokenURL: "https://api.line.me/oauth2/v2.1/token",
			},
		},
		profileEndpoint: "https://api.line.me/v2/profile",
		sessionID:       "session-id",
	}
	return authUsecase
}

func (usecase *authUsecase) GetAuthCodeURL(state string) (string, error) {
	url := usecase.lineOAuthConfig.AuthCodeURL(state)
	return url, nil
}

func (usecase *authUsecase) GetAccessToken(state string, code string) (string, error) {
	token, err := usecase.lineOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return token.AccessToken, nil
}

func (usecase *authUsecase) SaveState(state string, r *http.Request, w http.ResponseWriter) error {
	if err := usecase.sessionService.Save(r, w, usecase.sessionID, map[string]interface{}{"state": state}); err != nil {
		return err
	}
	return nil
}

func (usecase *authUsecase) CheckState(state string, r *http.Request) error {
	session, err := usecase.sessionService.Get(r, usecase.sessionID)
	if err != nil {
		return err
	}
	if session.Values["state"] != state {
		return errors.New("invalid oauth state")
	}
	return nil
}
