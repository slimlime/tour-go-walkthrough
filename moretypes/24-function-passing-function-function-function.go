package main

import (
	"fmt"
	"math"
)

func computePresetTest(
	fn func(float64, float64) float64,
) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(
		x, y float64,
	) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(computePresetTest(hypot))
	fmt.Println(computePresetTest(math.Pow))
}

