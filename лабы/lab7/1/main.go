package main

import "fmt"

func F(x1, x2 int) int {
	return 3*x1 + 2*x2
}

type answer struct {
	f      int
	x1, x2 int
}

func main() {
	max_answer := answer{
		f: -1000000,
	}
	for x1 := 0; x1 <= 4; x1++ {
		for x2 := 0; x2 <= 3; x2++ {
			if 3*x1+7*x2 <= 21 && x1+x2 <= 4 {
				f := F(x1, x2)
				if f > max_answer.f {
					max_answer.f = f
					max_answer.x1 = x1
					max_answer.x2 = x2
				}
			}
		}
	}
	fmt.Printf("Максимальное значение %d при х1=%d, х2=%d", max_answer.f, max_answer.x1, max_answer.x2)

}
