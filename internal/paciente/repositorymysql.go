package paciente

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create is a method that creates a paciente.
func (r *repositorymysql) Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error) {
	statement, err := r.db.Prepare(QueryInsertPaciente)
	if err != nil {
		return domain.Paciente{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Name,
		paciente.LastName,
		paciente.Address,
		paciente.DNI,
	)

	if err != nil {
		return domain.Paciente{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Paciente{}, ErrLastInsertedId
	}

	paciente.Id = int(lastId)

	return paciente, nil

}

// GetAll is a method that returns all pacientes.
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Paciente, error) {
	rows, err := r.db.Query(QueryGetAllPaciente)
	if err != nil {
		return []domain.Paciente{}, err
	}

	defer rows.Close()

	var pacientes []domain.Paciente

	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(
			&paciente.Id,
			&paciente.Name,
			&paciente.LastName,
			&paciente.Address,
			&paciente.DNI,
			&paciente.CreationDate,
		)
		if err != nil {
			return []domain.Paciente{}, err
		}

		pacientes = append(pacientes, paciente)
	}

	if err := rows.Err(); err != nil {
		return []domain.Paciente{}, err
	}

	return pacientes, nil
}

// GetByID is a method that returns a paciente by ID.
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Paciente, error) {
	row := r.db.QueryRow(QueryGetPacienteById, id)

	var paciente domain.Paciente
	err := row.Scan(
		&paciente.Id,
		&paciente.Name,
		&paciente.LastName,
		&paciente.Address,
		&paciente.DNI,
		&paciente.CreationDate,
	)

	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Update is a method that updates a paciente by ID.
func (r *repositorymysql) Update(
	ctx context.Context,
	paciente domain.Paciente,
	id int) (domain.Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdatePaciente)
	if err != nil {
		return domain.Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Name,
		paciente.LastName,
		paciente.Address,
		paciente.DNI,
		id,
	)

	if err != nil {
		return domain.Paciente{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Paciente{}, err
	}

	paciente.Id = id

	return paciente, nil

}

// Delete is a method that deletes a paciente by ID.
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePaciente, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates a paciente by ID.
func (r *repositorymysql) Patch(
	ctx context.Context,
	paciente domain.Paciente,
	id int) (domain.Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdatePaciente)
	if err != nil {
		return domain.Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Name,
		paciente.LastName,
		paciente.Address,
		paciente.DNI,
		id,
	)

	if err != nil {
		return domain.Paciente{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}
