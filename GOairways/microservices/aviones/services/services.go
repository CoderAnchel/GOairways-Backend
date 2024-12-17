package services

import "aviones/models"

func CreateNewModelo(modelo *models.Modelo) error {
	return models.CreateNewModelo(modelo)
}

func GetAllModelos() ([]models.Modelo, error) {
	return models.GetAllModelos(), nil
}

func GetModeloByID(id int) (models.Modelo, error) {
	return models.GetModeloByID(id)
}

func GetModeloByTipo(tipo string) ([]models.Modelo, error) {
	return models.GetModelosByTipo(tipo)
}
