package main 

import (
    "flag"
    "fmt"
    "log"

    "csv-parser/internal/parser"
)

func main() {
    filePath := flag.String("file", "", "Путь к CSV-файлу")
    flag.Parse()

    if *filePath == "" {
        log.Fatal("Укажите путь к CSV-файлу с помощью флага -file")
    }

	records, err := parser.ReadCSV(*filePath)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	
	for i, r := range records {
		fmt.Printf("Запись %d:\n", i+1)
		for k, v := range r {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
	
}
