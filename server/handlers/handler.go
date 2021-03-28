package handlers

import "github.com/go-chi/chi"

type Handler interface {
	Routes() chi.Router
}
