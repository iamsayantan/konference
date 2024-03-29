package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/iamsayantan/konference/config"
	"github.com/iamsayantan/konference/server/middlewares"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/iamsayantan/konference"
	"github.com/iamsayantan/konference/server/dto"
	"github.com/iamsayantan/konference/server/rendering"
)

type userHandler struct {
	service konference.UserService
}

func (u *userHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", u.login)
	r.Post("/register", u.register)

	// Routes that require authentication is grouped here.
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthChecker)
		r.Get("/me", u.me)
	})
	return r
}

func (u *userHandler) register(w http.ResponseWriter, r *http.Request) {
	var registrationDto dto.RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&registrationDto)
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.register.decode_request", http.StatusInternalServerError)
		return
	}

	validationErrors := registrationDto.Validate()
	if len(validationErrors) != 0 {
		rendering.RenderErrorsWithData(w, r, "validation errors", "user.register.validation_error", http.StatusUnprocessableEntity, validationErrors)
		return
	}

	err = u.service.CreateUser(r.Context(), registrationDto.Email, registrationDto.FirstName, registrationDto.LastName, registrationDto.Password)
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.register.creation_error", http.StatusBadRequest)
		return
	}

	rendering.RenderSuccess(w, r, "Registration Successful", http.StatusCreated)
}

func (u *userHandler) login(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.login.decode_request", http.StatusInternalServerError)
		return
	}

	validationErrors := loginDto.Validate()
	if len(validationErrors) != 0 {
		rendering.RenderErrorsWithData(w, r, "validation errors", "user.login.validation_error", http.StatusUnprocessableEntity, validationErrors)
		return
	}

	user, err := u.service.Authenticate(r.Context(), loginDto.Email, loginDto.Password)
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.login.authentication_error", http.StatusBadRequest)
		return
	}

	jwtExpirationTime := time.Now().Add(time.Hour * 24)
	jwtClaims := dto.UserClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwtExpirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := jwtToken.SignedString([]byte(config.AppSecret))
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.login.jwt_generation", http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "konference-auth",
		Value:    accessToken,
		Path:     "/",
		Expires:  jwtExpirationTime,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &sessionCookie)
	resp := dto.LoginResponse{User: *user, AccessToken: accessToken}
	rendering.RenderSuccessWithData(w, r, "login successful", http.StatusOK, resp)
}

func (u *userHandler) me(w http.ResponseWriter, r *http.Request) {
	loggedInUserID := middlewares.GetAuthenticatedUserIdFromContext(r.Context())
	user, err := u.service.GetUserDetails(r.Context(), loggedInUserID)

	if err != nil {
		rendering.RenderError(w, r, err.Error(), "user.me.user_details", http.StatusBadRequest)
		return
	}

	resp := struct {
		User konference.User `json:"user"`
	}{User: *user}

	rendering.RenderSuccessWithData(w, r, "success", http.StatusOK, resp)
}

func NewUserHandler(us konference.UserService) Handler {
	return &userHandler{service: us}
}
