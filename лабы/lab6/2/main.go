package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Point представляет точку в двумерном пространстве
type Point struct {
	x, y float64
}

// f - целевая функция, которую мы хотим минимизировать
func f(point Point) float64 {
	return point.x*point.x*point.x + 13.5*point.x*point.x - 66*point.x + (5*point.x-point.y)*(5*point.x-point.y)
}
func f2(point Point) float64 {
	return point.x*point.x + point.y*point.y
}

// Add складывает две точки
func Add(left, right Point) Point {
	return Point{left.x + right.x, left.y + right.y}
}

// Minus вычитает одну точку из другой
func Minus(left, right Point) Point {
	return Point{left.x - right.x, left.y - right.y}
}

// Mul умножает точку на скаляр
func Mul(num float64, right Point) Point {
	return Point{right.x * num, right.y * num}
}

// is_equal проверяет равенство двух точек с точностью до 1e-9
func is_equal(left, right Point) bool {
	return left.x == right.x && left.y == right.y
}

// inner выполняет локальное исследование функции в окрестности точки b1 с шагом h
func inner(b1 Point, h float64, f func(Point) float64) Point {
	fx := f(b1) // Вычисляем значение функции в точке b1

	// Исследование вдоль оси x
	fx1 := f(Add(b1, Mul(h, Point{1, 0}))) // Вычисляем значение функции в точке b1 + (h, 0)
	if fx1 < fx {
		b1 = Add(b1, Mul(h, Point{1, 0})) // Обновляем b1 на новую точку
		fx = fx1
	} else {
		fx2 := f(Add(b1, Mul(-h, Point{1, 0}))) // Вычисляем значение функции в точке b1 - (h, 0)
		if fx2 < fx {
			b1 = Add(b1, Mul(-h, Point{1, 0})) // Обновляем b1 на новую точку
			fx = fx2
		}
	}

	// Исследование вдоль оси y
	fx3 := f(Add(b1, Mul(h, Point{0, 1}))) // Вычисляем значение функции в точке b1 + (0, h)
	if fx3 < fx {
		b1 = Add(b1, Mul(h, Point{0, 1})) // Обновляем b1 на новую точку
		fx = fx3
	} else {
		fx4 := f(Add(b1, Mul(-h, Point{0, 1}))) // Вычисляем значение функции в точке b1 - (0, h)
		if fx4 < fx {
			b1 = Add(b1, Mul(-h, Point{0, 1})) // Обновляем b1 на новую точку
			fx = fx4
		}
	}

	return b1 // Возвращаем новую точку b1
}

// outer выполняет глобальное исследование функции, уменьшая шаг до достижения заданной точности
func outer(bPrev Point, h, e float64, f func(Point) float64) Point {
	for {
		bNext := inner(bPrev, h, f) // Выполняем локальное исследование
		if is_equal(bPrev, bNext) { // Если новая точка равна предыдущей с точностью до 1e-9
			h /= 2     // Уменьшаем шаг в 2 раза
			if h < e { // Если шаг меньше заданной точности
				break // Выходим из цикла
			}
		} else {
			bPrev = bNext // Обновляем базисную точку
			break         // Выходим из цикла
		}
	}
	return bPrev // Возвращаем новую базисную точку
}

func main() {
	bPrev := Point{1, 1} // Начальная точка
	start := bPrev
	H := 0.1     // Начальный шаг
	h := H       // Текущий шаг
	e := 0.00001 // Точность

	for {
		// Выполнить исследование
		b1 := outer(bPrev, h, e, f)

		// Проверка уменьшения функции
		if f(b1) < f(bPrev) {
			// Взять новую базисную точку

			// Увеличить шаг поиска по образцу
			pPrev := Add(bPrev, Mul(2, Minus(b1, bPrev)))
			bPrev = b1

			// Выполнить исследование
			p1 := outer(pPrev, h, e, f)

			// Проверка уменьшения функции
			if f(p1) < f(b1) {
				b1 = p1
			} else {
				h /= 2
				if h < e {
					break
				}
			}
		} else {
			// Уменьшить длину шага
			h /= 2
			if h < e {
				break
			}
		}
	}

	// Вывод результата
	fmt.Printf("Минимум найден в точке : (%f, %f)\n", bPrev.x, bPrev.y)
	fmt.Printf("Он равен: %f\n", f(bPrev))
	Plot(bPrev, start)
}
func Plot(p Point, start Point) {
	scatter3D := charts.NewScatter3D()
	scatter3D.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Супер График 2024",
		}),
	)

	data := make([]opts.ScatterData, 0)
	for x := -20.; x <= 20.; x += 0.1 {
		for y := -20.; y <= 20; y += 0.1 {
			z := f(Point{float64(x), float64(y)})
			data = append(data, opts.ScatterData{

				Value: []interface{}{x, y, z},
			})
		}
	}

	scatter3D.AddSeries("График", convertToChart3DData(data, p, false, "red"))
	point := []opts.ScatterData{
		{Value: []interface{}{p.x, p.y, f(p)}},
	}
	point_start := []opts.ScatterData{
		{Value: []interface{}{start.x, start.y, f(start)}},
	}

	// Добавление точки на график
	scatter3D.AddSeries("Минимум", convertToChart3DData(point, p, false, "green"))
	scatter3D.AddSeries("Точка Старта", convertToChart3DData(point_start, start, true, "yellow"))

	page := components.NewPage()
	page.AddCharts(scatter3D)

	f, err := os.Create("Plot.html")
	if err != nil {
		log.Println(err)
	}
	page.Render(f)

}

func convertToChart3DData(scatterData []opts.ScatterData, point Point, flag bool, color string) []opts.Chart3DData {
	chart3DData := make([]opts.Chart3DData, 0)

	for _, scatter := range scatterData {
		var chart3D opts.Chart3DData

		if scatter.Value.([]interface{})[0].(float64) == point.x && scatter.Value.([]interface{})[1].(float64) == point.y || flag {
			chart3D = opts.Chart3DData{

				Value: scatter.Value.([]interface{}),
				ItemStyle: &opts.ItemStyle{
					Color: color,
				},
			}
		} else {
			if flag != false {
				continue
			}
			chart3D = opts.Chart3DData{

				Value: scatter.Value.([]interface{}),
				ItemStyle: &opts.ItemStyle{
					Opacity: 0.009,
					Color:   color,
				},
			}
		}

		chart3DData = append(chart3DData, chart3D)
	}

	return chart3DData
}
