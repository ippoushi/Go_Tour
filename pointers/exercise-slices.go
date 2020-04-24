package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	images := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
        	for x := 0; x < dx; x++ {
			v := (x + y) / 2
			s[x] = uint8(v)
	        }
	        images[y] = s
	}
	return images
}

func main() {
	pic.Show(Pic)
}

