package repository

import (
	"context"
	"desadangdang/internal/core/domain/entity"
	"desadangdang/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AppointmentRepositoryInterface interface {
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	DeleteByIDAppointment(ctx context.Context, id int64) error
	CreateAppointment(ctx context.Context, req entity.AppointmentEntity) (string, error)
}

type appointmentRepository struct {
	DB *gorm.DB
}

// CreateAppointment implements AppointmentRepositoryInterface.
func (h *appointmentRepository) CreateAppointment(ctx context.Context, req entity.AppointmentEntity) (string, error) {
	modelAppointment := model.Appointment{
		ServiceID:   req.ServiceID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Brief:       req.Brief,
		Budget:      req.Budget,
		MeetAt:      req.MeetAt,
	}

	if err = h.DB.Create(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAppointment - 1: %v", err)
		return "", err
	}

	return modelAppointment.Email, nil

}

// DeleteByIDAppointment implements AppointmentInterface.
func (h *appointmentRepository) DeleteByIDAppointment(ctx context.Context, id int64) error {
	modelAppointment := model.Appointment{}

	if err = h.DB.Where("id = ?", id).First(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAppointment implements AppointmentInterface.
func (h *appointmentRepository) FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error) {
	rows, err := h.DB.
		Table("appointments as a").
		Select("a.id", "a.name", "a.email", "a.phone_number", "a.brief", "a.budget", "ss.name").
		Joins("inner join service_sections as ss on ss.id = a.service_id").
		Where("a.deleted_at IS NULL").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAppointment - 1: %v", err)
		return nil, err
	}

	var appointmentRepositoryEntities []entity.AppointmentEntity
	for rows.Next() {
		var appointment entity.AppointmentEntity
		err = rows.Scan(&appointment.ID, &appointment.Name, &appointment.Email, &appointment.PhoneNumber, &appointment.Brief, &appointment.Budget, &appointment.ServiceName)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAppointment - 2: %v", err)
			return nil, err
		}
		appointmentRepositoryEntities = append(appointmentRepositoryEntities, appointment)
	}

	return appointmentRepositoryEntities, nil
}

// FetchByIDAppointment implements AppointmentInterface.
func (h *appointmentRepository) FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error) {
	rows, err := h.DB.
		Table("appointments as a").
		Select("a.id", "a.phone_number", "brief", "meet_at", "a.name", "a.email", "a.budget", "ss.id", "ss.name").
		Joins("inner join service_sections as ss on ss.id = a.service_id").
		Where("a.id =? AND a.deleted_at IS NULL", id).
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDAppointment - 1: %v", err)
		return nil, err
	}

	appointment := &entity.AppointmentEntity{}
	for rows.Next() {
		err = rows.Scan(&appointment.ID, &appointment.PhoneNumber, &appointment.Brief, &appointment.MeetAt, &appointment.Name, &appointment.Email, &appointment.Budget, &appointment.ServiceID, &appointment.ServiceName)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchByIDAppointment - 2: %v", err)
			return nil, err
		}
	}

	return appointment, nil
}

func NewAppointmentRepository(DB *gorm.DB) AppointmentRepositoryInterface {
	return &appointmentRepository{
		DB: DB,
	}
}
