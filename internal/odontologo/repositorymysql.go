package odontologo

import (
	"context"
	"database/sql"

	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
)

type repositoryMySql struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) Repository {
	return &repositoryMySql{db: db}
}

// Create inserta un nuevo odontólogo.
func (r *repositoryMySql) Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error) {
	result, err := r.db.Exec(QueryInsertOdontologo, odontologo.Name, odontologo.LastName, odontologo.MedicalId)
	if err != nil {
		return domain.Odontologo{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Odontologo{}, err
	}

	odontologo.Id = int(lastID)
	return odontologo, nil
}

func (r *repositoryMySql) GetAll(ctx context.Context) ([]domain.Odontologo, error) {
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

func (r *repositoryMySql) GetByID(ctx context.Context, id int) (domain.Odontologo, error) {
	var odontologo domain.Odontologo
	err := r.db.QueryRow(QueryGetOdontologoById, id).Scan(&odontologo.Id, &odontologo.Name, &odontologo.LastName, &odontologo.MedicalId)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (r *repositoryMySql) Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	_, err := r.db.Exec(QueryUpdateOdontologo, odontologo.Name, odontologo.LastName, odontologo.MedicalId, id)
	if err != nil {
		return domain.Odontologo{}, err
	}

	odontologo.Id = id
	return odontologo, nil
}

func (r *repositoryMySql) Delete(ctx context.Context, id int) error {
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

func (r *repositoryMySql) Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error) {
	_, err := r.db.Exec(QueryPatchOdontologo, odontologo.Name, odontologo.LastName, odontologo.MedicalId, id)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}
