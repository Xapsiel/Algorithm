package main

import (
	"fmt"
	"math"
)

func first(x, y float64) float64 {
	return math.Cos(y) + x - 1.5
}
func second(x, y float64) float64 {
	return 2*y - math.Sin(x-0.5) - 0.8
}
func foo(x, y float64) float64 {
	return math.Sin(x+1) - y - 1.2
}
func foo2(x, y float64) float64 {
	return 223*x + math.Cos(y) - 2
}
func main() {
	foos := []func(float64, float64) float64{
		first,
		second,
	}
	NewtonRawson(2, 0.00001, 0, 0, foos)

}

func NewtonRawson(n int, eps float64, x, y float64, AllFunc []func(float64, float64) float64) {
	var F []float64 = make([]float64, n) // Создаем вектор F
	k := 0
	var X []float64                     // Инициализируем вектор X
	for iter := 0; iter < 200; iter++ { // Итерации метода Ньютона-Рафсона
		for i, e := range AllFunc {
			F[i] = -1 * e(x, y) // Создаем вектор функции
		}
		JacobyMatrix := Jacoby(AllFunc, x, y, n) // Получаем матрицу Якоби
		X = gauss(JacobyMatrix, F)               // Решаем систему линейных уравнений методом Гаусса
		x -= X[0]                                // Обновляем значение x
		y -= X[1]                                // Обновляем значение y
		k++
		if math.Abs(X[0]) < eps && math.Abs(X[1]) < eps { // Проверяем условие выхода
			break
		}
	}
	fmt.Println(k)
	fmt.Println(x, y) // Выводим результат

	for _, e := range AllFunc {
		fmt.Println(e(x, y))
	}
}

func Jacoby(allFunc []func(float64, float64) float64, x, y float64, variable int) [][]float64 {
	var matrix [][]float64 = make([][]float64, variable) // Создаем матрицу Якоби
	for index, foo := range allFunc {
		matrix[index] = make([]float64, variable)
		x0, y0 := ChastnayProizvodnaya(foo, x, y, 0.000000001) // Вычисляем частные производные
		matrix[index][0] = x0
		matrix[index][1] = y0
	}
	return matrix
}

func ChastnayProizvodnaya(foo func(x, y float64) float64, x, y, h float64) (float64, float64) {
	fx := (foo(x+h, y) - foo(x, y)) / h // Вычисляем производную по x
	fy := (foo(x, y+h) - foo(x, y)) / h // Вычисляем производную по y
	return fx, fy
}

func gauss(W [][]float64, F []float64) []float64 {
	n := len(W)
	WF := make([][]float64, n)

	// Преобразование матрицы W к виду [W|F].
	for i := 0; i < n; i++ {
		WF[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			WF[i][j] = W[i][j]
		}
		WF[i][n] = -F[i]
	}

	// Прямой ход метода Гаусса.
	for i := 0; i < n; i++ {
		maxRow := i
		for j := i + 1; j < n; j++ {
			if math.Abs(WF[j][i]) > math.Abs(WF[maxRow][i]) {
				maxRow = j
			}
		}
		WF[i], WF[maxRow] = WF[maxRow], WF[i]

		for j := i + 1; j < n; j++ {
			factor := WF[j][i] / WF[i][i]
			for k := i; k < n+1; k++ {
				WF[j][k] -= factor * WF[i][k]
			}
		}
	}

	// Обратный ход метода Гаусса.
	X := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		X[i] = WF[i][n] / WF[i][i]
		for j := i - 1; j >= 0; j-- {
			WF[j][n] -= WF[j][i] * X[i]
		}
	}

	return X
}
