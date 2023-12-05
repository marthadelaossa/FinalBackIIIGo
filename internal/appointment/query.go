package appointment

const (
	QueryInsertAppointment  = `INSERT INTO clinicaodontologica.turnos(description, id_odontologo, id_paciente, date_time) VALUES (?, ?, ?, ?)`
	QueryGetAllAppointments = `SELECT id, description, id_odontologo, id_paciente, date_time FROM clinicaodontologica.turnos`
	QueryDeleteAppointment  = `DELETE FROM clinicaodontologica.turnos WHERE id = ?`
	QueryGetAppointmentById = `SELECT id, description, id_odontologo, id_paciente, date_time FROM clinicaodontologica.turnos WHERE id = ?`
	QueryUpdateAppointment  = `UPDATE clinicaodontologica.turnos SET description = ?, id_odontologo = ?, id_paciente = ?, date_time = ? WHERE id = ?`
	QueryPatchAppointment   = `UPDATE clinicaodontologica.turnos SET description = COALESCE(?, description), id_odontologo = COALESCE(?, id_odontologo), id_paciente = COALESCE(?, id_paciente), date_time = COALESCE(?, date_time) WHERE id = ?`
)
