package odontologo

import (
	"context"
	"log"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Service interface {
	Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error)
	GetAll(ctx context.Context) ([]domain.Odontologo, error)
	GetByID(ctx context.Context, id int) (domain.Odontologo, error)
	Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
}

type service struct {
	repository Repository
}

func NewServiceOdontologo(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new odontologo.
func (s *service) Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error) {
	odontologo, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("[OdontologosService][Create] error creating dentist", err)
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// GetAll is a method that returns all odontologos.
func (s *service) GetAll(ctx context.Context) ([]domain.Odontologo, error) {
	listProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[OdontologosService][GetAll] error getting all odontologos", err)
		return []domain.Odontologo{}, err
	}

	return listProducts, nil
}

// GetByID is a method that returns an odontologo by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Odontologo, error) {
	odontologo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][GetByID] error getting odontologo by ID", err)
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// Update is a method that updates an odontologo by ID.
func (s *service) Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	odontologo, err := s.repository.Update(ctx, odontologo, id)
	if err != nil {
		log.Println("[OdontologosService][Update] error updating odontologo by ID", err)
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// Delete is a method that deletes an odontologo by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][Delete] error deleting product by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates an odontologo by ID.
func (s *service) Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	odontologoStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][Patch] error getting odontologo by ID", err)
		return domain.Odontologo{}, err
	}

	odontologoPatch, err := s.validatePatch(odontologoStore, odontologo)
	if err != nil {
		log.Println("[ProductsService][Patch] error validating product", err)
		return domain.Odontologo{}, err
	}

	odontologo, err = s.repository.Patch(ctx, odontologoPatch, id)
	if err != nil {
		log.Println("[ProductsService][Patch] error patching product by ID", err)
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(odontologoStore, odontologo domain.Odontologo) (domain.Odontologo, error) {

	if odontologo.Name != "" {
		odontologoStore.Name = odontologo.Name
	}

	if odontologo.LastName != "" {
		odontologoStore.LastName = odontologo.LastName
	}

	if odontologo.MedicalId != "" {
		odontologoStore.MedicalId = odontologo.MedicalId
	}

	return odontologoStore, nil

}
