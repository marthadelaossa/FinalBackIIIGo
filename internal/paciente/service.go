package paciente

import (
	"context"
	"log"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Service interface {
	Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error)
	GetAll(ctx context.Context) ([]domain.Paciente, error)
	GetByID(ctx context.Context, id int) (domain.Paciente, error)
	Update(ctx context.Context, producto domain.Paciente, id int) (domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, producto domain.Paciente, id int) (domain.Paciente, error)
}

type service struct {
	repository Repository
}

func NewServiceOdontologo(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new paciente.
func (s *service) Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.repository.Create(ctx, paciente)
	if err != nil {
		log.Println("[PacientesService][Create] error creating paciente", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// GetAll is a method that returns all products.
func (s *service) GetAll(ctx context.Context) ([]domain.Paciente, error) {
	listProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[PacientesService][GetAll] error getting all odontologos", err)
		return []domain.Paciente{}, err
	}

	return listProducts, nil
}

// GetByID is a method that returns a paciente by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Paciente, error) {
	paciente, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PacientesService][GetByID] error getting paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Update is a method that updates a paciente by ID.
func (s *service) Update(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error) {
	paciente, err := s.repository.Update(ctx, paciente, id)
	if err != nil {
		log.Println("[PacientesService][Update] error updating paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Delete is a method that deletes a paciente by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PacientesService][Delete] error deleting paciente by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a paciente by ID.
func (s *service) Patch(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error) {
	pacienteStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PacientesService][Patch] error getting paciente by ID", err)
		return domain.Paciente{}, err
	}

	pacientePatch, err := s.validatePatch(pacienteStore, paciente)
	if err != nil {
		log.Println("[ProductsService][Patch] error validating paciente", err)
		return domain.Paciente{}, err
	}

	paciente, err = s.repository.Patch(ctx, pacientePatch, id)
	if err != nil {
		log.Println("[ProductsService][Patch] error patching paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(pacienteStore, paciente domain.Paciente) (domain.Paciente, error) {

	if paciente.Name != "" {
		pacienteStore.Name = paciente.Name
	}

	if paciente.LastName != "" {
		pacienteStore.LastName = paciente.LastName
	}

	if paciente.DNI != "" {
		pacienteStore.DNI = paciente.DNI
	}

	if paciente.Address != "" {
		pacienteStore.Address = paciente.Address
	}

	if paciente.CreationDate != "" {
		pacienteStore.CreationDate = paciente.CreationDate
	}
	return pacienteStore, nil

}
