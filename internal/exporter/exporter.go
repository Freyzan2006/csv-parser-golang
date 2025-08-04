package exporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"csv-parser/internal/model"
)

import "github.com/xuri/excelize/v2"

func Export(records []model.Record, format, outputPath string) error {
	switch format {
	case "csv":
		return exportCSV(records, outputPath)
	case "json":
		return exportJSON(records, outputPath)
	case "excel":
		return exportExcel(records, outputPath) // Реализуем ниже
	default:
		return fmt.Errorf("неподдерживаемый формат вывода: %s", format)
	}
}

func exportCSV(records []model.Record, path string) error {
	if len(records) == 0 {
		return nil
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Пишем заголовки
	var headers []string
	for k := range records[0] {
		headers = append(headers, k)
	}
	writer.Write(headers)

	// Пишем строки
	for _, rec := range records {
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = rec[h]
		}
		writer.Write(row)
	}
	return nil
}

func exportJSON(records []model.Record, path string) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func exportExcel(records []model.Record, path string) error {
	if len(records) == 0 {
		return nil
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{}
	for k := range records[0] {
		headers = append(headers, k)
	}

	// Заголовки
	for i, h := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheet, cell, h)
	}

	// Данные
	for rowIdx, rec := range records {
		for colIdx, h := range headers {
			cell := fmt.Sprintf("%s%d", string(rune('A'+colIdx)), rowIdx+2)
			f.SetCellValue(sheet, cell, rec[h])
		}
	}

	if err := f.SaveAs(path); err != nil {
		return err
	}
	return nil
}