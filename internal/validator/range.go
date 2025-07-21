package validator

import (
	"fmt"
	"strconv"
)

import (
	"csv-parser/internal/model"
)


func RangeValidator(records []model.Record, ranger map[string][2]float64, verbose bool) ([]model.Record, error) {

	var valid []model.Record

	for i, rec := range records {
	
		for rangeKey, rangeValue := range ranger {


			if rec[rangeKey] == "" {
				continue
			}

			value, err := strconv.ParseFloat(rec[rangeKey], 64)
			if err != nil {
				return nil, fmt.Errorf("неверное значение поля '%s': %w", rangeKey, err)
			}

			if value < rangeValue[0] || value > rangeValue[1] {
				 if verbose {
					fmt.Printf("⚠️ Запись %d: значение поля '%s' выходит за допустимый диапазон\n", i+1, rangeKey)
				}
				continue
			}

			valid = append(valid, rec)
			
		}

	}


	return valid, nil
}