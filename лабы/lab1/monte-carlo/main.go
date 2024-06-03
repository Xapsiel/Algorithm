package main

import (
	"fmt"
	"math"
	"math/rand"
)

func first(x float64) float64 {
	return math.Log(math.Sin(x)) - 1/(x*x)
}

func main() {
	var left, right float64
	var n float64
	fmt.Println("Введите левую и правую границу графика")
	fmt.Scan(&left, &right)
	fmt.Println("Введите количество бросков")
	fmt.Scanln(&n)
	area, eps := MonteCarlo(first, 5, n, left, right)
	fmt.Printf("Интеграл равен %v, средняя погрешность %v\n", area, eps)

}

func MonteCarlo(fun func(x float64) float64, iter int, n, left, right float64) (float64, float64) {
	point := make(map[string]int)          // Создание словаря для хранения точек
	var maxy = -10 * 1000000.0             // Максимальный y
	var miny = 10 * 1000000.0              // Минимальный y
	for i := left; i < right; i += 0.001 { // Находим наибольшую и наименьшую точку
		y := fun(float64(i))
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
	}
	if maxy < 0 {
		maxy = 0
	}
	if miny > 0 {
		miny = 0
	}
	var SquareOfElem float64
	var SumOfElem float64
	var res []float64
	for iteration := 0; iteration < iter; iteration++ { // iter независимых попыток
		sum := 0.0
		l := 0.0
		for i := 0.0; i < n; i++ { // n независимых бросков
			x := left + rand.Float64()*(right-left)
			y := miny + rand.Float64()*(maxy-miny)
			key := fmt.Sprintf("%v-%v", x, y)

			if point[key] == 1 || point[key] == -1 {
				continue
			} else {
				f := float64(int(fun(x)*1000)) / 1000
				l++
				if f > y && y > 0 {
					sum += 1
					point[key] = 1
				} else if f < y && y < 0 {
					sum -= 1
					point[key] = -1
				}
			}
		}
		// добавляем результат в список
		res = append(res, (sum/l)*math.Abs((right-left)*(maxy-miny)))
		SquareOfElem += res[iteration] * res[iteration] // складываем квадрат площади
		SumOfElem += res[iteration]                     // сумма элементов
	}

	return res[0], SquareOfElem/n - (SumOfElem*SumOfElem)/(n*n)
}
