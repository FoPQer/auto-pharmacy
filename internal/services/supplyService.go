package services

import (
	"auto-pharmacy/internal/database"
	"auto-pharmacy/internal/models"
	"auto-pharmacy/internal/repository"
	"context"
	"fmt"
	"time"
)

type MedicineSupplyResponse struct {
	ID         uint      `json:"id"`
	MedicineID uint      `json:"medicine_id"`
	Quantity   float64       `json:"quantity"`
	ExpiredAt  time.Time `json:"expired_at"`
}

type CreateMedicineSupplyRequest struct {
	MedicineID uint      `json:"medicine_id" validate:"required"`
	Quantity   float64   `json:"quantity" validate:"required,min=1"`
	ExpiredAt  time.Time `json:"expired_at" validate:"required"`
}

type UpdateMedicineSupplyRequest struct {
	MedicineID *uint      `json:"medicine_id,omitempty"`
	Quantity   *float64   `json:"quantity,omitempty" validate:"omitempty,min=1"`
	ExpiredAt  *time.Time `json:"expired_at,omitempty"`
}

type MedicineSupplyService struct{
	repo repository.MedicineSupplyRepository
}

func NewMedicineSupplyService(repo repository.MedicineSupplyRepository) *MedicineSupplyService {
	return &MedicineSupplyService{repo: repo}
}

func (s *MedicineSupplyService) GetAllSupplies(ctx context.Context) ([]*MedicineSupplyResponse, error) {
	supplies, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Supplies Find error: %w", err)
	}

	var response = make([]*MedicineSupplyResponse, 0, len(supplies))
	for _, supply := range supplies {
		response = append(response, &MedicineSupplyResponse{
			ID:         supply.ID,
			MedicineID: supply.MedicineID,
			Quantity:   supply.Quantity,
			ExpiredAt:  supply.ExpiredAt,
		})
	}

	return response, nil
}

func (s *MedicineSupplyService) GetSupply(ctx context.Context, id uint) (*MedicineSupplyResponse, error) {
	supply, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Supply Find error: %w", err)
	}
	response := &MedicineSupplyResponse{
		ID:         supply.ID,
		MedicineID: supply.MedicineID,
		Quantity:   supply.Quantity,
		ExpiredAt:  supply.ExpiredAt,
	}
	return response, nil
}

func (s *MedicineSupplyService) SetSupply(ctx context.Context, req *CreateMedicineSupplyRequest) (*MedicineSupplyResponse, error) {
	supply := &models.MedicineSupply{
		MedicineID: req.MedicineID,
		Quantity:   req.Quantity,
		ExpiredAt:  req.ExpiredAt,
	}

	createdSupply, err := s.repo.Create(ctx, supply)
	if err != nil {
		return nil, fmt.Errorf("Supply create error: %w", err)
	}

	response := &MedicineSupplyResponse{
		ID:         createdSupply.ID,
		MedicineID: createdSupply.MedicineID,
		Quantity:   createdSupply.Quantity,
		ExpiredAt:  createdSupply.ExpiredAt,
	}
	return response, nil
}

func (s *MedicineSupplyService) UpdateSupply(ctx context.Context, id uint, req *UpdateMedicineSupplyRequest) (*MedicineSupplyResponse, error) {
	supply, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Supply find error: %w", err)
	}
	if req.MedicineID != nil {
		supply.MedicineID = *req.MedicineID
	}
	if req.Quantity != nil {
		supply.Quantity = *req.Quantity
	}
	if req.ExpiredAt != nil {
		supply.ExpiredAt = *req.ExpiredAt
	}
	
	if err := s.repo.Update(ctx, supply); err != nil {
		return nil, fmt.Errorf("Supply update error: %w", err)
	}
	if err := database.MysqlDB.DB.Save(&supply).Error; err != nil {
		return nil, fmt.Errorf("Supply save error: %w", err)
	}

	response := &MedicineSupplyResponse{
		ID:         supply.ID,
		MedicineID: supply.MedicineID,
		Quantity:   supply.Quantity,
		ExpiredAt:  supply.ExpiredAt,
	}
	return response, nil
}

func (s *MedicineSupplyService) DeleteSupply(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("Supply delete error: %w", err)
	}
	return nil
}

func (s *MedicineSupplyService) RestoreSupply(ctx context.Context, id uint) error {
	if err := s.repo.Restore(ctx, id); err != nil {
		return fmt.Errorf("Supply restore error: %w", err)
	}
	return nil
}

func (s *MedicineSupplyService) ForceDeleteSupply(ctx context.Context, id uint) error {
	if err := s.repo.ForceDelete(ctx, id); err != nil {
		return fmt.Errorf("Supply force delete error: %w", err)
	}
	return nil
}

func (s *MedicineSupplyService) MassDeleteSupply(ctx context.Context, ids []uint) error {
	if err := s.repo.MassDelete(ctx, ids); err != nil {
		return fmt.Errorf("Supply mass delete error: %w", err)
	}
	return nil
}

func (s *MedicineSupplyService) ReleaseSupply(ctx context.Context, medicineID uint) (*MedicineSupplyResponse, error) {
	supply, err := s.repo.GetOlderExpiringSupplyByMedicineID(ctx, medicineID)
	if err != nil {
		return nil, fmt.Errorf("Supply find error: %w", err)
	}

	if err := supply.Release(); err != nil {
		return nil, fmt.Errorf("Supply release error: %w", err)
	}

	if err := s.repo.Update(ctx, supply); err != nil {
		return nil, fmt.Errorf("Supply update error: %w", err)
	}
	
	return &MedicineSupplyResponse{
		ID:         supply.ID,
		MedicineID: supply.MedicineID,
		Quantity:   supply.Quantity,
		ExpiredAt:  supply.ExpiredAt,
	}, nil
}
