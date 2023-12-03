package odontologo

var (
	QueryInsertOdontologo = `INSERT INTO clinicaodontologica.odontologo(name,last_name,medical_ID)
	VALUES(?,?,?)`
	QueryGetAllOdontologos = `SELECT id,name,last_name,medical_ID 
	FROM clinicaodontologica.odontologo`
	QueryDeleteOdontologo  = `DELETE FROM clinicaodontologica.odontologo WHERE id = ?`
	QueryGetOdontologoById = `SELECT id,name,last_name,medical_ID
	FROM clinicaodontologica.odontologo WHERE id = ?`
	QueryUpdateOdontologo = `UPDATE clinicaodontologica.odontologo SET name = ?, last_name = ?, medical_ID = ?
	WHERE id = ?`
)
