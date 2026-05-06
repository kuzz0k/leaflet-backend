package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

// GetTile godoc
// @Summary      Получить тайл карты
// @Description  Возвращает один тайл карты по указанным координатам и названию карты.
// @Tags         tiles
// @Produce      image/png
// @Param        map_name path      string  true  "Название карты"
// @Param        z        path      int     true  "Уровень масштабирования (Z)"
// @Param        x        path      int     true  "Координата X"
// @Param        y        path      int     true  "Координата Y"
// @Success      200      {file}    string  "Бинарные данные тайла"
// @Failure      404      {object}  map[string]string "Тайл не найден"
// @Failure      500      {object}  map[string]string "Внутренняя ошибка сервера"
// @Router       /tiles/{map_name}/{z}/{x}/{y}.png [get]
func (h *Handler) GetTile(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 5 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	mapName := parts[1]
	z, errZ := strconv.Atoi(parts[2])
	x, errX := strconv.Atoi(parts[3])

	yStr := strings.TrimSuffix(parts[4], ".png")
	y, errY := strconv.Atoi(yStr)

	if errZ != nil || errX != nil || errY != nil {
		http.Error(w, "Invalid coordinates", http.StatusBadRequest)
		return
	}

	tileData, err := h.mbtilesService.GetTileData(mapName, z, x, y)
	if err != nil {
		slog.Error("Failed to get tile data", "error", err, "map", mapName, "z", z, "x", x, "y", y)
		http.Error(w, "Tile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(tileData)
}
