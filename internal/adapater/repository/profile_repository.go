package repository

import (
	"context"
	"desadangdang/internal/core/domain/entity"
	"desadangdang/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProfileInterface interface {
	FetchByIDProfile(ctx context.Context, id int64) (*entity.ProfileEntity, error)
	EditByIDProfile(ctx context.Context, req entity.ProfileEntity) error
}

type profile struct {
	DB *gorm.DB
}

// FetchByIDProfile implements ProfileInterface.
func (p *profile) FetchByIDProfile(ctx context.Context, id int64) (*entity.ProfileEntity, error) {
	modelProfile := model.Profile{}
	err := p.DB.Where("id = ?", id).First(&modelProfile).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDProfile - 1: %v", err)
		return nil, err
	}

	// Mapping model to entity
	return &entity.ProfileEntity{
		ID:        modelProfile.ID,
		Title:     modelProfile.Title,
		Content:   modelProfile.Content,
	}, nil
}

// EditByIDProfile implements ProfileInterface.
func (p *profile) EditByIDProfile(ctx context.Context, req entity.ProfileEntity) error {
	modelProfile := model.Profile{}

	err := p.DB.Where("id = ?", req.ID).First(&modelProfile).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDProfile - 1: %v", err)
		return err
	}

	// Update the profile fields
	modelProfile.Title = req.Title
	modelProfile.Content = req.Content

	// Save the updated profile
	err = p.DB.Save(&modelProfile).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDProfile - 2: %v", err)
		return err
	}
	return nil
}

func NewProfileRepository(DB *gorm.DB) ProfileInterface {
	return &profile{
		DB: DB,
	}
}
