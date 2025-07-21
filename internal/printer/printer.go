// internal/printer/printer.go
package printer

import (
	"fmt"
	"csv-parser/internal/model"
)

func PrintRecords(records []model.Record) {
	for i, r := range records {
		fmt.Printf("Запись %d:\n", i+1)
		for k, v := range r {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
}
