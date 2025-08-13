// // // internal/printer/printer.go
// // package printer

// // import (
// // 	"fmt"
// // 	"csv-parser/internal/model"
// // )

// // func PrintRecords(records []model.Record) {
// // 	for i, r := range records {
// // 		fmt.Printf("Запись %d:\n", i+1)
// // 		for k, v := range r {
// // 			fmt.Printf("  %s: %s\n", k, v)
// // 		}
// // 	}
// // }


// package printer

// import (
// 	"csv-parser/internal/model"
// 	"fmt"
// )

// // ANSI-цвета
// const (
// 	colorReset  = "\033[0m"
// 	colorRed    = "\033[31m"
// 	colorGreen  = "\033[32m"
// 	colorYellow = "\033[33m"
// 	colorCyan   = "\033[36m"
// )

// func PrintRecords(records []model.Record) {
// 	for i, r := range records {
// 		fmt.Printf("%sЗапись %d:%s\n", colorCyan, i+1, colorReset)
// 		for k, v := range r {
// 			keyColor := colorYellow
// 			valueColor := colorReset

// 			// Если значение пустое — выделим красным
// 			if v == "" {
// 				valueColor = colorRed
// 			}

// 			fmt.Printf("  %s%s%s: %s%s%s\n",
// 				keyColor, k, colorReset,
// 				valueColor, v, colorReset)
// 		}
// 	}
// }


// internal/printer/printer.go
package printer

import (
	"fmt"
	"strings"

	"csv-parser/internal/model"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

func PrintRecords(records []model.Record) {
	if len(records) == 0 {
		fmt.Println(colorYellow + "Нет данных для отображения" + colorReset)
		return
	}

	// Получаем список ключей (заголовков)
	headers := getHeaders(records)

	// Определяем ширину каждой колонки
	colWidths := make(map[string]int)
	for _, h := range headers {
		colWidths[h] = len(h)
	}
	for _, r := range records {
		for _, h := range headers {
			if len(r[h]) > colWidths[h] {
				colWidths[h] = len(r[h])
			}
		}
	}

	// Печатаем заголовок
	for _, h := range headers {
		fmt.Printf(colorBlue+"%-*s"+colorReset+"  ", colWidths[h], h)
	}
	fmt.Println()

	// Печатаем разделитель
	for _, h := range headers {
		fmt.Print(strings.Repeat("-", colWidths[h]) + "  ")
	}
	fmt.Println()

	// Печатаем строки
	for _, r := range records {
		for _, h := range headers {
			val := r[h]
			if strings.Contains(strings.ToLower(val), "error") {
				fmt.Printf(colorRed+"%-*s"+colorReset+"  ", colWidths[h], val)
			} else {
				fmt.Printf(colorGreen+"%-*s"+colorReset+"  ", colWidths[h], val)
			}
		}
		fmt.Println()
	}
}

func getHeaders(records []model.Record) []string {
	if len(records) == 0 {
		return []string{}
	}
	headers := make([]string, 0, len(records[0]))
	for h := range records[0] {
		headers = append(headers, h)
	}
	return headers
}


