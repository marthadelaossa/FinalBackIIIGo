package domain

// Paciente es una estructura que define ...
type Paciente struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	DNI          string `json:"DNI"`
	CreationDate string `json:"creation_date"`
}
