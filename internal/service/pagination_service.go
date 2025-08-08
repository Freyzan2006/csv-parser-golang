package service

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"fmt"
)

// func PaginationService(cfg *config.Config, records []model.Record) []model.Record {
// 	totalRecords := len(records) 
// 	fmt.Printf("Страница %d из %d (всего %d записей)\n", cfg.Page, (totalRecords+cfg.PerPage-1)/cfg.PerPage, totalRecords)

// 	if cfg.PerPage <= 0 {
// 		return records // защита от деления на ноль
// 	}
// 	if cfg.Page <= 0 {
// 		cfg.Page = 1
// 	}

// 	start := (cfg.Page - 1) * cfg.PerPage
// 	if start >= len(records) {
// 		return []model.Record{} // страница за пределами массива
// 	}

// 	end := start + cfg.PerPage
// 	if end > len(records) {
// 		end = len(records)
// 	}

	

// 	return records[start:end]
// }

func PaginationService(cfg *config.Config, records []model.Record) []model.Record {
	totalRecords := len(records)
	if cfg.PerPage <= 0 {
		return records // защита от деления на ноль
	}

	// считаем общее кол-во страниц (округляем вверх)
	totalPages := (totalRecords + cfg.PerPage - 1) / cfg.PerPage

	// защита от некорректных страниц
	if cfg.Page < 1 {
		cfg.Page = 1
	}
	if cfg.Page > totalPages {
		fmt.Printf("Запрошенная страница %d превышает максимальное количество страниц (%d). Показана последняя страница.\n", cfg.Page, totalPages)
		cfg.Page = totalPages
	}

	fmt.Printf("Страница %d из %d (всего %d записей)\n", cfg.Page, totalPages, totalRecords)

	start := (cfg.Page - 1) * cfg.PerPage
	end := start + cfg.PerPage
	if end > totalRecords {
		end = totalRecords
	}

	return records[start:end]
}

