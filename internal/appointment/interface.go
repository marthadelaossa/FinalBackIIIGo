package appointment

import (
	"context"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error)
	GetAll(ctx context.Context) ([]domain.Appointment, error)
	GetByID(ctx context.Context, id int) (domain.Appointment, error)
	Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
}
