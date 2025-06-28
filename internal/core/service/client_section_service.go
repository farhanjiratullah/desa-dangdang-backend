package service

import (
	"context"
	"desadangdang/internal/adapater/repository"
	"desadangdang/internal/core/domain/entity"
)

type ClientSectionServiceInterface interface {
	CreateClientSection(ctx context.Context, req entity.ClientSectionEntity) error
	FetchAllClientSection(ctx context.Context) ([]entity.ClientSectionEntity, error)
	FetchByIDClientSection(ctx context.Context, id int64) (*entity.ClientSectionEntity, error)
	EditByIDClientSection(ctx context.Context, req entity.ClientSectionEntity) error
	DeleteByIDClientSection(ctx context.Context, id int64) error
}

type clientSectionService struct {
	clientSectionRepo repository.ClientSectionInterface
}

// CreateClientSection implements ClientSectionServiceInterface.
func (c *clientSectionService) CreateClientSection(ctx context.Context, req entity.ClientSectionEntity) error {
	return c.clientSectionRepo.CreateClientSection(ctx, req)
}

// DeleteByIDClientSection implements ClientSectionServiceInterface.
func (c *clientSectionService) DeleteByIDClientSection(ctx context.Context, id int64) error {
	return c.clientSectionRepo.DeleteByIDClientSection(ctx, id)
}

// EditByIDClientSection implements ClientSectionServiceInterface.
func (c *clientSectionService) EditByIDClientSection(ctx context.Context, req entity.ClientSectionEntity) error {
	return c.clientSectionRepo.EditByIDClientSection(ctx, req)
}

// FetchAllClientSection implements ClientSectionServiceInterface.
func (c *clientSectionService) FetchAllClientSection(ctx context.Context) ([]entity.ClientSectionEntity, error) {
	return c.clientSectionRepo.FetchAllClientSection(ctx)
}

// FetchByIDClientSection implements ClientSectionServiceInterface.
func (c *clientSectionService) FetchByIDClientSection(ctx context.Context, id int64) (*entity.ClientSectionEntity, error) {
	return c.clientSectionRepo.FetchByIDClientSection(ctx, id)
}

func NewClientSectionService(clientSectionRepo repository.ClientSectionInterface) ClientSectionServiceInterface {
	return &clientSectionService{
		clientSectionRepo: clientSectionRepo,
	}
}
