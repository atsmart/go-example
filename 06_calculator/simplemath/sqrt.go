package simplemath

import "math"

func Sqrt(num int) int {
	v := math.Sqrt(float64(num))
	return int(v)
}
