package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func foo2(x float64) float64 {
	return math.Sin(x)
}
func foo(x float64) float64 {
	return math.Abs(math.Sin(x))
}
func main() {
	T := 2 * math.Pi
	N := 3
	left := 0.
	right := 2 * math.Pi
	a0, aCoeffs, bCoeffs := Furry(T, foo, N, 7, left, right)

	x := 2.5
	f := func(x float64) float64 {
		sum := 0.0
		for n := 1; n <= N; n++ {
			sum += (atCoeffs[n-1]*math.Cos(float64(n)*x) + bCoeffs[n-1]*math.Sin(float64(n)*x))
		}
		return a0/2 + sum
	}

	fmt.Printf("a0: %f\n", a0)
	fmt.Printf("aCoeffs: %v\n", aCoeffs)
	fmt.Printf("bCoeffs: %v\n", bCoeffs)
	fmt.Println(f(x))
	make_graphick(left, right, foo, f)

}
func make_graphick(left, right float64, foo func(float64) float64, furry func(float64) float64) {
	p := plot.New()
	p.Title.Text = "Разложение на ряды Фурье"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	//left := -1.0
	//right := 10.0
	step := 0.001
	pts := make(plotter.XYs, int((right-left)/step)+1)
	pts2 := make(plotter.XYs, int((right-left)/step)+1)
	index := 0
	for coor_x := left; coor_x < right; coor_x += step {
		pts2[index] = plotter.XY{X: coor_x, Y: furry(coor_x)}
		pts[index] = plotter.XY{X: coor_x, Y: foo(coor_x)}
		index++
	}
	l, _ := plotter.NewLine(pts)
	p.Add(l)
	l2, _ := plotter.NewLine(pts2)
	l2.Color = color.RGBA{255, 0, 0, 0}

	p.Add(l2)
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func Furry(T float64, foo func(float64) float64, N int, eps float64, left, right float64) (float64, []float64, []float64) {
	e := 1 / math.Pow(10, eps)
	a0 := (2 / T) * Simpson(left, right, e, 2, foo, 100)
	aCoeffs := make([]float64, N)
	bCoeffs := make([]float64, N)
	for n := 1; n <= N; n++ {
		an := (2 / T) * Simpson(left, right, e, 2, func(x float64) float64 {
			return foo(x) * math.Cos(float64(n)*x)
		}, 100000)
		bn := (2 / T) * Simpson(left, right, e, 2, func(x float64) float64 {
			return foo(x) * math.Sin(float64(n)*x)
		}, 100000)
		aCoeffs[n-1] = an
		bCoeffs[n-1] = bn
	}
	return a0, aCoeffs, bCoeffs
}

func Simpson(a, b, eps, n float64, f func(float64) float64, max_n float64) float64 {

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
				return S2
			}
		}
	}
	return 0
}
