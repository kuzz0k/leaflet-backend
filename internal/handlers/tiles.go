package handlers

import "net/http"

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
func GetTile(w http.ResponseWriter, r *http.Request) {
    
}
