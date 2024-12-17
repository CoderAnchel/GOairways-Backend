package models

import "rutas/database"

type Ruta struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	Origen_id         int    `json:"origen"`
	Destino_id        int    `json:"destino"`
	Duracion_minutos  int    `json:"duracion"`
	Distancia         int    `json:"distancia"`
	Frequencia_dia    int    `json:"frequencia_dia"`
	Frecuencia_semana int    `json:"frequencia_semana"`
	Activa            int    `json:"activa"`
	Monday            int    `json:"monday"`
	Tuesday           int    `json:"tuesday"`
	Wednesday         int    `json:"wednesday"`
	Thursday          int    `json:"thursday"`
	Friday            int    `json:"friday"`
	Saturday          int    `json:"saturday"`
	Sunday            int    `json:"sunday"`
	Type              string `json:"type"`
}

type Horario struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Ruta_id      int    `json:"ruta_id"`
	Hora_salida  string `json:"hora_salida"`
	Hora_llegada string `json:"hora_llegada"`
}

type ModeloRuta struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Ruta_id   int  `json:"ruta_id"`
	Modelo_id int  `json:"modelo"`
}

func (Ruta) TableName() string {
	return "rutas"
}

func (Horario) TableName() string {
	return "horarios_rutas"
}

func (ModeloRuta) TableName() string {
	return "modelos_rutas"
}

func GetAllRoutes() []Ruta {
	var rutas []Ruta
	database.DB.Find(&rutas)
	return rutas
}

func GetInactiveRoutes() []Ruta {
	var rutas []Ruta
	database.DB.Where("activa = 0").Find(&rutas)
	return rutas
}

func GetActiveRoutes() []Ruta {
	var rutas []Ruta
	database.DB.Where("activa = 1").Find(&rutas)
	return rutas
}

func CreateRoute(route *Ruta) error {
	return database.DB.Create(route).Error
}

func CreateSchedule(schedule *Horario) error {
	return database.DB.Create(schedule).Error
}

func CreateModelRoute(modelRoute *ModeloRuta) error {
	return database.DB.Create(modelRoute).Error
}

func ActivateRouteByID(id uint) error {
	return database.DB.Model(&Ruta{}).Where("id = ?", id).Update("activa", 1).Error
}

func GetNextRouteID() (uint, error) {
	var maxID uint
	result := database.DB.Table("rutas").Select("MAX(id)").Scan(&maxID)
	if result.Error != nil {
		return 0, result.Error
	}
	return maxID + 1, nil
}
