package common

import "math"

func getAverageFloat64(data []float64) float64 {
	sum := 0.0
	for _,v := range data {
		sum += v
	}
	return sum/float64(len(data))
}

func RoundFloat64(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixedFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundFloat64(num * output)) / output
}
