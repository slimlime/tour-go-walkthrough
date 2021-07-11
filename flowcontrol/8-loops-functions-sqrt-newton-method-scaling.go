// Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(
	x float64,
) (
	z float64,
) {
	// Strange bouncing around from arbitrary guess rather than backwards?
	z = 1.0
	adjustmentError:= 1.0
	
	for ; math.Abs(adjustmentError) > 0.0000000000001;  {
		guess := z * z
		adjustmentError = guess - x
		scaledDenominator := 2 * z
		z -= adjustmentError / scaledDenominator

		fmt.Printf("z %.2f, x %.2f, guess-squared %.2f, adjust %.2f, scaledDenominator %.2f\n",
			z,
			x,
			guess,
			adjustmentError,
			scaledDenominator,
		)
		// "x", x, "guess", guess, adjustmentError, scaledDenominator, z)
	}

	return
}

//
// (Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)

//
func main() {
	fmt.Println("custom Sqrt", Sqrt(16))
	fmt.Println("math   Sqrt", math.Sqrt(16))

	// 1.4166666666666667
	// 1.41421356237
	// Bouncing between numbers
	// 1.4142135623730951
	// 1.414213562373095

	// custom 1.414213562373095
	// math   1.4142135623730951



	testArr := [...]float64{1, 2, 3}
	testeArr := []float64{1, 2, 3}
	someMapIterable := map[int]float64{1: 1.1}
	
	for i := range testArr {
		fmt.Println(i)
	}
	for i := range testeArr {
		fmt.Println(i)
	}
	for k, v := range someMapIterable {
		fmt.Println(k, v)
	}

}

