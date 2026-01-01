package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
)

func GetAllTags() ([]models.Tag, error) {
	var tags = make([]models.Tag, 0)
	if err := database.MysqlDB.DB.Model(&models.Tag{}).Find(&tags).Error; err != nil {
		return nil, errors.Join(errors.New("Tags Find error"), err)
	}

	return tags, nil
}

func GetTag(id string) (models.Tag, error) {
	var tag models.Tag
	if err := database.MysqlDB.DB.Model(&models.Tag{}).First(&tag, id).Error; err != nil {
		return models.Tag{}, errors.Join(errors.New("Tag First error"), err)
	}

	return tag, nil
}

func SetTag(body *json.Decoder) (models.Tag, error) {
	tag := models.Tag{}
	if err := body.Decode(&tag); err != nil {
		return models.Tag{}, errors.Join(errors.New("Tag decode error"), err)
	}
	if err := database.MysqlDB.DB.Create(&tag).Error; err != nil {
		return models.Tag{}, errors.Join(errors.New("Tag create error"), err)
	}

	return tag, nil
}

func UpdateTag(id string, body *json.Decoder) (models.Tag, error) {
	tag, err := GetTag(id)
	if err != nil {
		return models.Tag{}, err
	}
	if err := body.Decode(&tag); err != nil {
		return models.Tag{}, errors.Join(errors.New("Tag decode error"), err)
	}
	if err := database.MysqlDB.DB.Save(&tag).Error; err != nil {
		return models.Tag{}, errors.Join(errors.New("Tag save error"), err)
	}

	return tag, nil
}

func DeleteTag(id string) error {
	if err := database.MysqlDB.DB.Delete(&models.Tag{}, id).Error; err != nil {
		return errors.Join(errors.New("Tag delete error"), err)
	}
	return nil
}

func RestoreTag(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Model(&models.Tag{}).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return errors.Join(errors.New("Tag restore error"), err)
	}
	return nil
}

func ForceDeleteTag(id string) error {
	if err := database.MysqlDB.DB.Unscoped().Delete(&models.Tag{}, id).Error; err != nil {
		return errors.Join(errors.New("Tag force delete error"), err)
	}
	return nil
}
