package services

import (
	"auto-pharmacy/internal/models"
	"auto-pharmacy/internal/repository"
	"context"
	"encoding/json"
	"fmt"
)

type MedicineResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Measurement *string `json:"measurement"`
	Dose float64 `json:"dose"`
	Measurement_dose *string `json:"measurement_dose"`
	Box *string `json:"box"`
	Place *string `json:"place"`
}

type GetMedicineRequest struct {
	ID uint `json:"id" param:"medicine" validate:"required"`
}

type CreateMedicineRequest struct {
	Name     string `json:"name" validate:"min=2,max=255"`
	Measurement *string `json:"measurement" validate:"min=1,max=255"`
	Dose float64 `json:"dose" validate:"min=1"`
	Measurement_dose *string `json:"measurement_dose" validate:"min=1,max=255"`
	Box *string `json:"box" validate:"min=1"`
	Place *string `json:"place" validate:"min=1,max=255"`
}

type UpdateMedicineRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Measurement *string `json:"measurement,omitempty" validate:"omitempty,min=1,max=255"`
	Dose *float64 `json:"dose,omitempty" validate:"omitempty,min=1"`
	Measurement_dose *string `json:"measurement_dose,omitempty" validate:"omitempty,min=1,max=255"`
	Box *string `json:"box,omitempty" validate:"omitempty,min=1"`
	Place *string `json:"place,omitempty" validate:"omitempty,min=1,max=255"`
}

type MedicineService struct {
    repo            repository.MedicineRepository
    medicineTagRepo repository.MedicineTagRepository
}

func NewMedicineService(repo repository.MedicineRepository) *MedicineService {
	return &MedicineService{repo: repo}
}

func (s *MedicineService) GetAllMedicines(ctx context.Context) ([]MedicineResponse, error) {
	medicines, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Medicine GetAll error: %w", err)
	}
	var response = make([]MedicineResponse, 0)
	for _, medicine := range medicines {
		response = append(response, MedicineResponse{
			ID:   medicine.ID,
			Name: medicine.Name,
			Measurement: medicine.Measurement,
			Dose: medicine.Dose,
			Measurement_dose: medicine.Measurement_dose,
			Box: medicine.Box,
			Place: medicine.Place,
		})
	}
	return response, nil
}

func (s *MedicineService) GetMedicine(ctx context.Context, id uint) (*MedicineResponse, error) {
	medicine, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Medicine GetByID error: %w", err)
	}

	response := &MedicineResponse{
		ID:   medicine.ID,
		Name: medicine.Name,
		Measurement: medicine.Measurement,
		Dose: medicine.Dose,
		Measurement_dose: medicine.Measurement_dose,
		Box: medicine.Box,
		Place: medicine.Place,
	}

	return response, nil
}

func (s *MedicineService) SetMedicine(ctx context.Context, req *CreateMedicineRequest) (*MedicineResponse, error) {
	medicine := models.NewMedicine(req.Name, req.Measurement, req.Dose, req.Measurement_dose, req.Box, req.Place)
	createdMedicine, err := s.repo.Create(ctx, medicine)
	if err != nil {
		return nil, fmt.Errorf("Medicine create error: %w", err)
	}

	return &MedicineResponse{
		ID:   createdMedicine.ID,
		Name: createdMedicine.Name,
		Measurement: createdMedicine.Measurement,
		Dose: createdMedicine.Dose,
		Measurement_dose: createdMedicine.Measurement_dose,
		Box: createdMedicine.Box,
		Place: createdMedicine.Place,
	}, nil
}

func (s *MedicineService) UpdateMedicine(ctx context.Context, id uint, body *json.Decoder) (*MedicineResponse, error) {
	medicine, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := body.Decode(&medicine); err != nil {
		return nil, fmt.Errorf("Medicine decode error: %w", err)
	}
	if err := s.repo.Update(ctx, medicine); err != nil {
		return nil, fmt.Errorf("Medicine save error: %w", err)
	}

	return &MedicineResponse{
		ID:   medicine.ID,
		Name: medicine.Name,
		Measurement: medicine.Measurement,
		Dose: medicine.Dose,
		Measurement_dose: medicine.Measurement_dose,
		Box: medicine.Box,
		Place: medicine.Place,
	}, nil
}

func (s *MedicineService) DeleteMedicine(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("Medicine delete error: %w", err)
	}
	return nil
}

func (s *MedicineService) RestoreMedicine(ctx context.Context, id uint) error {
	if err := s.repo.Restore(ctx, id); err != nil {
		return fmt.Errorf("Medicine restore error: %w", err)
	}
	return nil
}

func (s *MedicineService) ForceDeleteMedicine(ctx context.Context, id uint) error {
	if err := s.repo.ForceDelete(ctx, id); err != nil {
		return fmt.Errorf("Medicine force delete error: %w", err)
	}
	return nil
}

func (s *MedicineService) MassDeleteMedicine(ctx context.Context, ids []uint) error {
	if err := s.repo.MassDelete(ctx, ids); err != nil {
		return fmt.Errorf("Medicine mass delete error: %w", err)
	}
	return nil
}

func (s *MedicineService) AssociateTagToMedicine(ctx context.Context, medicineID uint, tagID uint) error {
    if err := s.medicineTagRepo.Associate(ctx, medicineID, tagID); err != nil {
        return fmt.Errorf("Associate tag to medicine error: %w", err)
    }
    return nil
}

func (s *MedicineService) DissociateTagFromMedicine(ctx context.Context, medicineID uint, tagID uint) error {
	if err := s.medicineTagRepo.Dissociate(ctx, medicineID, tagID); err != nil {
		return fmt.Errorf("Dissociate tag from medicine error: %w", err)
	}
	return nil
}
