package repository

import "github.com/danielthank/exchat-server/domain/model"

type ProfileRepository interface {
	Create(profile *model.Profile, accessToken string) error
}
