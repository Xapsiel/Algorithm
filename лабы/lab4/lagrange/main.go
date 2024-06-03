package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func main() {
	var n int
	fmt.Println("Сколько точек известно?")
	fmt.Scan(&n)
	X, X_Y := make_X_XY(n)
	polinom := Langrage(n, X)
	var x float64
	fmt.Println("Значение функции в какой точке вы желаете узнать")
	fmt.Scan(&x)
	fmt.Println(SUM(x, X, X_Y, polinom))
	p := plot.New()
	p.Title.Text = "Интерполяция по Лагранжу"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	left := -1.0 + X[0]
	right := 1 + X[n-1]
	//left := -1.0
	//right := 10.0
	step := 0.01
	pts := make(plotter.XYs, int((right-left)/step)+1)
	index := 0
	for coor_x := left; coor_x < right; coor_x += step {

		pts[index] = plotter.XY{X: coor_x, Y: SUM(coor_x, X, X_Y, polinom)}
		index++
	}
	l, _ := plotter.NewLine(pts)
	p.Add(l)
	pts2 := make(plotter.XYs, int((right-left)/step)+1)
	index = 0
	for x, y := range X_Y {
		pts2[index].X = x
		pts2[index].Y = y
		index++
	}
	dot, _ := plotter.NewScatter(pts2)
	dot.GlyphStyle.Radius = vg.Points(4)
	dot.Color = color.RGBA{0, 255, 0, 0}

	p.Add(dot)
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func make_X_XY(n int) (map[int]float64, map[float64]float64) {
	var X map[int]float64 = make(map[int]float64, n)

	var X_Y map[float64]float64 = make(map[float64]float64, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Println("Введи x и y через пробел")
		fmt.Scan(&x, &y)
		X[i] = x
		X_Y[x] = y
	}
	return X, X_Y

}
func Langrage(n int, X map[int]float64) Polinom {

	var Znamenatel = make(map[int]map[string]float64, n)
	var Chiclitel = make(map[int]map[int]float64, n)
	for j := 0; j < n; j++ {
		Znamenatel[j] = make(map[string]float64, n)
		Znamenatel[j]["Знаменатель"] = 1
		Chiclitel[j] = make(map[int]float64, n)
		k := 0
		for i := 0; i < n; i++ {
			if i == j {
				continue
			}
			Znamenatel[j]["Знаменатель"] *= (X[j] - X[i])
			Chiclitel[j][k] = -X[i]

			k++
		}
	}
	return Polinom{
		Znamenatel: Znamenatel,
		Chiclitel:  Chiclitel,
	}

}

type Polinom struct {
	Znamenatel map[int]map[string]float64
	Chiclitel  map[int]map[int]float64
}

func SUM(x float64, X map[int]float64, X_Y map[float64]float64, polinom Polinom) float64 {
	sum := 0.0
	for j := range polinom.Znamenatel {
		elem := 1.0
		for i := range polinom.Chiclitel[j] {
			elem *= (x + polinom.Chiclitel[j][i])
		}
		elem /= polinom.Znamenatel[j]["Знаменатель"]
		elem *= X_Y[X[j]]
		sum += elem
	}
	return sum
}
