// internal/parser/reader.go
package parser

import (
    "encoding/csv"
    "fmt"
    "os"
)

import (
    "csv-parser/internal/model"
)

func ReadCSV(path string) ([]model.Record, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("не удалось открыть файл: %w", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("ошибка при чтении CSV: %w", err)
    }

    if len(records) < 1 {
        return nil, fmt.Errorf("пустой файл или нет заголовков")
    }

    headers := records[0]
    var result []model.Record

    for _, row := range records[1:] {
        entry := make(model.Record)
        for i, cell := range row {
            if i < len(headers) {
                entry[headers[i]] = cell
            }
        }
        result = append(result, entry)
    }


    return result, nil
}
