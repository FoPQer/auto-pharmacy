package models

import (
	"auto-pharmacy/database"
	"fmt"
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
	Name             string  `gorm:"name"`
	Measurement      *string `gorm:"measurement"`
	Dose             float64 `gorm:"dose"`
	Measurement_dose *string `gorm:"measurementDose"`
	Box              *string `gorm:"box"`
	Place            *string `gorm:"place"`
	Supplies         []MedicineSupply
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

type MedicineSupply struct {
	gorm.Model
	Count      float64   `gorm:"count"`
	Expired_at time.Time `gorm:"expired_at"`
	MedicineID uint
}

func MedicineSupplyMigrate() error {
	return database.MysqlDB.DB.AutoMigrate(&MedicineSupply{})
}

func NewMedicineSupply(count float64, expired_at time.Time, medicineId uint) *MedicineSupply {
	return &MedicineSupply{
		Count:      count,
		Expired_at: expired_at,
		MedicineID: medicineId,
	}
}

func (m Medicine) Supply(count float64, expired_at time.Time) *MedicineSupply {
	return NewMedicineSupply(count, expired_at, m.ID)
}

func (ms MedicineSupply) Release() (bool, error) {
	var m Medicine
	database.MysqlDB.DB.Model(&ms).Association("Medicine").Find(&m)
	if ms.Count > m.Dose {
		ms.Count = ms.Count - m.Dose
		return true, nil
	}

	return false, MedicineReleaseErr{ms, fmt.Errorf("dose greater than count")}
}

type MedicineReleaseErr struct {
	Med MedicineSupply
	Err error
}

func (me MedicineReleaseErr) Error() string {
	var m Medicine
	database.MysqlDB.DB.Model(&me.Med).Association("Medicine").Find(&m)
	return fmt.Sprintf("%s medicine error: %v", m.Name, me.Err)
}

func (me MedicineReleaseErr) Unwrap() error {
	return me.Err
}
