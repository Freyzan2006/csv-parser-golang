package service

import "fmt"

import (
	"csv-parser/internal/config"
	"csv-parser/internal/exporter"
	"csv-parser/internal/model"
)


func ExportService(flags *config.Config, records []model.Record) ([]model.Record, error) {
	
	ext := map[string]string{
		"csv":   "csv",
		"json":  "json",
		"excel": "xlsx",
	}[flags.Export]

	outputFile := fmt.Sprintf("output.%s", ext)


	if len(flags.Export) > 0 {
		if err := exporter.Export(records, flags.Export, outputFile); err != nil {
			return nil, fmt.Errorf("ошибка экспорта: %w", err)
		}
		fmt.Println("Результаты экспортированы в:", outputFile)
	}
	return records, nil
}