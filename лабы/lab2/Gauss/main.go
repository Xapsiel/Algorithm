package main

import (
	"fmt"
)

func main() {
	A := [][]float64{
		{100, 1, -1, 7},
		{2, 3, 0, 7},
		{1, -1, 5, 11},
	}
	var size = 3
	//var size int
	//fmt.Print("Введите размерность матрицы: ")
	//fmt.Scan(&size)
	//fmt.Println()
	//
	//A := make([][]float64, size)
	//for row := 0; row < size; row++ {
	//	fmt.Println("Введите коэффициенты левой части", row+1, "уравнения")
	//	A[row] = make([]float64, size)
	//	for col := 0; col < size; col++ {
	//		var elem float64
	//		fmt.Scan(&elem)
	//		A[row][col] = elem
	//	}
	//}
	//for row := 0; row < size; row++ {
	//	fmt.Println("Введите чему равно", row+1, "уравнение")
	//	var elem float64
	//	fmt.Scan(&elem)
	//	A[row] = append(A[row], elem)
	//}
	gauss(A, size)
}

func gauss(A [][]float64, n int) {
	for index := 0; index < n-1; index++ {
		st_elem := (-1) * A[index][index]      // Первый элемент строки которую умножаем
		for row := index + 1; row < n; row++ { // следующие после умножаемой строки строки
			elem := A[row][index] / st_elem      //коэффициент умножения
			for col := index; col < n+1; col++ { //итерируемся по столбцам строки
				A[row][col] += (elem * A[index][col]) //умножаем на коэффициент
			}
		}
	}
	for index := n - 1; index >= 0; index-- { //Начиная с последней строки
		for i := index + 1; i < n; i++ { // Добавляем к правой части уравнения известные X умноженные на коэффициент
			A[index][n] += (-1) * (A[index][i] * A[i][i])
		}
		A[index][index] = A[index][n] / A[index][index] // находим X
		fmt.Printf("x%d = %f\n", index+1, A[index][index])
	}

}
