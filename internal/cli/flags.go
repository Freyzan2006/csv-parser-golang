// internal/cli/flags.go
package cli

import (
	"flag"
	"csv-parser/pkg"
	"csv-parser/internal/config"
)

import (
	"log"
)




func ParseFlags() *config.Config {
	var (
		filePath     string
		requiredRaw  string
		validateRaw  string
		rangeRaw     string
		verbose      bool
		filterRow    string
		sortRaw      string
		header       bool
		outFormat    string
		aggregate    string
		search 		 string
		page         int
		perPage      int 
		compareFile  string
	)

	flag.StringVar(&filePath, "file", "", "Путь к CSV-файлу")
	flag.StringVar(&requiredRaw, "required", "", "Обязательные поля (через запятую)")
	flag.StringVar(&validateRaw, "validate-type", "", "Проверка типов (пример: Age:int,Price:float)")
	flag.StringVar(&rangeRaw, "range", "", "Диапазоны значений (пример: Age:18-65)")
	flag.BoolVar(&verbose, "verbose", false, "Вывод подробной информации")
	flag.StringVar(&filterRow, "filter", "", "Фильтрация по строке (пример: Age>100)")
	flag.StringVar(&sortRaw, "sort", "", "Сортировка по колонке (пример: Age:desc или Name:asc)")
	flag.BoolVar(&header, "header", true, "Учитывать ли заголовки в первой строке")
	flag.StringVar(&outFormat, "out-format", "csv", "Формат вывода: csv, json, excel")
	flag.StringVar(&aggregate, "aggregate", "", "Агрегации вида col:op, например Age:sum или Price:avg")
	flag.StringVar(&search, "search", "", "Поиск строк по значению, пример: name=John")
	flag.IntVar(&page, "page", 1, "Номер страницы для постраничного вывода (начиная с 1)")
	flag.IntVar(&perPage, "per-page", 10, "Количество записей на странице")
	flag.StringVar(&compareFile, "compare", "", "Сравнить с другим CSV файлом")

	flag.Parse()

	aggOps, err := pkg.ParseAggregates(aggregate)
	if err != nil {
		log.Fatalf("ошибка парсинга агрегатов: %v", err)
	}


	

	return &config.Config{
		FilePath:     filePath,
		Required:     pkg.ParseCommaList(requiredRaw),
		ValidateType: pkg.ParseKeyValueMap(validateRaw, ":"),
		Range:        pkg.ParseRangeMap(rangeRaw),
		Verbose:      verbose,
		Filter:    	  filterRow,
		Sort:         sortRaw,
		Header:       header,
		Export:    	  outFormat,
		Aggregate:    aggOps,
		Search:       search,
		Page:     	  page,
		PerPage:  	  perPage,
		CompareFile: compareFile,
	}
}
