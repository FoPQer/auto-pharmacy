package repository

import (
	"auto-pharmacy/internal/models"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	Restore(ctx context.Context, id uint) error
	ForceDelete(ctx context.Context, id uint) error
	MassDelete(ctx context.Context, ids []uint) error
}

type MedicineRepository interface {
	Create(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error)
	GetByID(ctx context.Context, id uint) (*models.Medicine, error)
	GetAll(ctx context.Context) ([]models.Medicine, error)
	Update(ctx context.Context, medicine *models.Medicine) error
	Delete(ctx context.Context, id uint) error
	Restore(ctx context.Context, id uint) error
	ForceDelete(ctx context.Context, id uint) error
	MassDelete(ctx context.Context, ids []uint) error
}

type MedicineTagRepository interface {
    Associate(ctx context.Context, medicineID uint, tagID uint) error
    Dissociate(ctx context.Context, medicineID uint, tagID uint) error
    GetTagsByMedicineID(ctx context.Context, medicineID uint) ([]models.Tag, error)
}

type MedicineSupplyRepository interface {
	Create(ctx context.Context, supply *models.MedicineSupply) (*models.MedicineSupply, error)
	GetByID(ctx context.Context, id uint) (*models.MedicineSupply, error)
	GetOlderExpiringSupplyByMedicineID(ctx context.Context, medicineID uint) (*models.MedicineSupply, error)
	GetByMedicineID(ctx context.Context, medicineID uint) ([]models.MedicineSupply, error)
	GetAll(ctx context.Context) ([]models.MedicineSupply, error)
	Update(ctx context.Context, supply *models.MedicineSupply) error
	Delete(ctx context.Context, id uint) error
	Restore(ctx context.Context, id uint) error
	ForceDelete(ctx context.Context, id uint) error
	MassDelete(ctx context.Context, ids []uint) error
}

type TagRepository interface {
	Create(ctx context.Context, tag *models.Tag) (*models.Tag, error)
	GetByID(ctx context.Context, id uint) (*models.Tag, error)
	GetAll(ctx context.Context) ([]models.Tag, error)
	Update(ctx context.Context, tag *models.Tag) error
	Delete(ctx context.Context, id uint) error
	Restore(ctx context.Context, id uint) error
	ForceDelete(ctx context.Context, id uint) error
	MassDelete(ctx context.Context, ids []uint) error
}