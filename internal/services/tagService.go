package services

import (
	"auto-pharmacy/internal/models"
	"auto-pharmacy/internal/repository"
	"context"
	"fmt"
)

type TagResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
}

type GetTagRequest struct {
	ID uint `json:"id" param:"tag" validate:"required"`
}

type CreateTagRequest struct {
	Name     string `json:"name" validate:"min=2,max=255"`
}

type UpdateTagRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
}

type TagService struct{
	repo repository.TagRepository
}

func (s *TagService) GetAllTags(ctx context.Context) ([]*TagResponse, error) {
	tags, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Tags GetAll error: %w", err)
	}

	var response = make([]*TagResponse, len(tags))
	for i, tag := range tags {
		response[i] = &TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}
	return response, nil
}

func (s *TagService) GetTag(ctx context.Context, id uint) (*TagResponse, error) {
	tag, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Tag GetByID error: %w", err)
	}

	return &TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (s *TagService) SetTag(ctx context.Context, req *CreateTagRequest) (*TagResponse, error) {
	tag := models.NewTag(req.Name)
	createdTag, err := s.repo.Create(ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("Tag Create error: %w", err)
	}

	return &TagResponse{
		ID:   createdTag.ID,
		Name: createdTag.Name,
	}, nil
}

func (s *TagService) UpdateTag(ctx context.Context, id uint, req *UpdateTagRequest) (*TagResponse, error) {
	tag, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Tag GetByID error: %w", err)
	}
	if req.Name != nil {
		tag.Name = *req.Name
	}
	
	if err := s.repo.Update(ctx, tag); err != nil {
		return nil, fmt.Errorf("Tag Update error: %w", err)
	}

	return &TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (s *TagService) DeleteTag(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("Tag delete error: %w", err)
	}
	return nil
}

func (s *TagService) RestoreTag(ctx context.Context, id uint) error {
	if err := s.repo.Restore(ctx, id); err != nil {
		return fmt.Errorf("Tag restore error: %w", err)
	}
	return nil
}

func (s *TagService) ForceDeleteTag(ctx context.Context, id uint) error {
	if err := s.repo.ForceDelete(ctx, id); err != nil {
		return fmt.Errorf("Tag force delete error: %w", err)
	}
	return nil
}
