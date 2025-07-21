package validator

import (
	"fmt"
	"strings"
)

import (
	"csv-parser/internal/model"
)

func Required(records []model.Record, requiredFields []string) ([]model.Record, error) {
	var valid []model.Record
	for i, rec := range records {
		ok := true
		for _, field := range requiredFields {
			if val, exists := rec[field]; !exists || strings.TrimSpace(val) == "" {
				fmt.Printf("⚠️ Запись %d: отсутствует обязательное поле '%s'\n", i+1, field)
				ok = false
				break
			}
		}
		if ok {
			valid = append(valid, rec)
		}
	}
	return valid, nil
}