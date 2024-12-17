package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("HOLA DESDE LA API!! ✈️")
}

func getModelos(c *fiber.Ctx) error {
	// Hacer la solicitud al microservicio
	res, err := http.Get("http://127.0.0.1:3002/modelos")
	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}
	defer res.Body.Close()

	// Verificar si el código de estado de la respuesta es OK
	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	// Intentar decodificar el JSON
	var aviones interface{} // O el tipo específico de tu JSON (map, slice, etc.)
	if err := json.Unmarshal(body, &aviones); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	// Devolver la respuesta al cliente como JSON
	return c.JSON(aviones)
}

func getModeloByTipo(c *fiber.Ctx) error {
	modelo := c.Params("tipo")
	var response interface{}

	res, err := http.Get("http://127.0.0.1:3002/modelos/find/" + modelo)
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(response)
}

func getModeloById(c *fiber.Ctx) error {
	plane := c.Params("id")
	var response interface{}

	// Hacer la solicitud al microservicio
	res, err := http.Get("http://127.0.0.1:3002/modelos/" + plane)
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	// Verificar si el código de estado de la respuesta es OK
	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	// Intentar decodificar el JSON
	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(response)
}

func getAereopuertos(c *fiber.Ctx) error {
	// Hacer la solicitud al microservicio
	res, err := http.Get("http://127.0.0.1:3003/aereopuertos")

	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	var aereopuertos interface{}
	if err := json.Unmarshal(body, &aereopuertos); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(aereopuertos)
}

func getRoutes(c *fiber.Ctx) error {
	// Hacer la solicitud al microservicio
	res, err := http.Get("http://127.0.0.1:3006/rutas")
	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	//revisar status

	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	//leer respuesta

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	var rutas interface{}
	if err := json.Unmarshal(body, &rutas); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(rutas)
}

func getActiveRoutes(c *fiber.Ctx) error {
	res, err := http.Get("http://127.0.0.1:3006/rutas/active")
	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	//read
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	var rutas interface{}
	if err := json.Unmarshal(body, &rutas); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(rutas)
}

func getInactiveRoutes(c *fiber.Ctx) error {
	// Hacer la solicitud al microservicio
	res, err := http.Get("http://127.0.0.1:3006/rutas/inactive")
	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	//revisar status

	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	//leer respuesta

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	var rutas interface{}
	if err := json.Unmarshal(body, &rutas); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(rutas)
}

func activateRoute(c *fiber.Ctx) error {
	id := c.Params("id")
	var response interface{}

	res, err := http.Get("http://127.0.0.1:3006/rutas/accept/" + id)
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(response)
}

func createNewRoute(c *fiber.Ctx) error {
	//obtener NextID
	res, err := http.Get("http://127.0.0.1:3006/rutas/nextID")
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor (nextID)")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor (nextID)")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta (nextID)")
	}

	var nextID map[string]interface{}
	if err := json.Unmarshal(body, &nextID); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON (nextID)")
	}
	nextIDInt := int(nextID["next_id"].(float64))

	routeCreator, err := parseJSON(c.Body())
	if err != nil {
		return c.Status(400).SendString("Error al decodificar el JSON")
	}

	ruta := Ruta{
		Origen_id:         routeCreator.AirportFrom.ID,
		Destino_id:        routeCreator.AirportTo.ID,
		Duracion_minutos:  routeCreator.FlightTime.Horas*60 + routeCreator.FlightTime.Minutos,
		Distancia:         routeCreator.Distance,
		Frequencia_dia:    routeCreator.DailyFrequ,
		Frecuencia_semana: len(routeCreator.DaysAWeek),
		Activa:            0,
		Monday:            isSelected("Monday", routeCreator.SelectedDays),
		Tuesday:           isSelected("Tuesday", routeCreator.SelectedDays),
		Wednesday:         isSelected("Wednesday", routeCreator.SelectedDays),
		Thursday:          isSelected("Thursday", routeCreator.SelectedDays),
		Friday:            isSelected("Friday", routeCreator.SelectedDays),
		Saturday:          isSelected("Saturday", routeCreator.SelectedDays),
		Sunday:            isSelected("Sunday", routeCreator.SelectedDays),
		Type:              routeCreator.Tipo,
	}

	rutaJSON, err := json.Marshal(ruta)
	if err != nil {
		return c.Status(500).SendString("Error al codificar la ruta")
	}

	res, err = http.Post("http://127.0.0.1:3006/rutas/create", "application/json", bytes.NewBuffer(rutaJSON))
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor (rutas)")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		log.Printf("Error en la respuesta del servidor (rutas): %s", body)
		return c.Status(500).SendString("Error en la respuesta del servidor (rutas)")
	}

	for _, ticket := range routeCreator.TakeOffTickets {
		horario := Horario{
			Ruta_id:      nextIDInt,
			Hora_salida:  ticket.TakeOffTime,
			Hora_llegada: ticket.LandingTime,
		}

		horarioJSON, err := json.Marshal(horario)
		if err != nil {
			return c.Status(500).SendString("Error al codificar el horario")
		}

		res, err := http.Post("http://127.0.0.1:3006/rutas/schedule/create", "application/json", bytes.NewBuffer(horarioJSON))
		if err != nil {
			return c.Status(500).SendString("Error al contactar con el servidor (horarios)")
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(res.Body)
			log.Printf("Error en la respuesta del servidor (horarios): %s", body)
			return c.Status(500).SendString("Error en la respuesta del servidor (horarios)")
		}
	}

	for _, model := range routeCreator.SelectedModels {
		modeloRuta := ModeloRuta{
			Ruta_id:   nextIDInt,
			Modelo_id: model.ID,
		}

		modeloRutaJSON, err := json.Marshal(modeloRuta)
		if err != nil {
			return c.Status(500).SendString("Error al codificar el modeloRuta")
		}

		res, err := http.Post("http://127.0.0.1:3006/rutas/modelRoute/create", "application/json", bytes.NewBuffer(modeloRutaJSON))
		if err != nil {
			return c.Status(500).SendString("Error al contactar con el servidor (modelos)")
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(res.Body)
			log.Printf("Error en la respuesta del servidor (modelos): %s", body)
			return c.Status(500).SendString("Error en la respuesta del servidor (modelos)")
		}
	}

	return c.JSON("Ruta creada exitosamente")
}

func GenerateNewPin(c *fiber.Ctx) error {
	login := new(Login)
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON!",
		})
	}

	JSON, err := json.Marshal(login)

	if err != nil {
		return c.Status(500).SendString("Error al codificar el JSON")
	}

	response, error := http.Post("http://127.0.0.1:3001/users/modifyPIN", "application/json", bytes.NewBuffer(JSON))
	if error != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	data, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	return c.SendString(string(data))
}

func LoginFunc(c *fiber.Ctx) error {
	login := new(Login)
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON!",
		})
	}

	JSON, err := json.Marshal(login)
	if err != nil {
		return c.Status(500).SendString("Error al codificar el JSON")
	}

	response, error := http.Post("http://127.0.0.1:3001/users/login", "application/json", bytes.NewBuffer(JSON))
	if error != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	data, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	return c.SendString(string(data))
}

func getAviones(c *fiber.Ctx) error {
	res, err := http.Get("http://127.0.0.1:3002/aviones")

	if err != nil {
		log.Println("Error al contactar con el servidor:", err)
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Error en la respuesta del servidor:", res.Status)
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	var aviones interface{}
	if err := json.Unmarshal(body, &aviones); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(aviones)
}

func GetAvionPorID(c *fiber.Ctx) error {
	id := c.Params("id")
	var response interface{}

	res, err := http.Get("http://127.0.0.1:3002/aviones/" + id)
	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(response)
}

func checkMatricula(c *fiber.Ctx) error {
	matricula := c.Params("matricula")
	var response interface{}
	res, err := http.Get("http://127.0.0.1:3002/aviones/check/matricula/" + matricula)

	if err != nil {
		return c.Status(500).SendString("Error al contactar con el servidor")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(500).SendString("Error en la respuesta del servidor")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return c.Status(500).SendString("Error al leer el cuerpo de la respuesta")
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).SendString("Error al decodificar la respuesta JSON")
	}

	return c.JSON(fiber.Map{
		"available": response,
	})
}
