package handler

import (
	"github.com/go-chi/chi/v5"
	"main.go/internal/service"
	"net/http"
)

type Handler struct {
	router  chi.Router
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {

	h := &Handler{
		router:  chi.NewRouter(),
		service: service,
	}
	h.InitRoutes()
	return h
}

func (h *Handler) InitRoutes() {
	h.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	h.router.Route("/auth", func(r chi.Router) {
		r.Get("/sign-up", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("sign-up"))
		})
		r.Get("/sign-in", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("sign-in"))
		})
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
