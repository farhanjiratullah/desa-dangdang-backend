package repository

import (
	"context"
	"desadangdang/internal/core/domain/entity"
	"desadangdang/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type StatisticInterface interface {
	CreateStatistic(ctx context.Context, req entity.StatisticEntity) error
	FetchAllStatistic(ctx context.Context) ([]entity.StatisticEntity, error)
	FetchByIDStatistic(ctx context.Context, id int64) (*entity.StatisticEntity, error)
	EditByIDStatistic(ctx context.Context, req entity.StatisticEntity) error
	DeleteByIDStatistic(ctx context.Context, id int64) error
}

type statistic struct {
	DB *gorm.DB
}

// CreateStatistic implements StatisticInterface.
func (s *statistic) CreateStatistic(ctx context.Context, req entity.StatisticEntity) error {
	modelStatistic := model.Statistic{
		Name:  req.Name,
		Total: req.Total,
		Icon:  req.Icon,
	}

	if err := s.DB.Create(&modelStatistic).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateStatistic - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDStatistic implements StatisticInterface.
func (s *statistic) DeleteByIDStatistic(ctx context.Context, id int64) error {
	modelStatistic := model.Statistic{}

	err := s.DB.Where("id = ?", id).First(&modelStatistic).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDStatistic - 1: %v", err)
		return err
	}

	err = s.DB.Delete(&modelStatistic).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDStatistic - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDStatistic implements StatisticInterface.
func (s *statistic) EditByIDStatistic(ctx context.Context, req entity.StatisticEntity) error {
	modelStatistic := model.Statistic{}

	err := s.DB.Where("id = ?", req.ID).First(&modelStatistic).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDStatistic - 1: %v", err)
		return err
	}
	modelStatistic.Name = req.Name
	modelStatistic.Total = req.Total
	modelStatistic.Icon = req.Icon
	err = s.DB.Save(&modelStatistic).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDStatistic - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllStatistic implements StatisticInterface.
func (s *statistic) FetchAllStatistic(ctx context.Context) ([]entity.StatisticEntity, error) {
	modelStatistic := []model.Statistic{}
	err := s.DB.Select("id", "name", "total", "icon").Find(&modelStatistic).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllStatistic - 1: %v", err)
		return nil, err
	}

	var statisticEntities []entity.StatisticEntity
	for _, v := range modelStatistic {
		statisticEntities = append(statisticEntities, entity.StatisticEntity{
			ID:    v.ID,
			Name:  v.Name,
			Total: v.Total,
			Icon:  v.Icon,
		})
	}

	return statisticEntities, nil
}

// FetchByIDStatistic implements StatisticInterface.
func (s *statistic) FetchByIDStatistic(ctx context.Context, id int64) (*entity.StatisticEntity, error) {
	modelStatistic := model.Statistic{}
	err := s.DB.Where("id = ?", id).First(&modelStatistic).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDStatistic - 1: %v", err)
		return nil, err
	}

	return &entity.StatisticEntity{
		ID:    modelStatistic.ID,
		Name:  modelStatistic.Name,
		Total: modelStatistic.Total,
		Icon:  modelStatistic.Icon,
	}, nil
}

func NewStatisticRepository(DB *gorm.DB) StatisticInterface {
	return &statistic{
		DB: DB,
	}
}
