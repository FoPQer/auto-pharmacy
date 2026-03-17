package pgsql

import (
	"auto-pharmacy/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineSupplyRepository struct {
	DB *pgxpool.Pool
}

func NewMedicineSupplyRepository(db *pgxpool.Pool) *MedicineSupplyRepository {
	return &MedicineSupplyRepository{DB: db}
}

func (r *MedicineSupplyRepository) GetAll(ctx context.Context) ([]models.MedicineSupply, error) {
	var supplies []models.MedicineSupply
	rows, err := r.DB.Query(ctx, "SELECT id, medicine_id, quantity, expired_at FROM supplies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supply models.MedicineSupply
		if err := rows.Scan(
			&supply.ID,
			&supply.MedicineID,
			&supply.Quantity,
			&supply.ExpiredAt,
		); err != nil {
			return nil, fmt.Errorf("GetAll scan error: %w", err)
		}
		supplies = append(supplies, supply)
	}

	return supplies, nil
}

func (r *MedicineSupplyRepository) GetByID(ctx context.Context, id uint) (*models.MedicineSupply, error) {
	var supply models.MedicineSupply
	err := r.DB.QueryRow(
		ctx,
		"SELECT id, medicine_id, quantity, expired_at FROM supplies WHERE id=$1",
		id,
	).Scan(
		&supply.ID,
		&supply.MedicineID,
		&supply.Quantity,
		&supply.ExpiredAt,
	)
	if err != nil {
		return nil, fmt.Errorf("GetByID error: %w", err)
	}
	return &supply, nil
}

func (r *MedicineSupplyRepository) GetOlderExpiringSupplyByMedicineID(ctx context.Context, medicineID uint) (*models.MedicineSupply, error) {
	var supply models.MedicineSupply
	err := r.DB.QueryRow(
		ctx,
		"SELECT id, medicine_id, quantity, expired_at FROM supplies WHERE expired_at > NOW() AND medicine_id=$1 ORDER BY expired_at ASC LIMIT 1",
		medicineID,
	).Scan(
		&supply.ID,
		&supply.MedicineID,
		&supply.Quantity,
		&supply.ExpiredAt,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOlderExpiringSupplyByMedicineID error: %w", err)
	}
	return &supply, nil
}


func (r *MedicineSupplyRepository) GetByMedicineID(ctx context.Context, medicineID uint) ([]models.MedicineSupply, error) {
	var supplies []models.MedicineSupply
	rows, err := r.DB.Query(
		ctx,
		"SELECT id, medicine_id, quantity, expired_at FROM supplies WHERE medicine_id=$1",
		medicineID,
	)
	if err != nil {
		return nil, fmt.Errorf("GetByMedicineID error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var supply models.MedicineSupply
		if err := rows.Scan(
			&supply.ID,
			&supply.MedicineID,
			&supply.Quantity,
			&supply.ExpiredAt,
		); err != nil {
			return nil, fmt.Errorf("GetByMedicineID scan error: %w", err)
		}
		supplies = append(supplies, supply)
	}

	return supplies, nil
}

func (r *MedicineSupplyRepository) Create(ctx context.Context, supply *models.MedicineSupply) (*models.MedicineSupply, error) {
	var id uint
	if err := r.DB.QueryRow(
		ctx,
		"INSERT INTO supplies (medicine_id, quantity, expired_at) VALUES ($1, $2, $3) RETURNING id",
		supply.MedicineID,
		supply.Quantity,
		supply.ExpiredAt,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("Create supply error: %w", err)
	}

	supply.ID = id
	return supply, nil
}

func (r *MedicineSupplyRepository) Update(ctx context.Context, supply *models.MedicineSupply) error {
	_, err := r.DB.Exec(
		ctx,
		"UPDATE supplies SET medicine_id=$1, quantity=$2, expired_at=$3 WHERE id=$4",
		supply.MedicineID,
		supply.Quantity,
		supply.ExpiredAt,
		supply.ID,
	)
	if err != nil {
		return fmt.Errorf("Update supply error: %w", err)
	}
	return nil
}

func (r *MedicineSupplyRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM supplies WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Delete supply error: %w", err)
	}
	return nil
}

func (r *MedicineSupplyRepository) Restore(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "UPDATE supplies SET deleted_at=NULL WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Restore supply error: %w", err)
	}
	return nil
}

func (r *MedicineSupplyRepository) ForceDelete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM supplies WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Force delete supply error: %w", err)
	}
	return nil
}

func (r *MedicineSupplyRepository) MassDelete(ctx context.Context, ids []uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM supplies WHERE id = ANY($1)", ids)
	if err != nil {
		return fmt.Errorf("Mass delete supply error: %w", err)
	}
	return nil
}
