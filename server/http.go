package server

import (
	"github.com/go-chi/chi"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/iamsayantan/konference"
	"github.com/iamsayantan/konference/server/handlers"
	"github.com/iamsayantan/konference/server/rendering"
	"net/http"
)

type Server struct {
	UserService konference.UserService

	router chi.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(us konference.UserService) *Server {
	s := &Server{
		UserService: us,
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r := chi.NewRouter()

	// mount the middlewares
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(corsHandler.Handler)
	r.Use(chimw.AllowContentType("application/json"))

	r.Route("/users", func(r chi.Router) {
		uh := handlers.NewUserHandler(us)
		r.Mount("/v1", uh.Routes())
	})

	r.Get("/", HelloHandler)
	r.NotFound(notFoundHandler)
	r.MethodNotAllowed(methodNotAllowedHandler)

	s.router = r
	return s
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Author string `json:"author"`
		Email  string `json:"email"`
	}{Author: "Sayantan Das", Email: "das.sayantan94@gmail.com"}

	rendering.RenderSuccessWithData(w, r, "Welcome to Konference API server", http.StatusOK, resp)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	rendering.RenderError(w, r, "the requested route does not exist in our server", "route.not_found", http.StatusNotFound)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	rendering.RenderError(w, r, "the requested method is not supported for the given route", "route.method_not_allowed", http.StatusMethodNotAllowed)
}
