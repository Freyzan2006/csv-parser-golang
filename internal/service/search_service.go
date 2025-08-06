package service

import (
	"fmt"
	"strings"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
)

func SearchService(cfg *config.Config, records []model.Record) ([]model.Record, error) {
	if cfg.Search == "" {
		return records, nil
	}

	parts := strings.SplitN(cfg.Search, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("неверный формат поиска: %s (ожидается key=value)", cfg.Search)
	}

	field := parts[0]
	value := parts[1]

	var result []model.Record

	for _, rec := range records {
		if v, ok := rec[field]; ok && v == value {
			result = append(result, rec)
		}
	}

	if cfg.Verbose && len(result) == 0 {
		fmt.Printf("⚠️ Поиск не дал результатов\n")
	}

	return result, nil
}

