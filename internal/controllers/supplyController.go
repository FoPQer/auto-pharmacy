package controllers

import (
	"auto-pharmacy/internal/services"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

func SupplyIndex(c *echo.Context) error {
	supplies, err := services.GetAllSupplies()
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}

	return c.JSON(http.StatusOK, supplies)
}

func SupplyGet(c *echo.Context) error {
	supplyId := c.Param("supply")
	supply, err := services.GetSupply(supplyId)
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}

	return c.JSON(http.StatusOK, supply)
}

func SupplySet(c *echo.Context) error {
	body := json.NewDecoder(c.Request().Body)
	sup, err := services.SetSupply(body)
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.JSON(http.StatusCreated, sup)
}

func SupplyUpdate(c *echo.Context) error {
	supplyId := c.Param("supply")
	body := json.NewDecoder(c.Request().Body)
	sup, err := services.UpdateSupply(supplyId, body)
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.JSON(http.StatusOK, sup)
}

func SupplyDelete(c *echo.Context) error {
	supplyId := c.Param("supply")
	if err := services.DeleteSupply(supplyId); err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func SupplyRestore(c *echo.Context) error {
	supplyId := c.Param("supply")
	if err := services.RestoreSupply(supplyId); err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func SupplyForceDelete(c *echo.Context) error {
	supplyId := c.Param("supply")
	if err := services.ForceDeleteSupply(supplyId); err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func SupplyMassDelete(c *echo.Context) error {
	var data map[string][]int
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	ids, ok := data["ids"]
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Data find error ")
	}
	err := services.MassDeleteSupply(ids)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Supply error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Supply error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
