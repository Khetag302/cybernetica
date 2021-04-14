package internal

import "math"

func Pow(a, n int) int32 {
	var z = math.Pow(float64(a), float64(n))
	return int32(z)
}
