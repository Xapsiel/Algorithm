package main

import (
	"math"
)

func Simpson(a, b, e float64, f func(float642 float64) float64, max_div, min_div int) float64 {
	var prev_sum float64 = f(a) + f(b)
	separator := min_div%2 + min_div     // количество делений(четное, больше 0)
	step := (b - a) / float64(separator) // шаг значений
	for i := 1; i < separator; i++ {
		if i%2 == 1 {
			prev_sum += 4 * f(a+(step*float64(i))) //
		} else {
			prev_sum += 2 * f(a+(step*float64(i)))
		}
	}
	prev_sum *= (step / 3)
	for {
		separator *= 2 // количество делений кратно увеличиваются
		if separator > max_div {
			return prev_sum
		}
		step /= 2              // шаг кратно уменьшается
		cur_sum := f(a) + f(b) // интеграл при новых данных
		i := 1
		for i = 1; i < separator; i++ {
			if i%2 == 1 {
				cur_sum += 4 * f(a+(step*float64(i)))
			} else {
				cur_sum += 2 * f(a+(step*float64(i)))
			}
		}
		cur_sum *= (step / 3)
		if (math.Abs(cur_sum-prev_sum)/15) < e && a+(step*float64(i)) == b { // Проверка на точность по Рунге
			break
		} else if (math.Abs(cur_sum-prev_sum) / 15) < e {

		}
		prev_sum = cur_sum
	}
	return prev_sum

}
