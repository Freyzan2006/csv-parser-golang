package validator

import (
	"fmt"
	"sort"
	"strings"
	"strconv"

	"csv-parser/internal/model"
)

func Sort(records []model.Record, sortParam string, verbose bool) ([]model.Record, error) {
	if sortParam == "" {
		return records, nil
	}

	parts := strings.Split(sortParam, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("неверный формат сортировки: %s", sortParam)
	}

	field := parts[0]
	order := strings.ToLower(parts[1])

	// сортируем in-place
	sort.SliceStable(records, func(i, j int) bool {
		a := records[i][field]
		b := records[j][field]

		// Пробуем как числа
		aNum, errA := strconv.ParseFloat(a, 64)
		bNum, errB := strconv.ParseFloat(b, 64)

		if errA == nil && errB == nil {
			if order == "desc" {
				return aNum > bNum
			}
			return aNum < bNum
		} else {
			if verbose {
				fmt.Printf("⚠️ Запись %d: поле '%s' не число\n", i+1, field)
			}
		}

		// Иначе строковая сортировка
		if order == "desc" {
			return a > b
		}
		return a < b
	})

	return records, nil
}
