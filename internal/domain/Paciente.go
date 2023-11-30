package domain

// Paciente es una estructura que define ...
type Paciente struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	DNI          string `json:"DNI"`
	Address      string `json:"address"`
	CreationDate string `json:"creation_date"`
}
