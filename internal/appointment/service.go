package appointment

import (
	"context"
	"log"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Service interface {
	Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error)
	GetAll(ctx context.Context) ([]domain.Appointment, error)
	GetByID(ctx context.Context, id int) (domain.Appointment, error)
	Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
}

type service struct {
	repository Repository
}

func NewAppointmentService(r Repository) Service {
	return &service{repository: r}
}

// Create implements Service.
func (s *service) Create(ctx context.Context, results domain.Appointment) (domain.Appointment, error) {
	result, err := s.repository.Create(ctx, results)
	if err != nil {
		log.Println("[AppointmentService][create]: error creating appointment -> ", err)
	}
	log.Println("\n", result)
	return result, nil
}

// Delete implements Service.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[AppointmentService][delete]: error deleting appointment by id -> ", err)
		return err
	}
	return nil
}

// GetAll implements Service.
func (s *service) GetAll(ctx context.Context) ([]domain.Appointment, error) {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[AppointmentService][getAll]: error getting existing appointments -> ", err)
		return []domain.Appointment{}, err
	}
	return results, nil
}

// GetByID implements Service.
func (s *service) GetByID(ctx context.Context, id int) (domain.Appointment, error) {
	result, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[AppointmentService][getById]: error getting existing appointment by id -> ", err)
		return domain.Appointment{}, err
	}
	return result, nil
}

// Patch implements Service.
func (s *service) Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	result, err := s.repository.Patch(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentService][patch]: error patching existing appointment by id -> ", err)
		return domain.Appointment{}, err
	}
	return result, err
}

// Update implements Service.
func (s *service) Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	result, err := s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentService][update]: error updating existing appointment by id -> ", err)
		return domain.Appointment{}, err
	}
	return result, err
}
