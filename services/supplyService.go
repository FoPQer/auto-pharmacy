package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
)

func GetAllSupplies() ([]models.MedicineSupply, error) {
	var supplies = make([]models.MedicineSupply, 0)
	if err := database.MysqlDB.DB.Model(&models.MedicineSupply{}).Preload("Medicine", &models.Medicine{}).Find(&supplies).Error; err != nil {
		return nil, errors.Join(errors.New("Supplies Find error"), err)
	}

	return supplies, nil
}

func GetSupply(id string) (models.MedicineSupply, error) {
	var supply models.MedicineSupply
	if err := database.MysqlDB.DB.Model(&models.MedicineSupply{}).Preload("Medicine", &models.Medicine{}).First(&supply, id).Error; err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply First error"), err)
	}

	return supply, nil
}

func SetSupply(body *json.Decoder) (models.MedicineSupply, error) {
	supply := models.MedicineSupply{}

	if err := body.Decode(&supply); err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply decode error"), err)
	}
	if err := database.MysqlDB.DB.Create(&supply).Error; err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply create error"), err)
	}

	return supply, nil
}
