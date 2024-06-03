package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {

	var n int
	fmt.Println("Сколько бросков требуется выполнить?")
	fmt.Scan(&n)
	img := loadImage("/home/xapsiel/Desktop/Горская/лабы/lab1/BlackSea/BlackSea.png")
	k := 010
	all := 0
	rgbaimg, _ := img.(*image.RGBA)
	if 1118*604 < n {
		n = 1118 * 604
	}
	for i := 0; i < n; i++ {
		x := rand.Intn(1118)
		y := rand.Intn(604)
		left := img.At(x, y)
		right := color.Color(color.RGBA{124, 213, 233, 255})
		if left == right {
			if rgbaimg.At(x, y) != color.Color(color.RGBA{100, 100, 100, 1}) {
				k++
				rgbaimg.Set(x, y, color.RGBA{100, 100, 100, 1})
				all++

			}

		} else {
			if rgbaimg.At(x, y) != color.Color(color.RGBA{200, 200, 200, 1}) {
				rgbaimg.Set(x, y, color.RGBA{200, 200, 200, 1})
				all++

			}
		}

	}
	file, _ := os.Create("output.png")
	png.Encode(file, rgbaimg)
	fmt.Println((float64(k) / float64(all)) * 1150.0 * 620.0)

}

func loadImage(filepath string) image.Image {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println((*file).Name())
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}
