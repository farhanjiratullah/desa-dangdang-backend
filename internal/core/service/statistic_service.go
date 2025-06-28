package service

import (
	"context"
	"desadangdang/internal/adapater/repository"
	"desadangdang/internal/core/domain/entity"
)

type StatisticServiceInterface interface {
	CreateStatistic(ctx context.Context, req entity.StatisticEntity) error
	FetchAllStatistic(ctx context.Context) ([]entity.StatisticEntity, error)
	FetchByIDStatistic(ctx context.Context, id int64) (*entity.StatisticEntity, error)
	EditByIDStatistic(ctx context.Context, req entity.StatisticEntity) error
	DeleteByIDStatistic(ctx context.Context, id int64) error
}

type statisticService struct {
	statisticRepo repository.StatisticInterface
}

// CreateStatistic implements StatisticServiceInterface.
func (s *statisticService) CreateStatistic(ctx context.Context, req entity.StatisticEntity) error {
	return s.statisticRepo.CreateStatistic(ctx, req)
}

// DeleteByIDStatistic implements StatisticServiceInterface.
func (s *statisticService) DeleteByIDStatistic(ctx context.Context, id int64) error {
	return s.statisticRepo.DeleteByIDStatistic(ctx, id)
}

// EditByIDStatistic implements StatisticServiceInterface.
func (s *statisticService) EditByIDStatistic(ctx context.Context, req entity.StatisticEntity) error {
	return s.statisticRepo.EditByIDStatistic(ctx, req)
}

// FetchAllStatistic implements StatisticServiceInterface.
func (s *statisticService) FetchAllStatistic(ctx context.Context) ([]entity.StatisticEntity, error) {
	return s.statisticRepo.FetchAllStatistic(ctx)
}

// FetchByIDStatistic implements StatisticServiceInterface.
func (s *statisticService) FetchByIDStatistic(ctx context.Context, id int64) (*entity.StatisticEntity, error) {
	return s.statisticRepo.FetchByIDStatistic(ctx, id)
}

func NewStatisticService(statisticRepo repository.StatisticInterface) StatisticServiceInterface {
	return &statisticService{
		statisticRepo: statisticRepo,
	}
}
