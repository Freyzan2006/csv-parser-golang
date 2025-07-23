package validator

import (
	"fmt"
	"strconv"
)

import (
	"csv-parser/internal/model"
	"csv-parser/pkg"
)

func Filter(records []model.Record, raw string, verbose bool) ([]model.Record, error) {
	filter, err := pkg.ParseFilter(raw)
	if err != nil {
		return nil, err
	}

	var result []model.Record

	for i, r := range records {
		rawVal, ok := r[filter.Field]
		if !ok {
			if verbose {
				fmt.Printf("⚠️ Запись %d: поле '%s' не найдено\n", i+1, filter.Field)
			}
			continue
		}

		val, err := strconv.ParseFloat(rawVal, 64)
		if err != nil {
			if verbose {
				fmt.Printf("⚠️ Запись %d: значение '%s' не число\n", i+1, rawVal)
			}
			continue
		}

		if match(val, filter.Operator, filter.Value) {
			result = append(result, r)
		} else if verbose {
			fmt.Printf("⚠️ Запись %d не соответствует фильтру\n", i+1)
		}
	}

	return result, nil
}

// простая логика операторов
func match(val float64, op string, target float64) bool {
	switch op {
	case ">":
		return val > target
	case "<":
		return val < target
	case ">=":
		return val >= target
	case "<=":
		return val <= target
	case "==":
		return val == target
	case "!=":
		return val != target
	default:
		return false
	}
}
