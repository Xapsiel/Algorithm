package main

import (
	"fmt"
	"math"
)

func main() {
	A := [][]float64{
		{3, 0, -1, 7},
		{2, -5, 1, -2},
		{20, 2, 5, 1},
	}
	var size = 3
	fun(A, size, 40, 10)
}

func fun(A [][]float64, size int, maxiter int, e int) {
	eps := 1 / math.Pow(10, float64(e))
	var X map[int]float64 = make(map[int]float64)
	var E map[int]float64 = make(map[int]float64)
	for i := 0; i < size; i++ {
		var x float64
		fmt.Println("Введите начальное приближение для", i+1, "неизвестной: ")
		fmt.Scan(&x)
		E[i] = eps
		X[i] = x
	}

	for iter := 0; iter < maxiter; iter++ {
		matched := 0
		for row := range A {
			prevValue := X[row] //запоминаем Х полученное на предыдущей итерации
			X[row] = 0          //обнуляем его
			for col := range A[row] {

				if col != row && col != size { //пропускаем тот Х, который ищем и то, чему равно уравнение(Последний в массиве элемент)
					X[row] += -1 * (X[col]) * A[row][col]
				}

			}
			X[row] += A[row][size]                //Прибавляем значение уравнения
			X[row] /= A[row][row]                 //делим на коэффициент при рабочем Х
			if math.Abs(prevValue-X[row]) < eps { //находим погрешностЬ
				matched++
				E[row] = math.Abs(prevValue - X[row])
			}
		}
		if matched == size {
			break
		}
	}
	for x := 0; x < size; x++ {
		fmt.Printf("X%d равен %v, погрешность равна %.16f\n", x+1, X[x], E[x])
	}
}
