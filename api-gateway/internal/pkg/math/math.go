package math

import "math"

func RoundFloat2DecimalPrecison(f float64) float64 {
	return math.Round(f*100) / 100
}
