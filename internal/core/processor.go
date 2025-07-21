// internal/core/processor.go
package core

import (
	"fmt"
	"csv-parser/internal/config"
	"csv-parser/internal/parser"
	"csv-parser/internal/validator"
	"csv-parser/internal/printer"
)

func Process(cfg *config.Config) error {
	records, err := parser.ReadCSV(cfg.FilePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	if len(cfg.Required) > 0 {
		records, err = validator.Required(records, cfg.Required)
		if err != nil {
			return fmt.Errorf("валидация required: %w", err)
		}
	}

	if len(cfg.Range) > 0 {
		records, err = validator.RangeValidator(records, cfg.Range, cfg.Verbose)
		if err != nil {
			return fmt.Errorf("валидация range: %w", err)
		}
	}

	
	if len(cfg.ValidateType) > 0 {
		records, err = validator.ValidateTypes(records, cfg.ValidateType, cfg.Verbose)
		if err != nil {
			return fmt.Errorf("валидация типов: %w", err)
		}
	}


	printer.PrintRecords(records)
	return nil
}
