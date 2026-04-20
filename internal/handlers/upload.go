package handlers

import "net/http"

// UploadMap godoc
// @Summary      Загрузить новый .mbtiles файл
// @Description  Загружает файл карты, сохраняет его на сервере и добавляет запись в базу данных.
// @Tags         maps
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Файл карты в формате .mbtiles"
// @Success      200  {object}  map[string]string "Файл успешно загружен"
// @Failure      400  {object}  map[string]string "Ошибка в запросе"
// @Failure      500  {object}  map[string]string "Внутренняя ошибка сервера"
// @Router       /upload [post]
func UploadMap(w http.ResponseWriter, r *http.Request) {
    
}
