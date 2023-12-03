package odontologo

import (
	"context"
	"errors"
	"log"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

type repository struct {
	db []domain.Odontologo
}

// NewMemoryRepository ....
func NewMemoryRepository(db []domain.Odontologo) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new odontologo.
func (r *repository) Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error) {
	r.db = append(r.db, odontologo)
	return odontologo, nil
}

// GetAll is a method that returns all odontologos.
func (r *repository) GetAll(ctx context.Context) ([]domain.Odontologo, error) {

	contenidoContext := ctx.Value("rol")

	if contenidoContext != "" {
		log.Println("El contenido del contexto es:", contenidoContext)
	}

	if len(r.db) < 1 {
		return []domain.Odontologo{}, ErrEmpty
	}

	return r.db, nil
}

// GetByID is a method that returns an odontologo by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Odontologo, error) {
	var result domain.Odontologo
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id < 1 {
		return domain.Odontologo{}, ErrNotFound
	}

	return result, nil
}

// Update is a method that updates an odontologo by ID.
func (r *repository) Update(
	ctx context.Context,
	odontologo domain.Odontologo,
	id int) (domain.Odontologo, error) {

	var result domain.Odontologo
	for key, value := range r.db {
		if value.Id == id {
			odontologo.Id = id
			r.db[key] = odontologo
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Odontologo{}, ErrNotFound
	}

	return result, nil

}

// Delete is a method that deletes an odontologo by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	var result domain.Odontologo
	for key, value := range r.db {
		if value.Id == id {
			result = r.db[key]
			r.db = append(r.db[:key], r.db[key+1:]...)
			break
		}
	}

	if result.Id < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates an odontologo by ID.
func (r *repository) Patch(
	ctx context.Context,
	odontologo domain.Odontologo,
	id int) (domain.Odontologo, error) {

	var result domain.Odontologo
	for key, value := range r.db {
		if value.Id == id {
			if odontologo.Name != "" {
				r.db[key].Name = odontologo.Name
			}
			if odontologo.LastName != "" {
				r.db[key].LastName = odontologo.LastName
			}
			if odontologo.MedicalId != "" {
				r.db[key].MedicalId = odontologo.MedicalId
			}
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Odontologo{}, ErrNotFound
	}

	return result, nil
}
