package paciente

import (
	"context"
	"errors"
	"log"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("paciente not found")
)

type repository struct {
	db []domain.Paciente
}

// NewMemoryRepository ....
func NewMemoryRepository(db []domain.Paciente) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new paciente.
func (r *repository) Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error) {
	r.db = append(r.db, paciente)
	return paciente, nil
}

// GetAll is a method that returns all pacientes.
func (r *repository) GetAll(ctx context.Context) ([]domain.Paciente, error) {

	contenidoContext := ctx.Value("rol")

	if contenidoContext != "" {
		log.Println("El contenido del contexto es:", contenidoContext)
	}

	if len(r.db) < 1 {
		return []domain.Paciente{}, ErrEmpty
	}

	return r.db, nil
}

// GetByID is a method that returns a paciente by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Paciente, error) {
	var result domain.Paciente
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id < 1 {
		return domain.Paciente{}, ErrNotFound
	}

	return result, nil
}

// Update is a method that updates a paciente by ID.
func (r *repository) Update(
	ctx context.Context,
	paciente domain.Paciente,
	id int) (domain.Paciente, error) {

	var result domain.Paciente
	for key, value := range r.db {
		if value.Id == id {
			paciente.Id = id
			r.db[key] = paciente
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Paciente{}, ErrNotFound
	}

	return result, nil

}

// Delete is a method that deletes a paciente by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	var result domain.Paciente
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

// Patch is a method that updates a paciente by ID.
func (r *repository) Patch(
	ctx context.Context,
	paciente domain.Paciente,
	id int) (domain.Paciente, error) {

	var result domain.Paciente
	for key, value := range r.db {
		if value.Id == id {
			if paciente.Name != "" {
				r.db[key].Name = paciente.Name
			}
			if paciente.LastName != "" {
				r.db[key].LastName = paciente.LastName
			}
			if paciente.Address != "" {
				r.db[key].Address = paciente.Address
			}
			if paciente.DNI != "" {
				r.db[key].DNI = paciente.DNI
			}
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Paciente{}, ErrNotFound
	}

	return result, nil
}
