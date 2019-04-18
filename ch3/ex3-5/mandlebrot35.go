// Exercise 3.5: Implement a full-color Mandelbrot set using the function image.NewRGBA and the type color.RGBA or color.YCbCr.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// this actually just does random colors.. meaningful colors are probably the way to go
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	var palette = []color.Color{
		color.Black,
		color.RGBA{0x00, 0xFF, 0x00, 0xff},
		color.RGBA{0xFF, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xFF, 0xff},
		color.RGBA{0xDD, 0x33, 0x11, 0xff},
		color.RGBA{0x88, 0xC1, 0x00, 0xff},
		color.RGBA{0x11, 0x33, 0xDD, 0xff},
		color.RGBA{0x11, 0xDD, 0x33, 0xff},
		color.RGBA{0x33, 0xDD, 0x11, 0xff},
	}
	return palette[rand.Int()%8+1]
}
