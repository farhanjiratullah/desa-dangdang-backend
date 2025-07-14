package service

import (
	"context"
	"desadangdang/internal/adapater/repository"
	"desadangdang/internal/core/domain/entity"
)

type ProfileServiceInterface interface {
	FetchByIDProfile(ctx context.Context, id int64) (*entity.ProfileEntity, error)
	EditByIDProfile(ctx context.Context, req entity.ProfileEntity) error
}

type profileService struct {
	profileRepo repository.ProfileInterface
}

// FetchByIDProfile implements ProfileServiceInterface.
func (p *profileService) FetchByIDProfile(ctx context.Context, id int64) (*entity.ProfileEntity, error) {
	return p.profileRepo.FetchByIDProfile(ctx, id)
}

// EditByIDProfile implements ProfileServiceInterface.
func (p *profileService) EditByIDProfile(ctx context.Context, req entity.ProfileEntity) error {
	return p.profileRepo.EditByIDProfile(ctx, req)
}

func NewProfileService(profileRepo repository.ProfileInterface) ProfileServiceInterface {
	return &profileService{
		profileRepo: profileRepo,
	}
}
