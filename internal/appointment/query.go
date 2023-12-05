package appointment

const (
	QueryInsertAppointment  = `INSERT INTO clinicaodontologica.appointments(description, id_odontologo, id_paciente, date_time) VALUES (?, ?, ?, ?)`
	QueryGetAllAppointments = `SELECT id, description, id_odontologo, id_paciente, date_time FROM clinicaodontologica.appointments`
	QueryDeleteAppointment  = `DELETE FROM clinicaodontologica.appointments WHERE id = ?`
	QueryGetAppointmentById = `SELECT id, description, id_odontologo, id_paciente, date_time FROM clinicaodontologica.appointments WHERE id = ?`
	QueryUpdateAppointment  = `UPDATE clinicaodontologica.appointments SET description = ?, id_odontologo = ?, id_paciente = ?, date_time = ? WHERE id = ?`
	QueryPatchAppointment   = `UPDATE clinicaodontologica.appointments SET description = COALESCE(?, description), id_odontologo = COALESCE(?, id_odontologo), id_paciente = COALESCE(?, id_paciente), date_time = COALESCE(?, date_time) WHERE id = ?`
)
