package common

import "math"

func getMinMaxFloat64(data []float64) (float64, float64) {
	startMin := 1000000.0
	min := startMin
	max := 0.0
	for _,v := range data {
		if v == EmptyWeather {
			continue
		}
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	if min == startMin {
		min = 0.0
	}

	return min, max
}

func GetAverageFloat64(data []float64) float64 {
	sum := 0.0
	length := 0.0
	for _,v := range data {
		if v == EmptyWeather {
			continue
		}
		sum += v
		length++
	}
	return sum/length
}

func RoundFloat64(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixedFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundFloat64(num * output)) / output
}
