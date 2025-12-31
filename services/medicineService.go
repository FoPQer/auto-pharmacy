package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
)

func GetAllMedicines() ([]models.Medicine, error) {
	var medicines = make([]models.Medicine, 0)
	if err := database.MysqlDB.DB.Model(&models.Medicine{}).Preload("Supplies", &models.MedicineSupply{}).Find(&medicines).Error; err != nil {
		return nil, errors.Join(errors.New("Medicines Find error"), err)
	}

	return medicines, nil
}

func GetMedicine(id string) (models.Medicine, error) {
	var medicine models.Medicine
	if err := database.MysqlDB.DB.Model(&models.Medicine{}).Preload("Supplies", &models.MedicineSupply{}).First(&medicine, id).Error; err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine First error"), err)
	}

	return medicine, nil
}

func SetMedicine(body *json.Decoder) (models.Medicine, error) {
	medicine := models.Medicine{}
	supply := models.MedicineSupply{}
	if err := body.Decode(&medicine); err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine decode error"), err)
	}
	if err := database.MysqlDB.DB.Create(&medicine).Error; err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine create error"), err)
	}

	if err := body.Decode(&supply); err != nil {
		return models.Medicine{}, errors.Join(errors.New("Supply decode error"), err)
	}
	supply.MedicineID = medicine.ID
	if err := database.MysqlDB.DB.Create(&supply).Error; err != nil {
		return models.Medicine{}, errors.Join(errors.New("Supply create error"), err)
	}

	return medicine, nil
}
