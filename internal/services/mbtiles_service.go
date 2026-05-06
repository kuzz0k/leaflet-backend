package services

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type MbtilesService struct {
	storagePath string
}

func NewMbtilesService(storagePath string) *MbtilesService {
	return &MbtilesService{storagePath: storagePath}
}

func (s *MbtilesService) GetTileData(mapName string, z, x, y int) ([]byte, error) {
	dbPath := filepath.Join(s.storagePath, fmt.Sprintf("%s.mbtiles", mapName))

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open mbtiles file: %w", err)
	}
	defer db.Close()

	y = (1 << z) - 1 - y

	var tileData []byte
	query := "SELECT tile_data FROM tiles WHERE zoom_level = ? AND tile_column = ? AND tile_row = ?"
	err = db.QueryRow(query, z, x, y).Scan(&tileData)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tile not found for z=%d, x=%d, y=%d", z, x, y)
		}
		return nil, fmt.Errorf("failed to query tile: %w", err)
	}

	return tileData, nil
}