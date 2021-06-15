package fibonacci

import (
	"fmt"
	"math"
)

func Fibonacci(n interface{}) interface{} {
	switch v := n.(type) {
	case uint:
		return fibonacciUint(v)
	case int:
		return fibonacciInt(v)
	case float64:
		return fibonacciFloat(v)
	default:
		fmt.Printf("Type support for  %T! not implemented yet\n", v)
		return v
	}

}

func fibonacciUint(n uint) uint {
	return uint(calculate(float64(n)))

}

func fibonacciInt(n int) int {
	return int(calculate(float64(n)))
}

func fibonacciFloat(n float64) float64 {
	return calculate(n)
}

// math.Round((((1 + math.Sqrt(5))/2)^n - ((1 - math.Sqrt(5)) / 2)^n) / math.Sqrt(5))
func calculate(pos float64) float64 {
	part1 := math.Pow((1+math.Sqrt(5))/2, pos)
	part2 := math.Pow((1-math.Sqrt(5))/2, pos)
	result := (part1 - part2) / math.Sqrt(5)
	return math.Round(result)
}
