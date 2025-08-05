package service

import (
	"fmt"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/parser"
)

func AggregateService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	if len(flags.Aggregate) == 0 {
		return records, nil
	}

	results, err := parser.Aggregate(records, flags.Aggregate)
	if err != nil {
		return nil, fmt.Errorf("агрегация: %w", err)
	}

	aggRecord := model.Record{}
	for k, v := range results {
		aggRecord[k] = fmt.Sprintf("%.2f", v)
	}

	return []model.Record{aggRecord}, nil
}

