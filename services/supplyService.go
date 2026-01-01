package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
	"time"
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

func UpdateSupply(id string, body *json.Decoder) (models.MedicineSupply, error) {
	supply, err := GetSupply(id)
	if err != nil {
		return models.MedicineSupply{}, err
	}
	if err := body.Decode(&supply); err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply decode error"), err)
	}
	if err := database.MysqlDB.DB.Save(&supply).Error; err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply save error"), err)
	}

	return supply, nil
}

func DeleteSupply(id string) error {
	if err := database.MysqlDB.DB.Delete(&models.MedicineSupply{}, id).Error; err != nil {
		return errors.Join(errors.New("Supply delete error"), err)
	}
	return nil
}

func RestoreSupply(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Model(&models.MedicineSupply{}).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return errors.Join(errors.New("Supply restore error"), err)
	}
	return nil
}

func ForceDeleteSupply(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Delete(&models.MedicineSupply{}, id).Error; err != nil {
		return errors.Join(errors.New("Supply force delete error"), err)
	}
	return nil
}

func ReleaseSupply(id string) (models.MedicineSupply, error) {
	var supply models.MedicineSupply
	err := database.MysqlDB.DB.Model(models.MedicineSupply{}).Order("expired_at").Where("expired_at > ? AND medicine_id = ?", time.Now(), id).Preload("Medicine").Find(&supply).Error
	if err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply finding error "), err)
	}
	if err := supply.Release(); err != nil {
		return models.MedicineSupply{}, errors.Join(errors.New("Supply release error "), err)
	}
	database.MysqlDB.DB.Save(&supply)
	return supply, nil
}
