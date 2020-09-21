package infra

import (
	"github.com/danielthank/exchat-server/domain/model"
	"github.com/danielthank/exchat-server/domain/repository"
	"gorm.io/gorm"
)

type (
	ProfileRepository struct {
		*SqlHandler
	}
	ProfileGorm struct {
		gorm.Model
		UserID      string
		DisplayName string
		PictureURL  string
		AccessToken string
	}
)

func NewProfileRepository(sqlHandler *SqlHandler) repository.ProfileRepository {
	profileRepository := &ProfileRepository{sqlHandler}
	profileRepository.Conn.AutoMigrate(&ProfileGorm{})
	return profileRepository
}

func (t *ProfileRepository) Create(profile *model.Profile, accessToken string) error {
	if err := t.Conn.Create(&ProfileGorm{
		UserID:      profile.UserID,
		DisplayName: profile.DisplayName,
		PictureURL:  profile.PictureURL,
		AccessToken: accessToken,
	}).Error; err != nil {
		return err
	}
	return nil
}
