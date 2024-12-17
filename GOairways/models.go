package main

type Airport struct {
	CodigoIATA  string `json:"codigo_iata"`
	ID          int    `json:"id"`
	Latitud     string `json:"latitud"`
	Longitud    string `json:"longitud"`
	Nombre      string `json:"nombre"`
	Ubicacion   string `json:"ubicacion"`
	URL         string `json:"url"`
	ZonaHoraria string `json:"zona_horaria"`
}

type Model struct {
	Alcanze       int    `json:"alcanze"`
	Capacidad     int    `json:"capacidad"`
	Configuracion string `json:"configuracion"`
	Envergadura   int    `json:"enevergadura"`
	Filas         int    `json:"filas"`
	ID            int    `json:"id"`
	Longitud      int    `json:"longitud"`
	Modelo        string `json:"modelo"`
	Motores       int    `json:"motores"`
	Tipo          string `json:"tipo"`
}

type FlightTime struct {
	Horas   int `json:"horas"`
	Minutos int `json:"minutos"`
}

type TakeOffTicket struct {
	ID             int        `json:"id"`
	AirportFrom    Airport    `json:"airportFrom"`
	AirportTo      Airport    `json:"airportTo"`
	TakeOffTime    string     `json:"takeOffTime"`
	LandingTime    string     `json:"landingTime"`
	FlightTime     FlightTime `json:"flightTime"`
	SelectedModels []Model    `json:"selectedModels"`
	SelectedDays   []string   `json:"SelectedDays"`
}

type RouteCreator struct {
	ID             string          `json:"$id"`
	AirportFrom    Airport         `json:"airportFrom"`
	AirportTo      Airport         `json:"airportTo"`
	ShowDistance   bool            `json:"showDistance"`
	Tipo           string          `json:"tipo"`
	SelectedModels []Model         `json:"selectedModels"`
	DailyFrequ     int             `json:"dailyFrequ"`
	DaysAWeek      []string        `json:"DaysAWeek"`
	SelectedDays   []string        `json:"SelectedDays"`
	TakeOffTime    interface{}     `json:"takeOffTime"`
	LandingTime    interface{}     `json:"landingTime"`
	FlightTime     FlightTime      `json:"flightTime"`
	TakeOffTickets []TakeOffTicket `json:"takeOffTickets"`
	IsOptionsAPI   bool            `json:"_isOptionsAPI"`
	Distance       int             `json:"distance"`
}

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

type Login struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
	PIN      int    `json:"pin"`
}
