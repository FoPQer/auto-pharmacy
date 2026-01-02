package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
)

func GetAllUsers() ([]models.User, error) {
	var users = make([]models.User, 0)
	if err := database.MysqlDB.DB.Model(&models.User{}).Omit("password").Find(&users).Error; err != nil {
		return nil, errors.Join(errors.New("Users Find error"), err)
	}

	return users, nil
}

func GetUser(id string) (models.User, error) {
	var user models.User
	if err := database.MysqlDB.DB.Model(&models.User{}).First(&user, id).Error; err != nil {
		return models.User{}, errors.Join(errors.New("User First error"), err)
	}

	return user, nil
}

func SetUser(body *json.Decoder) (models.User, error) {
	user := models.User{}
	if err := body.Decode(&user); err != nil {
		return models.User{}, errors.Join(errors.New("User decode error"), err)
	}
	if err := database.MysqlDB.DB.Create(&user).Error; err != nil {
		return models.User{}, errors.Join(errors.New("User create error"), err)
	}

	return user, nil
}

func UpdateUser(id string, body *json.Decoder) (models.User, error) {
	user, err := GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	if err := body.Decode(&user); err != nil {
		return models.User{}, errors.Join(errors.New("User decode error"), err)
	}
	if err := database.MysqlDB.DB.Save(&user).Error; err != nil {
		return models.User{}, errors.Join(errors.New("User save error"), err)
	}

	return user, nil
}

func DeleteUser(id string) error {
	if err := database.MysqlDB.DB.Delete(&models.User{}, id).Error; err != nil {
		return errors.Join(errors.New("User delete error"), err)
	}
	return nil
}

func RestoreUser(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return errors.Join(errors.New("User restore error"), err)
	}
	return nil
}

func ForceDeleteUser(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		return errors.Join(errors.New("User force delete error"), err)
	}
	return nil
}
