// package main

// import (
// 	"fmt"
// 	"log"

// 	"csv-parser/internal/cli"
// 	"csv-parser/internal/parser"
// )

// func main() {
	// cfg := cli.NewConfig()
	// flags := cfg.ParseFlags()

	// if flags.FilePath == "" {
	// 	log.Fatal("Укажите путь к CSV-файлу с помощью флага -file")
	// }

	// if flags.Required != "" {
	// 	fmt.Printf("Обязательные поля: %s\n", flags.Required)
	// }

	// records, err := parser.ReadCSV(flags.FilePath)
	// if err != nil {
	// 	log.Fatalf("Ошибка: %v", err)
	// }

	// for i, r := range records {
	// 	fmt.Printf("Запись %d:\n", i+1)
	// 	for k, v := range r {
	// 		fmt.Printf("  %s: %s\n", k, v)
	// 	}
	// }
// }


package main

import (
	"csv-parser/internal/cli"
	"csv-parser/internal/core"
	"csv-parser/internal/logger"
	"log"
)

func main() {
	cfg := cli.ParseFlags()

	logger.Init(cfg.LogFile)

	logger.Logger.Println("Старт программы")

	if err := core.Process(cfg); err != nil {
		log.Fatal(err)
	}

	logger.Logger.Println("Старт Завершено")
}
