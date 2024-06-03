package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
)

type CubicSpline struct {
	splines []SplineTuple
}

type SplineTuple struct {
	a, b, c, d, x float64
}

func main() {
	functions := []struct {
		f    func(float64) float64
		name string
	}{
		{f: math.Exp, name: "e^x"},
		{f: math.Sinh, name: "sh(x)"},
		{f: math.Cosh, name: "ch(x)"},
		{f: math.Sin, name: "sin(x)"},
		{f: math.Cos, name: "cos(x)"},
		{f: math.Log, name: "ln(x)"},
		{f: func(x float64) float64 { return math.Exp(-x) }, name: "e^-x"},
	}
	X := 1.17
	n := 6

	x := []float64{1.0, 1.04, 1.08, 1.12, 1.16, 1.2}

	for _, function := range functions {
		var y []float64 = make([]float64, n)

		var cs CubicSpline

		for i, e := range x {
			y[i] = function.f(e)
		}
		cs.BuildSpline(x, y, n)
		res := cs.Interpolate(X)
		fmt.Printf("x = %.4f, y=%.6f, f(x)=%s\n", X, res, function.name)
	}
	left := x[0]
	right := 1.2

	step := 0.04
	for _, function := range functions {
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Кубический сплайн функции %s", function.name)
		p.X.Label.Text = "x"
		p.Y.Label.Text = "y"
		pts := make(plotter.XYs, int((right-left)/step)+1)
		var y []float64 = make([]float64, n)

		var cs CubicSpline

		for i, e := range x {
			y[i] = function.f(e)
		}
		cs.BuildSpline(x, y, n)
		index := 0
		for coor_x := left; coor_x < right; coor_x += step {

			pts[index] = plotter.XY{X: coor_x, Y: cs.Interpolate(coor_x)}
			index++
		}
		l, _ := plotter.NewLine(pts)
		p.Add(l)
		if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("%s.png", function.name)); err != nil {
			panic(err)
		}

	}
	//left := -1.0
	//right := 10.0
}

func (cs *CubicSpline) BuildSpline(x []float64, y []float64, n int) {
	cs.splines = make([]SplineTuple, n)
	for i := 0; i < n; i++ {
		cs.splines[i].x = x[i]
		cs.splines[i].a = y[i]
	}
	cs.splines[0].c = 0.0
	cs.splines[n-1].c = 0.0

	alpha := make([]float64, n-1)
	beta := make([]float64, n-1)
	alpha[0] = 0.0
	beta[0] = 0.0
	for i := 1; i < n-1; i++ {
		hi := x[i] - x[i-1]
		hi1 := x[i+1] - x[i]
		A := hi
		C := 2.0 * (hi + hi1)
		B := hi1
		F := 6.0 * ((y[i+1]-y[i])/hi1 - (y[i]-y[i-1])/hi)
		z := (A*alpha[i-1] + C)
		alpha[i] = -B / z
		beta[i] = (F - A*beta[i-1]) / z
	}

	for i := n - 2; i > 0; i-- {
		cs.splines[i].c = alpha[i]*cs.splines[i+1].c + beta[i]
	}

	for i := n - 1; i > 0; i-- {
		hi := x[i] - x[i-1]
		cs.splines[i].d = (cs.splines[i].c - cs.splines[i-1].c) / hi
		cs.splines[i].b = hi*(2.0*cs.splines[i].c+cs.splines[i-1].c)/6.0 + (y[i]-y[i-1])/hi
	}
}

func (cs *CubicSpline) Interpolate(x float64) float64 {
	if cs.splines == nil {
		return math.NaN()
	}

	n := len(cs.splines)
	var s SplineTuple

	if x <= cs.splines[0].x {
		s = cs.splines[0]
	} else if x >= cs.splines[n-1].x {
		s = cs.splines[n-1]
	} else {
		i := 0
		j := n - 1
		for i+1 < j {
			k := i + (j-i)/2
			if x <= cs.splines[k].x {
				j = k
			} else {
				i = k
			}
		}
		s = cs.splines[j]
	}

	dx := x - s.x
	return s.a + (s.b+(s.c/2.0+s.d*dx/6.0)*dx)*dx
}
