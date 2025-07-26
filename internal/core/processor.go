package core

import (
	"csv-parser/internal/config"
	"csv-parser/internal/parser"
	"csv-parser/internal/printer"
	"csv-parser/internal/service"
	"fmt"
)

func Process(cfg *config.Config) error {
	records, err := parser.ReadCSV(cfg.FilePath, cfg.Header)
	if err != nil {
		return fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	records, err = service.RequiredService(cfg, records)
	if err != nil {
		return fmt.Errorf("валидация required: %w", err)
	}
	

	records, err = service.RangeValidatorService(cfg, records)
	if err != nil {
		return fmt.Errorf("валидация range: %w", err)
	}


	records, err = service.ValidateTypeService(cfg, records)
	if err != nil {
		return fmt.Errorf("валидация типов: %w", err)
	}
	

	records, err = service.FilterService(cfg, records)
	if err != nil {
		return fmt.Errorf("фильтрация: %w", err)
	}


	records, err = service.SortService(cfg, records)
	if err != nil {
		return fmt.Errorf("сортировка: %w", err)
	}
	

	printer.PrintRecords(records)
	return nil
}
