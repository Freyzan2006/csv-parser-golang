package pkg

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

func ParseCommaList(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func ParseKeyValueMap(s, sep string) map[string]string {
	result := make(map[string]string)
	if s == "" {
		return result
	}
	pairs := strings.Split(s, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, sep, 2)
		if len(kv) == 2 {
			result[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return result
}

func ParseRangeMap(s string) map[string][2]float64 {
	result := make(map[string][2]float64)
	if s == "" {
		return result
	}
	pairs := strings.Split(s, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) != 2 {
			continue
		}
		field := strings.TrimSpace(kv[0])
		bounds := strings.SplitN(kv[1], "-", 2)
		if len(bounds) != 2 {
			continue
		}
		var min, max float64
		fmt.Sscanf(bounds[0], "%f", &min)
		fmt.Sscanf(bounds[1], "%f", &max)
		result[field] = [2]float64{min, max}
	}
	return result
}





type FilterCondition struct {
	Field    string
	Operator string
	Value    float64
}


func ParseFilter(input string) (FilterCondition, error) {
	re := regexp.MustCompile(`^(\w+)\s*(==|!=|>=|<=|>|<)\s*(\d+\.?\d*)$`)
	matches := re.FindStringSubmatch(strings.TrimSpace(input))
	if len(matches) != 4 {
		return FilterCondition{}, fmt.Errorf("неверный формат фильтра: %s", input)
	}

	val, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return FilterCondition{}, fmt.Errorf("не удалось преобразовать значение: %w", err)
	}

	return FilterCondition{
		Field:    matches[1],
		Operator: matches[2],
		Value:    val,
	}, nil
}






type AggregateOp struct {
	Column string
	Op     string // sum, avg, max, min
}

func ParseAggregates(arg string) ([]AggregateOp, error) {
	if arg == "" {
		return nil, nil
	}

	var result []AggregateOp
	parts := strings.Split(arg, ",")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) != 2 {
			return nil, fmt.Errorf("неверный формат агрегата: %s", part)
		}
		col := kv[0]
		op := kv[1]
		if op != "sum" && op != "avg" && op != "min" && op != "max" {
			return nil, fmt.Errorf("неподдерживаемая операция агрегата: %s", op)
		}
		result = append(result, AggregateOp{Column: col, Op: op})
	}
	return result, nil
}
