package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/validator"
)


func FilterService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.Filter) > 0 {
		updatedRecords, err := validator.Filter(records, flags.Filter, flags.Verbose)
		if err != nil {
			return nil, fmt.Errorf("фильтрация: %w", err)
		}
		return updatedRecords, nil
	}
	return records, nil
}