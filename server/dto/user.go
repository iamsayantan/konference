package dto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/iamsayantan/konference"
)

type RegistrationRequest struct {
	FirstName string `json:"first_name" validate:"required,max=191"`
	LastName  string `json:"last_name" validate:"required,max=191"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

func (r *RegistrationRequest) Validate() map[string]string {
	err := validate.Struct(r)
	if err == nil {
		return make(map[string]string)
	}

	validationErrors := err.(validator.ValidationErrors)
	return translateErrors(validationErrors)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (lr *LoginRequest) Validate() map[string]string {
	err := validate.Struct(lr)
	if err == nil {
		return make(map[string]string)
	}

	validationErrors := err.(validator.ValidationErrors)
	return translateErrors(validationErrors)
}

type LoginResponse struct {
	User        konference.User `json:"user"`
	AccessToken string          `json:"access_token"`
}

// UserClaims struct holds the data that would be encoded to a jwt.
type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
