package domain

type Appointment struct {
	Id           int    `json:"id"`
	Description  string `json:"description"`
	OdontologoId int    `json:"id_odontologo"`
	PacienteId   int    `json:"id_paciente"`
	DateTime     string `json:"date_time"`
}
