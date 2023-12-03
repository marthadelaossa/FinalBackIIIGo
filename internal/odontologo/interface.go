package odontologo

import (
	"context"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error)
	GetAll(ctx context.Context) ([]domain.Odontologo, error)
	GetByID(ctx context.Context, id int) (domain.Odontologo, error)
	Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
}
