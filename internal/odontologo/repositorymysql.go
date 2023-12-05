package odontologo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrExecStatement    = errors.New("error executing statement")
	ErrLastInsertedId   = errors.New("error getting last inserted ID")
)

type OdontologoMySqlRepository struct {
	db *sql.DB
}

// NewMySqlRepository creates a new instance of MySqlRepository.
func NewMySqlRepository(db *sql.DB) Repository {
	return &OdontologoMySqlRepository{db: db}
}

// Create inserts a new odontologo into the database.
func (r *OdontologoMySqlRepository) Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryInsertOdontologo)
	if err != nil {
		return domain.Odontologo{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(odontologo.Name, odontologo.LastName, odontologo.MedicalId)
	if err != nil {
		return domain.Odontologo{}, ErrExecStatement
	}

	lastID, err := result.LastInsertId()
	
	if err != nil {
		return domain.Odontologo{}, ErrLastInsertedId
	}

	odontologo.Id = int(lastID)

	return odontologo, nil
}

// GetAll retrieves all odontologos from the database.
func (r *OdontologoMySqlRepository) GetAll(ctx context.Context) ([]domain.Odontologo, error) {
	rows, err := r.db.Query(QueryGetAllOdontologos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var odontologos []domain.Odontologo

	for rows.Next() {
		var odontologo domain.Odontologo
		err := rows.Scan(&odontologo.Id, &odontologo.Name, &odontologo.LastName, &odontologo.MedicalId)
		if err != nil {
			return nil, err
		}

		odontologos = append(odontologos, odontologo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return odontologos, nil
}

// GetByID retrieves an odontologo by ID from the database.
func (r *OdontologoMySqlRepository) GetByID(ctx context.Context, id int) (domain.Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo domain.Odontologo
	err := row.Scan(&odontologo.Id, &odontologo.Name, &odontologo.LastName, &odontologo.MedicalId)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// Update updates an odontologo by ID in the database.
func (r *OdontologoMySqlRepository) Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(odontologo.Name, odontologo.LastName, odontologo.MedicalId, id)
	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	odontologo.Id = id

	return odontologo, nil
}

// Delete deletes an odontologo by ID from the database.
func (r *OdontologoMySqlRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
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

// Patch updates an odontologo by ID in the database.
func (r *OdontologoMySqlRepository) Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(odontologo.Name, odontologo.LastName, odontologo.MedicalId, id)
	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}
