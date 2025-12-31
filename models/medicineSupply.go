package models

import (
	"auto-pharmacy/database"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MedicineSupply struct {
	gorm.Model
	Count      float64   `gorm:"count" json:"count"`
	Expired_at time.Time `gorm:"expired_at" json:"expired_at"`
	MedicineID uint      `json:"medicine_id"`
	Medicine   Medicine
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
