package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/validator"
)


func RequiredService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.Required) > 0 {
		updatedRecords, err := validator.Required(records, flags.Required)
		if err != nil {
			return nil, fmt.Errorf("валидация required: %w", err)
		}
		return updatedRecords, nil
	}
	return records, nil
}