package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 3*x*x - 14*x - 5
}
func main() {
	left := -1000.6
	right := 412.0
	step := 0.1
	ditoxomia(fo2o, left, right, 10, step)
}

func foo(x float64) float64 {
	return math.Pow(math.E, x) + math.Pow(math.E, -3*x) - 4
}
func fo2o(x float64) float64 {
	return x - math.Sin(2*x)
}

func test(x float64) float64 {
	return 3*math.Pow(x, 4) - 5*math.Pow(x, 3) - 17*math.Pow(x, 2) + 13*x + 6
}
func ditoxomia(foo func(float64) float64, a0, b0 float64, e float64, h float64) {
	eps := 1 / math.Pow(10, e)
	a := a0
	b := a + h
	for b <= b0 {
		for {
			x := (a + b) / 2.0
			fa := foo(a)
			fx := foo(x)
			if fx*fa > 0 {
				a = x
			} else {
				b = x
			}

			if math.Abs(b-a) < eps && math.Abs(fx) < eps*10 {
				fmt.Println(x)
				break
			}
			if math.Abs(b-a) < eps {
				break
			}

		}
		a0 += h
		a = a0
		b = a + h

	}

}
