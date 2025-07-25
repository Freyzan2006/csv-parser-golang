package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/validator"
)


func SortService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.Sort) > 0 {
		updatedRecords, err := validator.Sort(records, flags.Sort, flags.Verbose)
		if err != nil {
			return nil, fmt.Errorf("сортировка: %w", err)
		}
		return updatedRecords, nil
	}
	return records, nil
}