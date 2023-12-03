package paciente

import (
	"context"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error)
	GetAll(ctx context.Context) ([]domain.Paciente, error)
	GetByID(ctx context.Context, id int) (domain.Paciente, error)
	Update(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
}
