// internal/cli/flags.go
package cli

import (
	"flag"
	"csv-parser/pkg"
	"csv-parser/internal/config"
)

func ParseFlags() *config.Config {
	var (
		filePath     string
		requiredRaw  string
		validateRaw  string
		rangeRaw     string
		verbose      bool
	)

	flag.StringVar(&filePath, "file", "", "Путь к CSV-файлу")
	flag.StringVar(&requiredRaw, "required", "", "Обязательные поля (через запятую)")
	flag.StringVar(&validateRaw, "validate-type", "", "Проверка типов (пример: Age:int,Price:float)")
	flag.StringVar(&rangeRaw, "range", "", "Диапазоны значений (пример: Age:18-65)")
	flag.BoolVar(&verbose, "verbose", false, "Вывод подробной информации")


	flag.Parse()

	return &config.Config{
		FilePath:     filePath,
		Required:     pkg.ParseCommaList(requiredRaw),
		ValidateType: pkg.ParseKeyValueMap(validateRaw, ":"),
		Range:        pkg.ParseRangeMap(rangeRaw),
		Verbose:      verbose,
	}
}
