// package parser

// import (
//     "encoding/csv"
//     "fmt"
//     "os"
// )



// func ReadCSV(path string) ([]model.Record, error) {
//     file, err := os.Open(path)
//     if err != nil {
//         return nil, fmt.Errorf("не удалось открыть файл: %w", err)
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         return nil, fmt.Errorf("ошибка при чтении CSV: %w", err)
//     }

//     if len(records) < 1 {
//         return nil, fmt.Errorf("пустой файл или нет заголовков")
//     }

//     headers := records[0]
//     var result []model.Record

//     for _, row := range records[1:] {
//         entry := make(model.Record)
//         for i, cell := range row {
//             if i < len(headers) {
//                 entry[headers[i]] = cell
//             }
//         }
//         result = append(result, entry)
//     }


//     return result, nil
// }


package parser

import (
    "encoding/csv"
    "fmt"
    "os"
    "csv-parser/internal/model"
)

func ReadCSV(path string, headerEnabled bool) ([]model.Record, error) {
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

    if len(records) == 0 {
        return nil, fmt.Errorf("CSV пустой")
    }

    var headers []string
    var data [][]string

    if headerEnabled {
        headers = records[0]
        data = records[1:]
    } else {
        // Генерация заголовков: col_1, col_2, ...
        numCols := len(records[0])
        for i := 0; i < numCols; i++ {
            headers = append(headers, fmt.Sprintf("col_%d", i+1))
        }
        data = records
    }

    var result []model.Record
    for _, row := range data {
        record := make(model.Record)
        for i, cell := range row {
            if i < len(headers) {
                record[headers[i]] = cell
            }
        }
        result = append(result, record)
    }

    return result, nil
}
