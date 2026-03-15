package pgsql

import (
	"auto-pharmacy/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	rows, err := r.DB.Query(ctx, "SELECT id, name, email, password_hash FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(ctx, "SELECT id, name, email, password_hash FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(ctx, "SELECT id, name, email, password_hash FROM users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	var id uint
	if err := r.DB.QueryRow(
		ctx, 
		"INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id", 
		user.Name, 
		user.Email,
		user.Password,
	).Scan(&id); err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.DB.Exec(ctx, "UPDATE users SET name=$1, email=$2, password_hash=$3 WHERE id=$4", user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}

func (r *UserRepository) Restore(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "UPDATE users SET deleted_at=NULL WHERE id=$1", id)
	return err
}

func (r *UserRepository) ForceDelete(ctx context.Context, id uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}

func (r *UserRepository) MassDelete(ctx context.Context, ids []uint) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM users WHERE id = ANY($1)", ids)
	return err
}