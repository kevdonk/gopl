/*
Exercise 1.12: Modify the Lissajous server to read parameter values from the URL. For exam- ple, you might arrange it so that
a URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi
function to convert the string parameter into an integer. You can see its documentation with go doc strconv.Atoi.
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w, r.URL.Query())
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xff}, color.RGBA{0xFF, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xFF, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, params map[string][]string) {
	var (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if _, ok := params["cycles"]; ok {
		c, err := strconv.Atoi(params["cycles"][0])
		cycles = c
		if err != nil {
			fmt.Fprintf(out, "invalid cycles")
		}
	}

	if _, ok := params["size"]; ok {
		s, err := strconv.Atoi(params["size"][0])
		size = s
		if err != nil {
			fmt.Fprintf(out, "invalid size")
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(rand.Int()%3+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
