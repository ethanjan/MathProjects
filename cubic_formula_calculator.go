package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	// This declares all of the necessary initial variables.
	var (
		a, b, c, d float64
	)
	// This asks for the necessary user input.
	fmt.Println("\n\x1b[31mPlease enter the a, b, and c coefficients for the cubic equation",
		"equation, and it will be solved (ax³+bx²+cx+d=0).\x1b[30m")
	fmt.Println("\x1b[32ma:")
	fmt.Scanf("%f", &a)
	fmt.Println("b:")
	fmt.Scanf("%f", &b)
	fmt.Println("c:")
	fmt.Scanf("%f", &c)
	fmt.Println("d:")
	fmt.Scanf("%f", &d)
	fmt.Println("\n\x1b[35mYou want to solve:")
	fmt.Printf("%fx³+%fx²+%fx+%f=0\n\x1b[30m", a, b, c, d)
	// This calculates the discriminant.
	discriminant := (18 * a * b * c * d) - (4 * math.Pow(b, 3) * d) + (math.Pow(b, 2) * math.Pow(c, 2)) - (4 * a * math.Pow(c, 3)) - (27 * math.Pow(a, 2) * math.Pow(d, 2))
	delta0 := math.Pow(b, 2) - (3 * a * c)
	delta1 := (2 * math.Pow(b, 3)) - (9 * a * b * c) + (27 * math.Pow(a, 2) * d)
	if discriminant > 0 {
		// This converts the variables to complex.
		complexA := complex(a, 0)
		complexB := complex(b, 0)
		complexDelta0 := complex(delta0, 0)
		complexDelta1 := complex(delta1, 0)
		place := math.Pow((delta1+math.Sqrt(math.Pow(delta1, 2)-(4*math.Pow(delta0, 3))))/2, 1.0/3.0)
		C := cmplx.Pow((complexDelta1+cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
		if place == 0 {
			C = cmplx.Pow((complexDelta1-cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
		}
		primitiveCubeRootOfUnity := complex(-0.5, 0.5*math.Sqrt(3)*1)
		k1 := complex(1.0, 0)
		k2 := complex(2.0, 0)
		// This finds and prints the solutions.
		x1 := -1 / (3 * complexA) * (complexB + C + complexDelta0/C)
		x2 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k1) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k1)*C))
		x3 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k2) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k2)*C))
		fmt.Println("\x1b[36mAnswers:", x1, ",", x2, ",", x3, ".\n\x1b[30m")
	} else if discriminant < 0 {
		if delta0 == 0 {
			// This converts the variables to complex.
			complexA := complex(a, 0)
			complexB := complex(b, 0)
			complexDelta0 := complex(delta0, 0)
			complexDelta1 := complex(delta1, 0)
			place := math.Pow((delta1+math.Sqrt(math.Pow(delta1, 2)-(4*math.Pow(delta0, 3))))/2, 1.0/3.0)
			C := cmplx.Pow((complexDelta1+cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
			if place == 0 {
				C = cmplx.Pow((complexDelta1-cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
			}
			primitiveCubeRootOfUnity := complex(-0.5, 0.5*math.Sqrt(3)*1)
			k1 := complex(1.0, 0)
			k2 := complex(2.0, 0)
			// This calculates and displays the solutions.
			x1 := -1 / (3 * complexA) * (complexB + C + complexDelta0/C)
			x2 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k1) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k1)*C))
			x3 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k2) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k2)*C))
			fmt.Println("\x1b[36mAnswers:", x1, ",", x2, ",", x3, ".\n\x1b[30m")
		} else {
			// This converts the variables to complex.
			complexA := complex(a, 0)
			complexB := complex(b, 0)
			complexDelta0 := complex(delta0, 0)
			complexDelta1 := complex(delta1, 0)
			place := math.Pow((delta1+math.Sqrt(math.Pow(delta1, 2)-(4*math.Pow(delta0, 3))))/2, 1.0/3.0)
			C := cmplx.Pow((complexDelta1+cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
			if place == 0 {
				C = cmplx.Pow((complexDelta1-cmplx.Sqrt(cmplx.Pow(complexDelta1, 2)-(4*cmplx.Pow(complexDelta0, 3))))/2, 1.0/3.0)
			}
			primitiveCubeRootOfUnity := complex(-0.5, 0.5*math.Sqrt(3)*1)
			k1 := complex(1.0, 0)
			k2 := complex(2.0, 0)
			// This finds and prints the solutions.
			x1 := -1 / (3 * complexA) * (complexB + C + complexDelta0/C)
			x2 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k1) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k1)*C))
			x3 := -1 / (3 * complexA) * (complexB + (cmplx.Pow(primitiveCubeRootOfUnity, k2) * C) + complexDelta0/(cmplx.Pow(primitiveCubeRootOfUnity, k2)*C))
			fmt.Println("\x1b[36mAnswers:", x1, ",", x2, ",", x3, ".\n\x1b[30m")
		}
	} else {
		// This ensures that the cubic equation solver can also find the solutions for roots that have a multiplicity greater than 1.
		if delta0 == 0 {
			x1 := (-1 * b) / (3 * a)
			fmt.Println("\x1b[36mAnswers:", x1, ".\n\x1b[30m")
		} else {
			x1 := (9*a*d - b*c) / (2 * delta0)
			x2 := (4*a*b*c - 9*math.Pow(a, 2)*d - math.Pow(b, 3)) / (a * delta0)
			fmt.Println("\x1b[36mAnswers:", x1, ",", x2, ".\n\x1b[30m")
		}

	}
}
