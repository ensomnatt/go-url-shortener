package handler

import (
	"net/http"
	"urlshortener/internal/storage/db"
)

type Request struct {
  URL string `json:"url"`
  Alias string `json:"alias"`
}

type Response struct {
  Status string `json:"status"`
  Message string `json:"message"`
  URL string `json:"url,omitempty"`
}

type Handler struct {
  db *db.Storage
}

func New(db *db.Storage) Handler {
  return Handler{db: db}
}

func (h *Handler) CreateRoutes(r *http.ServeMux) {
  r.HandleFunc("CREATE /shorten", h.Save)
  r.HandleFunc("GET /shorten/{alias}", h.Get)
  r.HandleFunc("DELETE /shorten/{alias}", h.Delete)
}
