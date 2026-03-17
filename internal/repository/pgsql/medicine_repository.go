package pgsql

import (
	"auto-pharmacy/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineRepository struct {
	DB *pgxpool.Pool
}

func NewMedicineRepository(db *pgxpool.Pool) *MedicineRepository {
	return &MedicineRepository{DB: db}
}

func (r *MedicineRepository) GetAll(ctx context.Context) ([]models.Medicine, error) {
	var medicines []models.Medicine
	rows, err := r.DB.Query(ctx, "SELECT id, name, measurement, dose, measurement_dose, box, place FROM medicines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var medicine models.Medicine
		if err := rows.Scan(
			&medicine.ID,
			&medicine.Name,
			&medicine.Measurement,
			&medicine.Dose,
			&medicine.Measurement_dose,
			&medicine.Box,
			&medicine.Place,
		); err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}

	return medicines, nil
}

func (r *MedicineRepository) GetByID(ctx context.Context, id uint) (*models.Medicine, error) {
	var medicine models.Medicine
	err := r.DB.QueryRow(
		ctx,
		"SELECT id, name, measurement, dose, measurement_dose, box, place FROM medicines WHERE id=$1",
		id,
	).Scan(
		&medicine.ID,
		&medicine.Name,
		&medicine.Measurement,
		&medicine.Dose,
		&medicine.Measurement_dose,
		&medicine.Box,
		&medicine.Place,
	)
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *MedicineRepository) Create(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error) {
	var id uint
	if err := r.DB.QueryRow(
		ctx,
		"INSERT INTO medicines (name, measurement, dose, measurement_dose, box, place) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		medicine.Name,
		medicine.Measurement,
		medicine.Dose,
		medicine.Measurement_dose,
		medicine.Box,
		medicine.Place,
	).Scan(&id); err != nil {
		return nil, err
	}

	medicine.ID = id
	return medicine, nil
}

func (r *MedicineRepository) Update(ctx context.Context, medicine *models.Medicine) error {
	_, err := r.DB.Exec(
		ctx,
		"UPDATE medicines SET name=$1, measurement=$2, dose=$3, measurement_dose=$4, box=$5, place=$6 WHERE id=$7",
		medicine.Name,
		medicine.Measurement,
		medicine.Dose,
		medicine.Measurement_dose,
		medicine.Box,
		medicine.Place,
		medicine.ID,
	)
	return err
}

func (r *MedicineRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM medicines WHERE id=$1", id)
	return err
}

func (r *MedicineRepository) Restore(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "UPDATE medicines SET deleted_at=NULL WHERE id=$1", id)
	return err
}

func (r *MedicineRepository) ForceDelete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM medicines WHERE id=$1", id)
	return err
}

func (r *MedicineRepository) MassDelete(ctx context.Context, ids []uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM medicines WHERE id = ANY($1)", ids)
	return err
}
