package main
import (
	"bytes"
	"image"
	"image/png"
	"os"
)
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}
func main() {
	const (
		dx = 256
		dy = 256
	)
	data := Pic(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create("/tmp/img.png")
	file.Write(buf.Bytes())
	file.Close()
}

