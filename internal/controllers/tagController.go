package controllers

import (
	"auto-pharmacy/internal/services"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

func TagIndex(c *echo.Context) error {
	tags, err := services.GetAllTags()
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}

	return c.JSON(http.StatusOK, tags)
}

func TagGet(c *echo.Context) error {
	tagId := c.Param("tag")
	tag, err := services.GetTag(tagId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, tag)
}

func TagSet(c *echo.Context) error {
	body := json.NewDecoder(c.Request().Body)
	tag, err := services.SetTag(body)
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	return c.JSON(http.StatusCreated, tag)
}

func TagUpdate(c *echo.Context) error {
	tagId := c.Param("tag")
	body := json.NewDecoder(c.Request().Body)
	tag, err := services.UpdateTag(tagId, body)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	c.Logger().Error("Tag not found " + err.Error())
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	return c.JSON(http.StatusOK, tag)
}

func TagDelete(c *echo.Context) error {
	tagId := c.Param("tag")
	err := services.DeleteTag(tagId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func TagRestore(c *echo.Context) error {
	tagId := c.Param("tag")
	err := services.RestoreTag(tagId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func TagForceDelete(c *echo.Context) error {
	tagId := c.Param("tag")
	err := services.ForceDeleteTag(tagId)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }
	if err != nil {
		c.Logger().Error("Tag error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Tag error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
