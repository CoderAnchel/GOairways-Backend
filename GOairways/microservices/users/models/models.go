package models

import (
	"errors"
	"time"
	"users/database"
	"users/services"
)

type User struct {
	ID                int    `json:"id" gorm:"primaryKey"`
	Id_rol            int    `json:"id_rol"`
	Nombre            string `json:"nombre"`
	Apellido          string `json:"apellido"`
	Email             string `json:"email"`
	Telefono          string `json:"telefono"`
	Password          string `json:"password"`
	Fecha_naciemiento string `json:"fecha_nacimiento"`
	Nacionalidad      string `json:"nacionalidad"`
	Direccion         string `json:"direccion"`
	Ciudad            string `json:"ciudad"`
	Codigo_postal     string `json:"codigo_postal"`
	Pais              string `json:"pais"`
	Pasaporte         string `json:"pasaporte"`
	DNI               string `json:"dni"`
	PIN               int    `json:"pin"`
}

type Login struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
	PIN      int    `json:"pin"`
}

func (u *User) TableName() string {
	return "usuarios"
}

func CreateUser(user *User) error {
	return database.DB.Create(user).Error
}

func LoginFunc(DNI string, password string, PIN int) (interface{}, error) {
	var user User

	result := database.DB.Where("dni = ?", DNI).First(&user)
	if result.Error != nil {
		return nil, nil
	}

	check := services.CheckPasswordHash(password, user.Password)

	if !check || user.PIN != PIN {
		return nil, nil
	}

	return user, nil
}

func ModyfyPIN(login *Login) error {
	var user User

	// Verificar si el DNI es correcto
	result := database.DB.Where("dni = ?", login.DNI).First(&user)
	if result.Error != nil {
		return result.Error
	}

	// Verificar si la contraseña es correcta
	check := services.CheckPasswordHash(login.Password, user.Password)
	if !check {
		return errors.New("invalid password")
	}

	// Generar y actualizar el PIN
	pin := services.GenerateRandomPIN()
	response := database.DB.Model(&User{}).Where("dni = ?", login.DNI).Update("PIN", pin).Error
	if response != nil {
		return response
	}

	services.SendMail(pin)

	go func(dni string) {
		time.Sleep(2 * time.Minute)
		newPin := services.GenerateRandomPIN()
		database.DB.Model(&User{}).Where("dni = ?", dni).Update("PIN", newPin)
	}(login.DNI)

	return nil
}
