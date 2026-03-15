package controllers

import (
	"auto-pharmacy/internal/database"
	"auto-pharmacy/internal/models"
	"auto-pharmacy/internal/services"
	"crypto/subtle"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type LoginController struct {
	userService *services.UserService
}

func NewLoginController(userService *services.UserService) *LoginController {
	return &LoginController{userService: userService}
}

func (lc *LoginController) LoginHandler(c *echo.Context) error {
	email := c.FormValue("email")
	tokenLifetime, err := strconv.ParseInt(os.Getenv("TOKEN_LIFETIME"), 10, 32)
	expirationTime := time.Now().Add(time.Duration(tokenLifetime) * time.Minute)
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "auto-pharmacy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.Logger().Error("Token generation error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Token generation error "+err.Error())
	}

	// Return the token to the client
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (lc *LoginController) RegisterHandler(c *echo.Context) error {
	var req services.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("User request parse error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User request parse error "+err.Error())
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("User request validation error " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error "+err.Error())
	}

	user, err := lc.userService.SetUser(c.Request().Context(), &req)
	if err != nil {
		c.Logger().Error("User creation error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "User creation error "+err.Error())
	}

	tokenLifetime, err := strconv.ParseInt(os.Getenv("TOKEN_LIFETIME"), 10, 32)
	expirationTime := time.Now().Add(time.Duration(tokenLifetime) * time.Minute)
	claims := &models.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "auto-pharmacy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.Logger().Error("Token generation error " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Token generation error "+err.Error())
	}

	// Return the token to the client
	return c.JSON(http.StatusCreated, map[string]any{
		"token": tokenString,
		"user": user,
	})
}

func BasicAuth(c *echo.Context, email, password string) (bool, error) {
	var dbUser models.User
	err := database.MysqlDB.DB.Find(&dbUser, "email = ?", email).Error
	// if (errors.Is(err, gorm.ErrRecordNotFound)) || dbUser.ID == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("Email or Password doesn't match"))
	// 	return
	// }
	if err != nil {
		c.Logger().Error("User Find error " + err.Error())
		return false, errors.Join(errors.New("User Find error"), err)
	}
	if subtle.ConstantTimeCompare([]byte(email), []byte("joe")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
		return true, nil
	}
	return false, nil
}
