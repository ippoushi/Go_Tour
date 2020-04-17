package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0  
	res := 0.0
	
	for {
		res = z - ( z*z - x ) / ( 2*z )
		if math.Abs( res - z ) < 0.001 { break }
		
		fmt.Printf("Sqrt x >>> %\n", x)
		fmt.Printf("Sqrt z >>> %\n" z")

func main() {
	fmt.Println(Sqrt(2))
}

