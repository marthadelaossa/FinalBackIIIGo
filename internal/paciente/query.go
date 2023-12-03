package paciente

var (
	QueryInsertPaciente = `INSERT INTO clinicaodontologica.paciente(id,name,last_name,DNI,address,creation_date)
	VALUES(?,?,?)`
	QueryGetAllPacientes = `SELECT id,name,last_name,DNI,address,creation_date
	FROM clinicaodontologica.paciente`
	QueryDeletePaciente    = `DELETE FROM clinicaodontologica.paciente WHERE id = ?`
	QueryGetOdontologoById = `SELECT id,name,last_name,DNI,address,creation_date
	FROM clinicaodontologica.paciente WHERE id = ?`
	QueryUpdatePaciente = `UPDATE clinicaodontologica.paciente SET name = ?, last_name = ?, DNI= ?, address = ?, creation_date= ?
	WHERE id = ?`
	QueryPatchPaciente = `UPDATE clinicaodontologica.paciente SET name = COALESCE(?, name), last_name = COALESCE(?, last_name), DNI = COALESCE(?, DNI), address = COALESCE(?, address), creation_date = COALESCE(?, creation_date)
	WHERE id = ?`
)
