package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	MedicalId string `json:"medical_ID"`
}
