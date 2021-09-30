// print args
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		c, err := strconv.ParseFloat(r.URL.Query().Get("c"), 64)
		if err != nil {
			c = 1
		}
		lissajous(w, c)
	}
	http.HandleFunc("/", handler)
	
	http.HandleFunc("/count", counter)
	http.HandleFunc("/discr", discr)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}



func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func discr(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		a = 0
	}
	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}
	c, err := strconv.ParseFloat(r.URL.Query().Get("c"), 64)
	if err != nil {
		c = 0
	}
	d := b*b - 4*a*c
	if d<0 {
		fmt.Print("D < 0\n")
	}
	x1 := (-1*b - math.Pow(d, 0.5))/(2*a)
	x2 := (-1*b + math.Pow(d, 0.5))/(2*a)
	
	fmt.Fprintf("x1 = %.03f\nx2 = %.03f", x1, x2)
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, c float64) {
	const (
		res = 0.001
		size = 200
		nframes = 64
		delay = 8
	)
	cycles := c
	
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
} 