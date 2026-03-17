package pgsql

import (
	"auto-pharmacy/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ErrTagNotFound struct {
	ID uint
	err error
}

func (e *ErrTagNotFound) Error() string {
	return "Tag not found"
}

func (e *ErrTagNotFound) Unwrap() error {
	return e.err
}

type TagRepository struct {
	DB *pgxpool.Pool
}

func NewTagRepository(db *pgxpool.Pool) *TagRepository {
	return &TagRepository{DB: db}
}

func (r *TagRepository) GetAll(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag
	rows, err := r.DB.Query(ctx, "SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *TagRepository) GetByID(ctx context.Context, id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.DB.QueryRow(ctx, "SELECT id, name FROM tags WHERE id=$1", id).Scan(&tag.ID, &tag.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, &ErrTagNotFound{ID: id, err: err}
	}
	if err != nil {
		return nil, fmt.Errorf("Tag GetByID error: %w", err)
	}
	return &tag, nil
}

func (r *TagRepository) Create(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	var id uint
	if err := r.DB.QueryRow(
		ctx,
		"INSERT INTO tags (name) VALUES ($1) RETURNING id",
		tag.Name,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("Tag Create error: %w", err)
	}
	
	tag.ID = id
	return tag, nil
}

func (r *TagRepository) Update(ctx context.Context, tag *models.Tag) error {
	rows, err := r.DB.Exec(ctx, "UPDATE tags SET name=$1 WHERE id=$2", tag.Name, tag.ID)
	if err != nil {
		return fmt.Errorf("Tag Update error: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return &ErrTagNotFound{ID: tag.ID, err: fmt.Errorf("Tag Update error: no rows affected")}
	}
	return nil
}

func (r *TagRepository) Delete(ctx context.Context, id uint) error {
	rows, err := r.DB.Exec(ctx, "DELETE FROM tags WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Tag Delete error: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return &ErrTagNotFound{ID: id, err: fmt.Errorf("Tag Delete error: no rows affected")}
	}
	return nil
}

func (r *TagRepository) Restore(ctx context.Context, id uint) error {
	rows, err := r.DB.Exec(ctx, "UPDATE tags SET deleted_at=NULL WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Tag Restore error: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return &ErrTagNotFound{ID: id, err: fmt.Errorf("Tag Restore error: no rows affected")}
	}
	return nil
}

func (r *TagRepository) ForceDelete(ctx context.Context, id uint) error {
	rows, err := r.DB.Exec(ctx, "DELETE FROM tags WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Tag ForceDelete error: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return &ErrTagNotFound{ID: id, err: fmt.Errorf("Tag ForceDelete error: no rows affected")}
	}
	return nil
}

func (r *TagRepository) MassDelete(ctx context.Context, ids []uint) error {
	rows, err := r.DB.Exec(ctx, "DELETE FROM tags WHERE id = ANY($1)", ids)
	if err != nil {
		return fmt.Errorf("Tag MassDelete error: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return &ErrTagNotFound{ID: 0, err: fmt.Errorf("Tag MassDelete error: no rows affected")}
	}
	return nil
}
