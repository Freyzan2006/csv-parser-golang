package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/validator"
)


func ValidateTypeService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.ValidateType) > 0 {
		updatedRecords, err := validator.ValidateTypes(records, flags.ValidateType, flags.Verbose)
		if err != nil {
			return nil, fmt.Errorf("валидация типов: %w", err)
		}
		return updatedRecords, nil
	}
	return records, nil
}