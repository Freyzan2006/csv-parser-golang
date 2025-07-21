package validator

import (
	"fmt"
	"strconv"
	"unicode"

	"csv-parser/internal/model"
)

func ValidateTypes(records []model.Record, types map[string]string, verbose bool) ([]model.Record, error) {
	var valid []model.Record

	for i, rec := range records {
		isValid := true

		for field, expectedType := range types {
			val := rec[field]

			if val == "" {
				continue // пропускаем пустые
			}

			switch expectedType {
			case "int":
				if _, err := strconv.Atoi(val); err != nil {
					isValid = false
					if verbose {
						fmt.Printf("⚠️ Запись %d: поле '%s' должно быть целым числом\n", i+1, field)
					}
				}
			case "float":
				if _, err := strconv.ParseFloat(val, 64); err != nil {
					isValid = false
					if verbose {
						fmt.Printf("⚠️ Запись %d: поле '%s' должно быть числом\n", i+1, field)
					}
				}
			case "string":
				// необязательно, но можно проверять что это хотя бы буквы/текст
				if !isMostlyLetters(val) && verbose {
					fmt.Printf("⚠️ Запись %d: поле '%s' содержит неожиданные символы для строки\n", i+1, field)
				}
			default:
				return nil, fmt.Errorf("неизвестный тип '%s' для поля '%s'", expectedType, field)
			}
		}

		if isValid {
			valid = append(valid, rec)
		}
	}

	return valid, nil
}

// необязательная доп. проверка
func isMostlyLetters(s string) bool {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count >= len(s)/2
}
