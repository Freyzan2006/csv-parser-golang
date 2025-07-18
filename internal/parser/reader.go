// internal/parser/reader.go
package parser

import (
    "encoding/csv"
    "fmt"
    "os"
)

func ReadCSV(path string) ([]map[string]string, error) {
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
    var result []map[string]string

    for _, row := range records[1:] {
        entry := make(map[string]string)
        for i, cell := range row {
            if i < len(headers) {
                entry[headers[i]] = cell
            }
        }
        result = append(result, entry)
    }

    return result, nil
}
