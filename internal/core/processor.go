package core

import (
	"csv-parser/internal/config"
	"csv-parser/internal/parser"
	"csv-parser/internal/printer"
	"csv-parser/internal/service"
	"fmt"
)

func Process(cfg *config.Config) error {
	records, err := parser.ReadCSV(cfg.FilePath, cfg.Header, cfg.Encoding)
	if err != nil {
		return fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	records, err = service.CompareService(cfg, records)
	if err != nil {
		return fmt.Errorf("сравнение: %w", err)
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


	records, err = service.SearchService(cfg, records)
	if err != nil {
		return fmt.Errorf("поиск: %w", err)
	}


	records, err = service.AggregateService(cfg, records)
	if err != nil {
		return fmt.Errorf("агрегация: %w", err)
	}


	records, err = service.ExportService(cfg, records)
	if err != nil {
		return fmt.Errorf("экспорт: %w", err)
	}
	
	
	records = service.PaginationService(cfg, records)
	

	printer.PrintRecords(records)
	return nil
}
