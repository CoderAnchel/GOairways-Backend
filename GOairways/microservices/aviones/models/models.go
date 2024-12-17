package models

import "aviones/database"

type Modelo struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Modelo        string  `json:"modelo"`
	Capacidad     int     `json:"capacidad"`
	Filas         int     `json:"filas"`
	Configuracion string  `json:"configuracion"`
	Tipo          string  `json:"tipo"`
	Longitud      float64 `json:"longitud"`
	Enevergadura  float64 `json:"enevergadura"`
	Alcanze       int     `json:"alcanze"`
	Motores       int     `json:"motores"`
}

type Avion struct {
	ID                         uint   `json:"id" gorm:"primaryKey"`
	Modelo_id                  uint   `json:"modelo_id"`
	Nombre                     string `json:"nombre"`
	Matricula                  string `json:"matricula"`
	Ubicacion                  string `json:"ubicacion"`
	Status                     string `json:"status"`
	Km_recorridos              int    `json:"km_recorridos"`
	Fecha_entrada_servicio     string `json:"fecha_entrada_servicio"`
	Fecha_ultimo_mantenimiento string `json:"fecha_ultimo_mantenimiento"`
	Proximo_mantenimiento      string `json:"proximo_mantenimiento"`
}

func (Avion) TableName() string {
	return "aviones"
}

func (Modelo) TableName() string {
	return "modelos"
}

func CreateNewModelo(avion *Modelo) error {
	return database.DB.Create(avion).Error
}

func GetPlaneByID(id int) (Avion, error) {
	var avion Avion
	err := database.DB.Where("id = ?", id).First(&avion).Error
	return avion, err
}

func GetModeloByID(id int) (Modelo, error) {
	var avion Modelo
	err := database.DB.Where("id = ?", id).First(&avion).Error
	return avion, err
}

func GetAllModelos() []Modelo {
	var aviones []Modelo
	database.DB.Find(&aviones)
	return aviones
}

func GetAllPlanes() []Avion {
	var aviones []Avion
	database.DB.Find(&aviones)
	return aviones
}

func GetModelosByTipo(tipo string) ([]Modelo, error) {
	var modelos []Modelo
	err := database.DB.Where("tipo = ?", tipo).Find(&modelos).Error
	return modelos, err
}

func CheckPlateExists(plate string) bool {
	var avion Avion
	err := database.DB.Where("matricula = ?", plate).First(&avion).Error
	return err == nil
}
