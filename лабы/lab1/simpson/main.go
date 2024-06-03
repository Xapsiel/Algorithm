package main

import (
	"fmt"
	"math"
)

func first(x float64) float64 {
	return (math.Sin(x))
}
func main() {

	var left float64
	var right float64
	var e float64
	fmt.Println("Введите левую границу интегрирования, правую, а также желаемая точность(показатель степени 10)")
	fmt.Scan(&left, &right, &e)
	e = 1.0 / (math.Pow(10.0, e))
	fmt.Println(S(left, math.Pi, e, 2, first, 10000))
}

func S(a, b, eps, n float64, f func(float64) float64, max_n float64) float64 {

	SumSimpson := func(a, b, step float64) float64 {
		sum := f(a) + f(b)
		counter := 0
		for a < b {
			a += step
			if counter%2 == 0 {
				sum += 4 * f(a)
			} else {
				sum += 2 * f(a)
			}
			counter++
		}
		return sum * step / 3
	}

	for {
		S1 := SumSimpson(a, b, (b-a)/n)
		S2 := SumSimpson(a, b, (b-a)/n/2)
		if math.Abs(S2-S1)/15 < eps {
			return S2
			break
		} else {
			n *= 2
			if n > max_n {
				fmt.Println("Достигнуто максимальное разделение")
				return S2
			}
		}
	}
	return 0
}
