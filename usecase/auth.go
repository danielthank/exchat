package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/danielthank/exchat-server/domain/model"
	"github.com/danielthank/exchat-server/domain/repository"
	"github.com/danielthank/exchat-server/domain/service"
	"golang.org/x/oauth2"
)

type AuthUsecase interface {
	GetAuthCodeURL(state string) (string, error)
	GetAccessToken(state string, code string) (string, error)
	SaveState(state string, r *http.Request, w http.ResponseWriter) error
	CheckState(state string, r *http.Request) error
	DeleteSesssion(r *http.Request, w http.ResponseWriter) error
	GetProfileByAccessToken(accessToken string) (*model.Profile, error)
	PersistProfile(profile *model.Profile) error
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

func (usecase *authUsecase) DeleteSesssion(r *http.Request, w http.ResponseWriter) error {
	if err := usecase.sessionService.Delete(r, w, usecase.sessionID); err != nil {
		return err
	}
	return nil
}

func (usecase *authUsecase) GetProfileByAccessToken(accessToken string) (*model.Profile, error) {
	profile := &model.Profile{}
	profile.AccessToken = accessToken

	req, err := http.NewRequest("GET", usecase.profileEndpoint, nil)
	if err != nil {
		return nil, errors.New("failed creating request")
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(profile)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return profile, nil
}

func (usecase *authUsecase) PersistProfile(profile *model.Profile) error {
	if err := usecase.profileRepository.Create(profile); err != nil {
		return err
	}
	return nil
}
