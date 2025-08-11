// package parser

// import (
//     "encoding/csv"
//     "fmt"
//     "os"
//     "csv-parser/internal/model"
// )

// func ReadCSV(path string, headerEnabled bool) ([]model.Record, error) {
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

//     if len(records) == 0 {
//         return nil, fmt.Errorf("CSV пустой")
//     }

//     var headers []string
//     var data [][]string

//     if headerEnabled {
//         headers = records[0]
//         data = records[1:]
//     } else {
//         // Генерация заголовков: col_1, col_2, ...
//         numCols := len(records[0])
//         for i := 0; i < numCols; i++ {
//             headers = append(headers, fmt.Sprintf("col_%d", i+1))
//         }
//         data = records
//     }

//     var result []model.Record
//     for _, row := range data {
//         record := make(model.Record)
//         for i, cell := range row {
//             if i < len(headers) {
//                 record[headers[i]] = cell
//             }
//         }
//         result = append(result, record)
//     }

//     return result, nil
// }


package parser

import (
    "csv-parser/internal/model"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strings"

    "golang.org/x/text/encoding/charmap"
    "golang.org/x/text/transform"
)

// ReadCSV читает CSV с поддержкой разных кодировок (utf-8, windows-1251 и т.д.)
func ReadCSV(path string, headerEnabled bool, encoding string) ([]model.Record, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("не удалось открыть файл: %w", err)
    }
    defer file.Close()

    var reader io.Reader = file

    // Декодируем в UTF-8, если кодировка не utf-8
    switch strings.ToLower(encoding) {
    case "windows-1251", "cp1251":
        reader = transform.NewReader(file, charmap.Windows1251.NewDecoder())
    case "iso-8859-1":
        reader = transform.NewReader(file, charmap.ISO8859_1.NewDecoder())
    case "utf-8", "":
        // Ничего не делаем
    default:
        return nil, fmt.Errorf("неизвестная кодировка: %s", encoding)
    }

    csvReader := csv.NewReader(reader)
    records, err := csvReader.ReadAll()
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
