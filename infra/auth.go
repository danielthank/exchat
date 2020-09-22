package infra

import (
	"github.com/danielthank/exchat-server/domain/model"
	"github.com/danielthank/exchat-server/domain/repository"
	"gorm.io/gorm"
)

type (
	profileRepository struct {
		*SqlHandler
	}
	profileGorm struct {
		gorm.Model
		UserID      string
		DisplayName string
		PictureURL  string
		AccessToken string
	}
)

func NewProfileRepository(sqlHandler *SqlHandler) repository.ProfileRepository {
	profileRepository := &profileRepository{sqlHandler}
	profileRepository.Conn.AutoMigrate(&profileGorm{})
	return profileRepository
}

func (t *profileRepository) Create(profile *model.Profile) error {
	if err := t.Conn.Create(&profileGorm{
		UserID:      profile.UserID,
		DisplayName: profile.DisplayName,
		PictureURL:  profile.PictureURL,
		AccessToken: profile.AccessToken,
	}).Error; err != nil {
		return err
	}
	return nil
}
