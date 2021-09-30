package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height	= 900, 700
	cells			= 100
	xyrange			= 30.0
	xyscale			= width/2/xyrange
	zscale			= height*2
	angle			= math.Pi/6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main(){
	
	file, err := os.Create("svg.xml")
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close() 
    file.WriteString("<svg xmlns='http://www.w3.org/2000/svg' ")
    file.WriteString("style='stroke: grey; fill: white; stroke-width: 0.7;' ")
    file.WriteString("width='")
    file.WriteString(strconv.Itoa(width))
    file.WriteString("' height='")
    file.WriteString(strconv.Itoa(height))
    file.WriteString("'>")
    
	for i:=0; i < cells; i++ {
		for j:=0; j<cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			
			file.WriteString("<polygon points='")
			file.WriteString(fmt.Sprintf("%f",ax))
			file.WriteString(", ")
			file.WriteString(fmt.Sprintf("%f",ay))
			file.WriteString(" ")
			file.WriteString(fmt.Sprintf("%f",bx))
			file.WriteString(", ")
			file.WriteString(fmt.Sprintf("%f",by))
			file.WriteString(" ")
			file.WriteString(fmt.Sprintf("%f",cx))
			file.WriteString(", ")
			file.WriteString(fmt.Sprintf("%f",cy))
			file.WriteString(" ")
			file.WriteString(fmt.Sprintf("%f",dx))
			file.WriteString(", ")
			file.WriteString(fmt.Sprintf("%f",dy))
			file.WriteString("'/>\n")
			
		}
	}
	file.WriteString("</svg>")
	
	fmt.Println("Done.")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)/r
}