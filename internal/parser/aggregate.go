package parser 

import (
	"strconv"
	"csv-parser/internal/model"
	"csv-parser/pkg"
)


func Aggregate(records []model.Record, ops []pkg.AggregateOp) (map[string]float64, error) {
	results := make(map[string]float64)
	counts := make(map[string]int)

	for _, op := range ops {
		results[op.Column+":"+op.Op] = 0
	}

	for _, rec := range records {
		for _, op := range ops {
			valStr, ok := rec[op.Column]
			if !ok {
				continue
			}
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				continue
			}

			key := op.Column + ":" + op.Op

			switch op.Op {
			case "sum", "avg":
				results[key] += val
				counts[key]++
			case "min":
				if counts[key] == 0 || val < results[key] {
					results[key] = val
				}
				counts[key]++
			case "max":
				if counts[key] == 0 || val > results[key] {
					results[key] = val
				}
				counts[key]++
			}
		}
	}

	// вычисляем среднее
	for _, op := range ops {
		key := op.Column + ":" + op.Op
		if op.Op == "avg" && counts[key] > 0 {
			results[key] = results[key] / float64(counts[key])
		}
	}

	return results, nil
}
