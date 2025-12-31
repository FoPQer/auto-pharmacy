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
	if err := body.Decode(&medicine); err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine decode error"), err)
	}
	if err := database.MysqlDB.DB.Create(&medicine).Error; err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine create error"), err)
	}

	return medicine, nil
}

func UpdateMedicine(id string, body *json.Decoder) (models.Medicine, error) {
	medicine, err := GetMedicine(id)
	if err != nil {
		return models.Medicine{}, err
	}
	if err := body.Decode(&medicine); err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine decode error"), err)
	}
	if err := database.MysqlDB.DB.Save(&medicine).Error; err != nil {
		return models.Medicine{}, errors.Join(errors.New("Medicine save error"), err)
	}

	return medicine, nil
}

func DeleteMedicine(id string) error {
	if err := database.MysqlDB.DB.Delete(&models.Medicine{}, id).Error; err != nil {
		return errors.Join(errors.New("Medicine delete error"), err)
	}
	return nil
}

func RestoreMedicine(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Model(&models.Medicine{}).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return errors.Join(errors.New("Medicine restore error"), err)
	}
	return nil
}

func ForceDeleteMedicine(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Association("Supplies").Clear(); err != nil {
		return errors.Join(errors.New("MedicineSupply clear error"), err)
	}
	if err := database.MysqlDB.DB.Unscoped().Delete(&models.Medicine{}, id).Error; err != nil {
		return errors.Join(errors.New("Medicine force delete error"), err)
	}
	return nil
}
