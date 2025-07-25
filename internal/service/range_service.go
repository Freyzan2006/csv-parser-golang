package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/validator"
)



func RangeValidatorService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.Range) > 0 {
		updatedRecords, err := validator.RangeValidator(records, flags.Range, flags.Verbose)
		if err != nil {
			return nil, fmt.Errorf("валидация range: %w", err)
		}
		return updatedRecords, nil
	}
	return records, nil
}
