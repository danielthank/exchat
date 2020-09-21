package usecase

import (
	"os"

	"github.com/danielthank/exchat-server/domain/repository"
	"golang.org/x/oauth2"
)

type AuthUsecase interface {
}

type authUsecase struct {
	profileRepository *repository.ProfileRepository
	lineOAuthConfig   *oauth2.Config
	profileEndpoint   string
}

func NewAuthUsecase(profileRepository *repository.ProfileRepository) *authUsecase {
	authUsecase := &authUsecase{
		profileRepository: profileRepository,
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
	}
	return authUsecase
}

func (usecase *authUsecase) AuthCodeURL(state string) error {
	return nil
}
