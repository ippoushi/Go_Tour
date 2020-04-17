package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	res := 0.0

	for {

		res = z - ( z*z - x ) / ( 2*z )
		if math.Abs(res-z) < 0.001 { break }

		fmt.Printf("Sqrt x   >>> %f\n", x)
		fmt.Printf("Sqrt z   >>> %f\n", z)
		fmt.Printf("Sqrt res >>> %f\n", res)
		fmt.Printf("Sqrt (z - x) >>> %f\n\n", z- x)

		z=res
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
}
