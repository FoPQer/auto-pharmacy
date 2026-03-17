package controllers

import (
	"auto-pharmacy/internal/services"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

type MedicineController struct {
	medicineService *services.MedicineService
	supplyService  *services.MedicineSupplyService
}

func (mc *MedicineController) MedicineIndex(c *echo.Context) error {
	medicines, err := mc.medicineService.GetAllMedicines(c.Request().Context())
	if err != nil {
		c.Logger().Error("Medicines error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicines error "+err.Error())
	}

	return c.JSON(http.StatusOK, medicines)
}

func (mc *MedicineController) MedicineGet(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	medicine, err := mc.medicineService.GetMedicine(c.Request().Context(), medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	return c.JSON(http.StatusOK, medicine)
}

func (mc *MedicineController) MedicineSet(c *echo.Context) error {
	req := new(services.CreateMedicineRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	med, err := mc.medicineService.SetMedicine(c.Request().Context(), req)
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	
	return c.JSON(http.StatusCreated, med)
}

func (mc *MedicineController) MedicineUpdate(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	req := new(services.UpdateMedicineRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	med, err := mc.medicineService.UpdateMedicine(c.Request().Context(), medicineId, req)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, med)
}

func (mc *MedicineController) MedicineDelete(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	err = mc.medicineService.DeleteMedicine(c.Request().Context(), medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (mc *MedicineController) MedicineRestore(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	err = mc.medicineService.RestoreMedicine(c.Request().Context(), medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (mc *MedicineController) MedicineForceDelete(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	err = mc.medicineService.ForceDeleteMedicine(c.Request().Context(), medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (mc *MedicineController) MedicineMassDelete(c *echo.Context) error {
	var data map[string][]uint
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	ids, ok := data["ids"]
	if !ok {
		c.Logger().Error("Data find error ")
		return echo.NewHTTPError(http.StatusInternalServerError, "Data find error ")
	}
	err := mc.medicineService.MassDeleteMedicine(c.Request().Context(), ids)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (mc *MedicineController) MedicineRelease(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	_, err = mc.supplyService.ReleaseSupply(c.Request().Context(), medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}

	medicine, err := mc.medicineService.GetMedicine(c.Request().Context(), medicineId)
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, medicine)
}

func (mc *MedicineController) MedicineAssociateTag(c *echo.Context) error {
	medicineId, err := echo.PathParam[uint](c, "medicine")
	if err != nil {
		c.Logger().Error("Invalid medicine ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid medicine ID "+err.Error())
	}
	tagId, err := echo.PathParam[uint](c, "tag")
	if err != nil {
		c.Logger().Error("Invalid tag ID " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag ID "+err.Error())
	}
	
	if err := mc.medicineService.AssociateTagToMedicine(c.Request().Context(), medicineId, tagId); err != nil {
		c.Logger().Error("Medicine associate error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine associate error "+err.Error())
	}
	
	return c.NoContent(http.StatusNoContent)
}
