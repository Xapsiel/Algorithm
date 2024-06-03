package main

import (
	"fmt"
	"math"
)

func foo(x float64) float64 {
	return math.Pow(math.E, x) + math.Pow(math.E, -3*x) - 4
}
func foo2(x float64) float64 {
	return math.Pow(math.E, x) + math.Pow(math.E, 4*x) - 4
}
func main() {
	newton(0.000001, -10, 30, foo)
}

func newton(eps, a, b float64, f func(float64) float64) {
	x := a
	h := 0.1
	for a0 := a; a0 < b; {
		fx := f(a0)
		fx2 := f(a0 + h)
		if fx*fx2 == 1 {
			continue
		}
		for i := 0; i < 100; i++ {
			sigma := (f(x) / chastnay_proivodnay(f, x, eps))
			x = x - sigma

			if x <= a0 {
				a0 = a0 + 10*h
				break
			}
			if x > b {
				a0 = a0 + 10*h
				break
			}
			if math.Abs(sigma) <= eps {
				fmt.Println(x)
				a0 = x + h

				break
			}
		}
		x = a0

	}

}

func chastnay_proivodnay(f func(float64) float64, x, h float64) float64 {
	return f(x+h)/h - f(x)/h
}
