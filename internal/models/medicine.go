package models

import (
	"auto-pharmacy/internal/database"
	"errors"
	"time"

	"gorm.io/gorm"
)

func MedMigrate() error {
	if err := MedicineMigrate(); err != nil {
		return err
	}
	if err := MedicineSupplyMigrate(); err != nil {
		return err
	}
	return nil
}

type Medicine struct {
	gorm.Model
	Name             string  `gorm:"name" json:"name"`
	Measurement      *string `gorm:"measurement" json:"measurement"`
	Dose             float64 `gorm:"dose" json:"dose"`
	Measurement_dose *string `gorm:"measurementDose" json:"measurementDose"`
	Box              *string `gorm:"box" json:"box"`
	Place            *string `gorm:"place" json:"place"`
	Supplies         []MedicineSupply
	Tags             []Tag `gorm:"many2many:medicine_tag;"`
}

func MedicineMigrate() error {
	return database.MysqlDB.DB.AutoMigrate(&Medicine{})
}

func NewMedicine(name string, measurement *string, dose float64, measurement_dose *string, box *string, place *string) *Medicine {
	return &Medicine{
		Name:             name,
		Measurement:      measurement,
		Dose:             dose,
		Measurement_dose: measurement_dose,
		Box:              box,
		Place:            place,
	}
}

func (m Medicine) Supply(count float64, expired_at time.Time) *MedicineSupply {
	return NewMedicineSupply(count, expired_at, m.ID)
}

func (med *Medicine) AppendTag(t *Tag) error {
	if err := database.MysqlDB.DB.Model(med).Association("Tags").Append(t); err != nil {
		return errors.Join(errors.New("Medicine Append Tag error"), err)
	}
	return nil
}
