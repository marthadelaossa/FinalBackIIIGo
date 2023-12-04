package paciente

var (
	QueryInsertPaciente = `INSERT INTO clinicaodontologica.paciente(name,last_name,address,DNI)
	VALUES(?,?,?,?)`
	QueryGetAllPaciente = `SELECT id,name,last_name,address,DNI,creation_date
	FROM clinicaodontologica.paciente`
	QueryDeletePaciente  = `DELETE FROM clinicaodontologica.paciente WHERE id = ?`
	QueryGetPacienteById = `SELECT id,name,last_name,address,DNI,creation_date
	FROM clinicaodontologica.paciente WHERE id = ?`
	QueryUpdatePaciente = `UPDATE clinicaodontologica.paciente SET name = ?,last_name = ?,address = ?,DNI = ?
	WHERE id = ?`
	QueryPatchPaciente = `UPDATE clinicaodontologica.paciente SET name = COALESCE(?, name),last_name = COALESCE(?, last_name),address = COALESCE(?, address),DNI = COALESCE(?, DNI)
	WHERE id = ?`
)
