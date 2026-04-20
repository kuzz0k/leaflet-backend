package models

// Map представляет информацию о карте, хранящуюся в базе данных.
type Map struct {
    ID          int    `json:"id" db:"id"`
    Name        string `json:"name" db:"name"`
    FilePath    string `json:"file_path" db:"file_path"`
    Description string `json:"description" db:"description"`
    // Дополнительные поля, например, дата загрузки, размер и т.д.
}
