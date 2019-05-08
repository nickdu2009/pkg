package xmath

import (
	"fmt"
	"math"
	"strconv"
)

var EPSILON float64 = 0.00000001

func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	}
	return false
}

func FloatIsZero(v float64) bool {
	return FloatEquals(v, 0)
}

func Round(input float64, places int) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", places), input), 64)
	return v
}
