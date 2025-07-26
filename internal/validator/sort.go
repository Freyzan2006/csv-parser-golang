package validator

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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

	// Локальная функция сортировки по числам
	sortOfNumber := func(i, j int) bool {
		a, b := records[i][field], records[j][field]
		aNum, errA := strconv.ParseFloat(a, 64)
		bNum, errB := strconv.ParseFloat(b, 64)
		if errA == nil && errB == nil {
			if order == "desc" {
				return aNum > bNum
			}
			return aNum < bNum
		}
		return false // если хотя бы одно значение не число
	}

	// Локальная функция сортировки по длине строки
	sortOfString := func(i, j int) bool {
		a, b := records[i][field], records[j][field]
		if order == "desc" {
			return len(a) > len(b)
		}
		return len(a) < len(b)
	}

	sort.SliceStable(records, func(i, j int) bool {
		// Сначала пробуем сортировку как число
		a := records[i][field]
		b := records[j][field]
		_, errA := strconv.ParseFloat(a, 64)
		_, errB := strconv.ParseFloat(b, 64)

		if errA == nil && errB == nil {
			return sortOfNumber(i, j)
		}

		if verbose {
			fmt.Printf("⚠️ Запись %d или %d: поле '%s' не число, сортировка по длине строки\n", i+1, j+1, field)
		}
		return sortOfString(i, j)
	})

	return records, nil
}
