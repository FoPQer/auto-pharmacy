package models

import (
	"auto-pharmacy/internal/database"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string
}

func NewTag(name string) *Tag {
	return &Tag{Name: name}
}

func (t *Tag) AppendTag(med *Medicine) error {
	if err := database.MysqlDB.DB.Model(med).Association("Tags").Append(t); err != nil {
		return err
	}
	return nil
}
