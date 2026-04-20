package handlers

import "net/http"

// GetMapInfo godoc
// @Summary      Получить информацию о доступных картах
// @Description  Возвращает список всех карт, информация о которых хранится в базе данных.
// @Tags         maps
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Map "Успешный ответ"
// @Router       /maps/info [get]
func GetMapInfo(w http.ResponseWriter, r *http.Request) {
    
}
