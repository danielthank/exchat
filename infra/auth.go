package infra

import (
	"log"

	"github.com/danielthank/exchat-server/domain/model"
	"github.com/danielthank/exchat-server/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	profileRepository struct {
		*SqlHandler
	}
	profileGorm struct {
		gorm.Model
		UserID      string `gorm:"uniqueIndex;size:40"`
		DisplayName string
		PictureURL  string
		AccessToken string
	}
)

func NewProfileRepository(sqlHandler *SqlHandler) repository.ProfileRepository {
	profileRepository := &profileRepository{sqlHandler}
	profileRepository.AutoMigrate(&profileGorm{})
	return profileRepository
}

func (t *profileRepository) Create(profile *model.Profile) error {
	log.Println("creating")
	if err := t.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"display_name", "picture_url", "access_token"}),
	}).Create(&profileGorm{
		UserID:      profile.UserID,
		DisplayName: profile.DisplayName,
		PictureURL:  profile.PictureURL,
		AccessToken: profile.AccessToken,
	}).Error; err != nil {
		return err
	}
	return nil
}
