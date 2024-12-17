package models

import "aereopuertos/database"

type Aereopuerto struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Nombre       string `json:"nombre"`
	Codigo_iata  string `json:"codigo_iata"`
	Ubicacion    string `json:"ubicacion"`
	Latitud      string `json:"latitud"`
	Longitud     string `json:"longitud"`
	Zona_horaria string `json:"zona_horaria"`
	Url          string `json:"url"`
}

func GetAllAirports() []Aereopuerto {
	var aeropuertos []Aereopuerto
	database.DB.Find(&aeropuertos)
	return aeropuertos
}
