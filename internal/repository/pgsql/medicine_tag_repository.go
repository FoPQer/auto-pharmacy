package pgsql

import (
	"auto-pharmacy/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ErrAssociate struct {
	MedicineID uint
	TagID      uint
	err        error
}

func (e *ErrAssociate) Error() string {
	return fmt.Sprintf("failed to associate tag %d with medicine %d: %v", e.TagID, e.MedicineID, e.err)
}

func (e *ErrAssociate) Unwrap() error {
	return e.err
}

type MedicineTagRepository struct {
	DB *pgxpool.Pool
}

func (r *MedicineTagRepository) Associate(ctx context.Context, medicineID, tagID uint) error {
	_, err := r.DB.Exec(ctx,
		"INSERT INTO medicine_tag (medicine_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		medicineID, tagID,
	)
	if err != nil {
		return &ErrAssociate{MedicineID: medicineID, TagID: tagID, err: err}
	}
	return nil
}

func (r *MedicineTagRepository) Dissociate(ctx context.Context, medicineID, tagID uint) error {
	_, err := r.DB.Exec(ctx,
		"DELETE FROM medicine_tag WHERE medicine_id=$1 AND tag_id=$2",
		medicineID, tagID,
	)
	if err != nil {
		return fmt.Errorf("MedicineTag dissociate error: %w", err)
	}
	return nil
}

func (r *MedicineTagRepository) GetTagsByMedicineID(ctx context.Context, medicineID uint) ([]models.Tag, error) {
	var tags []models.Tag
	rows, err := r.DB.Query(ctx,
		`SELECT t.id, t.name
		 FROM tags t
		 JOIN medicine_tag mt ON t.id = mt.tag_id
		 WHERE mt.medicine_id = $1`, medicineID,
	)
	if err != nil {
		return nil, fmt.Errorf("MedicineTag get tags by medicine ID error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, fmt.Errorf("MedicineTag scan tag error: %w", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MedicineTag rows error: %w", err)
	}

	return tags, nil
}