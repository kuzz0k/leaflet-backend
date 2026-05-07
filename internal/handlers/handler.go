package handlers

import (
	"leaflet-back/internal/services"
	"net/http"
)

type Handler struct {
	mbtilesService *services.MbtilesService
}

func NewHandler(mbtilesService *services.MbtilesService) *Handler {
	return &Handler{
		mbtilesService: mbtilesService,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tiles/", h.GetTile)
	return mux
}