package controllers

import (
	"auto-pharmacy/internal/services"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

func MedicineIndex(c *echo.Context) error {
	medicines, err := services.GetAllMedicines()
	if err != nil {
		c.Logger().Error("Medicines error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicines error "+err.Error())
	}

	return c.JSON(http.StatusOK, medicines)
}

func MedicineGet(c *echo.Context) error {
	medicineId := c.Param("medicine")
	medicine, err := services.GetMedicine(medicineId)
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

func MedicineSet(c *echo.Context) error {
	body := json.NewDecoder(c.Request().Body)
	med, err := services.SetMedicine(body)
	if err != nil {
		c.Logger().Error("Medicine error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine error "+err.Error())
	}
	
	return c.JSON(http.StatusCreated, med)
}

func MedicineUpdate(c *echo.Context) error {
	medicineId := c.Param("medicine")
	body := json.NewDecoder(c.Request().Body)
	med, err := services.UpdateMedicine(medicineId, body)
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

func MedicineDelete(c *echo.Context) error {
	medicineId := c.Param("medicine")
	err := services.DeleteMedicine(medicineId)
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

func MedicineRestore(c *echo.Context) error {
	medicineId := c.Param("medicine")
	err := services.RestoreMedicine(medicineId)
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

func MedicineForceDelete(c *echo.Context) error {
	medicineId := c.Param("medicine")
	err := services.ForceDeleteMedicine(medicineId)
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

func MedicineMassDelete(c *echo.Context) error {
	var data map[string][]int
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	ids, ok := data["ids"]
	if !ok {
		c.Logger().Error("Data find error ")
		return echo.NewHTTPError(http.StatusInternalServerError, "Data find error ")
	}
	err := services.MassDeleteMedicine(ids)
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

func MedicineRelease(c *echo.Context) error {
	medicineId := c.Param("medicine")
	medicine, err := services.ReleaseSupply(medicineId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Medicine not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, "Medicine not found")
	// }
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, medicine)
}

func MedicineAssociateTag(c *echo.Context) error {
	medicineId := c.Param("medicine")
	tagId := c.Param("tag")
	medicine, err := services.AssociateTagToMedicine(medicineId, tagId)
	if err != nil {
		c.Logger().Error("Medicine associate error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Medicine associate error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, medicine)
}
