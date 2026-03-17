package models

import (
	"auto-pharmacy/internal/database"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type MedicineSupply struct {
	gorm.Model
	Quantity      float64   `gorm:"Quantity" json:"Quantity"`
	ExpiredAt time.Time `gorm:"expired_at" json:"expired_at"`
	MedicineID uint      `json:"medicine_id"`
	Medicine   *Medicine
}

func MedicineSupplyMigrate() error {
	return database.MysqlDB.DB.AutoMigrate(&MedicineSupply{})
}

func NewMedicineSupply(quantity float64, expiredAt time.Time, medicineId uint) *MedicineSupply {
	return &MedicineSupply{
		Quantity:      quantity,
		ExpiredAt: expiredAt,
		MedicineID: medicineId,
	}
}

func (ms *MedicineSupply) Release() error {
	m := ms.Medicine
	if ms.Quantity > m.Dose {
		ms.Quantity = ms.Quantity - m.Dose
		return nil
	}

	return MedicineReleaseErr{*ms, errors.New("dose " + strconv.FormatFloat(m.Dose, 'f', 2, 64) + " greater than Quantity " + strconv.FormatFloat(ms.Quantity, 'f', 2, 64))}
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
