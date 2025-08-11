package service

import (
	"fmt"
	"strings"
	"sort"
)

import (
	"csv-parser/internal/config"
	"csv-parser/internal/model"
	"csv-parser/internal/parser"
)

func CompareService(cfg *config.Config, records []model.Record) ([]model.Record, error) {
	if cfg.CompareFile == "" {
		return records, nil
	}

	otherRecords, err := parser.ReadCSV(cfg.CompareFile, cfg.Header, cfg.Encoding)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения второго CSV: %w", err)
	}


	toKey := func(r model.Record) string {
		keys := make([]string, 0, len(r))
		for k := range r {
			keys = append(keys, k)
		}
		sort.Strings(keys) 
		var sb strings.Builder
		for _, k := range keys {
			sb.WriteString(fmt.Sprintf("%s=%s;", k, r[k]))
		}
		return sb.String()
	}

	firstSet := make(map[string]bool)
	for _, r := range records {
		firstSet[toKey(r)] = true
	}

	secondSet := make(map[string]bool)
	for _, r := range otherRecords {
		secondSet[toKey(r)] = true
	}

	fmt.Println("=== Различия между файлами ===")


	for _, r := range records {
		if !secondSet[toKey(r)] {
			fmt.Println("- Удалена:", r)
		}
	}


	for _, r := range otherRecords {
		if !firstSet[toKey(r)] {
			fmt.Println("+ Добавлена:", r)
		}
	}

	return records, nil
}


