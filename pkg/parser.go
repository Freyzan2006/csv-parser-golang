package pkg

import (
	"fmt"
	"strings"
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
