package services

import (
	"auto-pharmacy/internal/models"
	"auto-pharmacy/internal/repository"
	"context"
	"errors"
	"fmt"
	"log"
)

type ErrUserNotFound struct {
	ID uint
	Email string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("User not found with ID: %d or Email: %s", e.ID, e.Email)
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type GetUserRequest struct {
	ID uint `json:"id" param:"user" validate:"required_without=Email"`
	Email string `json:"email" query:"email" validate:"required_without=ID,email"`
}

type CreateUserRequest struct {
    Name     string `json:"name" validate:"min=2,max=255"`
    Email    string `json:"email" validate:"email,max=255"`
    Password string `json:"password" validate:"min=8,max=72"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=8,max=72"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type UserService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*UserResponse, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []*UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponses, nil
}

func (s *UserService) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	var user *models.User
	var err error
	switch {
		case req.ID != 0:
			user, err = s.repo.GetByID(ctx, req.ID)
		case req.Email != "":
			user, err = s.repo.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}

		return &UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
	return nil, &ErrUserNotFound{ID: req.ID, Email: req.Email}
}

func (s *UserService) SetUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	user, err := models.NewUser(
		req.Email,
		req.Password,
		req.Name,
	)
	user.Password, err = models.HashPassword(user.Password)
	if err != nil {
		log.Printf("Hashing Password error: %s", err.Error())
		return nil, errors.Join(errors.New("Hashing Password error"), err)
	}

	user, err = s.repo.Create(ctx, user)
	if err != nil {
		log.Printf("User create error: %s", err.Error())
		return nil, errors.Join(errors.New("User create error"), err)
	}

	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req *UpdateUserRequest) (*UserResponse, error) {
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    if req.Name != nil {
        user.Name = *req.Name
    }
    if req.Email != nil {
        user.Email = *req.Email
    }
    if req.Password != nil {
        hashed, err := models.HashPassword(*req.Password)
        if err != nil {
            return nil, err
        }
        user.Password = hashed
    }

    if err := s.repo.Update(ctx, user); err != nil {
        return nil, err
    }

    return &UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uint, req *ChangePasswordRequest) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	user.Password, err = models.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("Hashing Password error: %w", err)
	}

	if err := s.repo.Update(ctx, user); err != nil {
        return fmt.Errorf("User update error: %w", err)
    }

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("User delete error: %w", err)
	}
	return nil
}

func (s *UserService) RestoreUser(ctx context.Context, id uint) error {
	if err := s.repo.Restore(ctx, id); err != nil {
		return fmt.Errorf("User restore error: %w", err)
	}
	return nil
}

func (s *UserService) ForceDeleteUser(ctx context.Context, id uint) error {
	if err := s.repo.ForceDelete(ctx, id); err != nil {
		return fmt.Errorf("User force delete error: %w", err)
	}
	return nil
}

func (s *UserService) MassDeleteUser(ctx context.Context, ids []uint) error {
	if err := s.repo.MassDelete(ctx, ids); err != nil {
		return fmt.Errorf("User mass delete error: %w", err)
	}
	return nil
}
