package service

import (
	"context"
	"desadangdang/internal/adapater/repository"
	"desadangdang/internal/core/domain/entity"
)

type ServiceSectionServiceInterface interface {
	CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error)
	FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error)
	EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	DeleteByIDServiceSection(ctx context.Context, id int64) error
}

type serviceSectionService struct {
	serviceSectionRepo repository.ServiceSectionRepositoryInterface
}

// CreateServiceSection implements ServiceSectionServiceInterface.
func (c *serviceSectionService) CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	return c.serviceSectionRepo.CreateServiceSection(ctx, req)
}

// DeleteByIDServiceSection implements ServiceSectionServiceInterface.
func (c *serviceSectionService) DeleteByIDServiceSection(ctx context.Context, id int64) error {
	return c.serviceSectionRepo.DeleteByIDServiceSection(ctx, id)
}

// EditByIDServiceSection implements ServiceSectionServiceInterface.
func (c *serviceSectionService) EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	return c.serviceSectionRepo.EditByIDServiceSection(ctx, req)
}

// FetchAllServiceSection implements ServiceSectionServiceInterface.
func (c *serviceSectionService) FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error) {
	return c.serviceSectionRepo.FetchAllServiceSection(ctx)
}

// FetchByIDServiceSection implements ServiceSectionServiceInterface.
func (c *serviceSectionService) FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error) {
	return c.serviceSectionRepo.FetchByIDServiceSection(ctx, id)
}

func NewServiceSectionService(serviceSectionRepo repository.ServiceSectionRepositoryInterface) ServiceSectionServiceInterface {
	return &serviceSectionService{
		serviceSectionRepo: serviceSectionRepo,
	}
}
