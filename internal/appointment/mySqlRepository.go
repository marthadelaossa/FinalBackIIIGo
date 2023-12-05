package appointment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

var (
	notAppointmentFoundForDeletion = errors.New("No one medical appointment were found for deleting.")
)

type AppointmentMySqlRepository struct {
	db *sql.DB
}

// Create implements Repository.
func (ar *AppointmentMySqlRepository) Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error) {
	sqlStatement, err := ar.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	defer sqlStatement.Close()

	result, err := sqlStatement.Exec(
		appointment.Description,
		appointment.OdontologoId,
		appointment.PacienteId,
		appointment.DateTime,
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	lastId, err := result.LastInsertId()

	if err == nil {
		return domain.Appointment{}, err
	}

	appointment.Id = int(lastId)

	return appointment, nil
}

// Delete implements Repository.
func (ar *AppointmentMySqlRepository) Delete(ctx context.Context, id int) error {
	result, err := ar.db.Exec(QueryDeleteAppointment, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return notAppointmentFoundForDeletion
	}
	return nil
}

// GetAll implements Repository.
func (ar *AppointmentMySqlRepository) GetAll(ctx context.Context) ([]domain.Appointment, error) {
	rows, err := ar.db.Query(QueryGetAllAppointments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicalAppointments []domain.Appointment

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var medicalAppointment domain.Appointment
		err := rows.Scan(
			&medicalAppointment.Id,
			&medicalAppointment.Description,
			&medicalAppointment.OdontologoId,
			&medicalAppointment.PacienteId,
			&medicalAppointment.DateTime,
		)
		if err != nil {
			return nil, err
		}
		medicalAppointments = append(medicalAppointments, medicalAppointment)
	}

	return medicalAppointments, nil
}

// GetByID implements Repository.
func (ar *AppointmentMySqlRepository) GetByID(ctx context.Context, id int) (domain.Appointment, error) {
	row := ar.db.QueryRow(QueryGetAppointmentById, id)

	var medicalAppointment domain.Appointment
	err := row.Scan(
		&medicalAppointment.Id,
		&medicalAppointment.Description,
		&medicalAppointment.OdontologoId,
		&medicalAppointment.PacienteId,
		&medicalAppointment.DateTime,
	)
	if err != nil {
		return domain.Appointment{}, err
	}
	return medicalAppointment, nil
}

// Patch implements Repository.
func (ar *AppointmentMySqlRepository) Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	sqlStatement, err := ar.db.Prepare(QueryUpdateAppointment)

	if err != nil {
		return domain.Appointment{}, nil
	}

	defer sqlStatement.Close()

	result, err := sqlStatement.Exec(
		appointment.Id,
		appointment.Description,
		appointment.OdontologoId,
		appointment.PacienteId,
		appointment.DateTime,
	)

	if err != nil {
		return domain.Appointment{}, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return domain.Appointment{}, nil
	}

	return appointment, nil
}

// Update implements Repository.
func (ar *AppointmentMySqlRepository) Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	sqlStatement, err := ar.db.Prepare(QueryUpdateAppointment)

	if err != nil {
		return domain.Appointment{}, nil
	}

	defer sqlStatement.Close()

	result, err := sqlStatement.Exec(
		appointment.Id,
		appointment.OdontologoId,
		appointment.PacienteId,
		appointment.DateTime,
	)

	if err != nil {
		return domain.Appointment{}, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return domain.Appointment{}, nil
	}

	return appointment, nil
}

func NewAppointmentMySqlRepository(db *sql.DB) Repository {
	return &AppointmentMySqlRepository{db: db}
}
