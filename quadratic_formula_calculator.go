package main

// This imports all of the necessary packages.
import (
	"fmt"
	"math/cmplx"
	"math"
)

// This begins the program.
func main() {
	// This declares the necessary variables.
	var (
	a, b, c float64
	)
	// This asks the user to input the necessary values.
	fmt.Println("\n\x1b[31mPlease enter the a, b, and c coefficients for the quadratic",
		"equation, and it will be solved (ax²+bx+c=0).\x1b[30m")
	fmt.Println("\x1b[32ma:")
	fmt.Scanf("%f", &a)
	fmt.Println("b:")
	fmt.Scanf("%f", &b)
	fmt.Println("c:")
	fmt.Scanf("%f", &c)
	fmt.Println("\n\x1b[35mYou want to solve:")
	fmt.Printf("%fx²+%fx+%f=0\n\x1b[30m", a, b, c)
	// This calculates the discriminant and determines the solutions.
	discriminant := math.Pow(b, 2) - (4 * a * c)
	if discriminant < 0 {
		negativeDiscriminant := (cmplx.Sqrt(complex(discriminant, 0)))
		complexA := complex(a, 0)
		complexB := complex(b, 0)
		result1 := ((complexB * -1) + negativeDiscriminant) / (2 * complexA)
        	result2 := ((complexB * -1) - negativeDiscriminant) / (2 * complexA)
		fmt.Println("\x1b[36mThe solutions are:", result1, ",", result2, ".\n\x1b[30m")
	}
	if discriminant > 0 {
		result1 := ((b * -1) + math.Sqrt(discriminant)) / (2 * a)
                result2 := ((b * -1) - math.Sqrt(discriminant)) / (2 * a)
                fmt.Println("\x1b[36mThe solutions are:", result1, ",", result2, ".\n\x1b[30m")
	}
	if discriminant == 0 {
		result1 := ((b * -1) + math.Sqrt(discriminant)) / (2 * a)
                fmt.Println("\x1b[36mThe solution is:", result1, ".\n\x1b[30m")
	}
}
