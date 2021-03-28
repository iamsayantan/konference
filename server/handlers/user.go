package handlers

import (
	"github.com/go-chi/chi"
	"github.com/iamsayantan/konference"
	"net/http"
)

type userHandler struct {
	service konference.UserService
}

func (u *userHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", u.register)

	return r
}

func (u *userHandler) register(w http.ResponseWriter, r *http.Request) {

}

func NewUserHandler(us konference.UserService) Handler {
	return &userHandler{service: us}
}
