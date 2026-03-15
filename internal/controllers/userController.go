package controllers

import (
	"auto-pharmacy/internal/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Index(c *echo.Context) error {
	users, err := uc.service.GetAllUsers(c.Request().Context())
	if err != nil {
		c.Logger().Error("Users error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) Get(c *echo.Context) error {
	var req services.GetUserRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("User request parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User request parse error "+err.Error())
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("User request validation error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error "+err.Error())
	}

	var myErr *services.ErrUserNotFound
	user, err := uc.service.GetUser(c.Request().Context(), &req)
	if errors.As(err, &myErr) {
		c.Logger().Error("User not found " + err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Set(c *echo.Context) error {
	var req services.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("User request parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User request parse error "+err.Error())
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("User request validation error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error "+err.Error())
	}

	med, err := uc.service.SetUser(c.Request().Context(), &req)
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}

	return c.JSON(http.StatusCreated, med)
}

func (uc *UserController) Update(c *echo.Context) error {
	userId, err := echo.PathParam[uint](c, "user")
	if err != nil {
		c.Logger().Error("User ID parse error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "User ID parse error "+err.Error())
	}

	var req services.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("User request parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User request parse error "+err.Error())
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("User request validation error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error "+err.Error())
	}

	med, err := uc.service.UpdateUser(c.Request().Context(), userId, &req)
	var myErr *services.ErrUserNotFound
	if errors.As(err, &myErr) {
		c.Logger().Error("User not found " + err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}
	
	return c.JSON(http.StatusOK, med)
}

func (uc *UserController) Delete(c *echo.Context) error {
	userId, err := echo.PathParam[uint](c, "user")
	if err != nil {
		c.Logger().Error("User ID parse error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "User ID parse error "+err.Error())
	}
	err = uc.service.DeleteUser(c.Request().Context(), userId)
	var myErr *services.ErrUserNotFound
	if errors.As(err, &myErr) {
		c.Logger().Error("User not found " + err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) Restore(c *echo.Context) error {
	userId, err := echo.PathParam[uint](c, "user")
	if err != nil {
		c.Logger().Error("User ID parse error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "User ID parse error "+err.Error())
	}
	err = uc.service.RestoreUser(c.Request().Context(), userId)
	var myErr *services.ErrUserNotFound
	if errors.As(err, &myErr) {
		c.Logger().Error("User not found " + err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) ForceDelete(c *echo.Context) error {
	userId, err := echo.PathParam[uint](c, "user")
	if err != nil {
		c.Logger().Error("User ID parse error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "User ID parse error "+err.Error())
	}
	err = uc.service.ForceDeleteUser(c.Request().Context(), userId)
	var myErr *services.ErrUserNotFound
	if errors.As(err, &myErr) {
		c.Logger().Error("User not found " + err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) MassDelete(c *echo.Context) error {
	var data []uint
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		c.Logger().Error("Body parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Body parse error "+err.Error())
	}
	err := uc.service.MassDeleteUser(c.Request().Context(), data)

	if err != nil {
		c.Logger().Error("User error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User error "+err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
