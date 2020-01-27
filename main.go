package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"os"
)

// type RGB struct {
// 	R uint32
// 	G uint32
// 	B uint32
// 	A uint32
// }

type RGB struct {
	R, G, B, A int64
}

func (r RGB) String() string {
	s := fmt.Sprintf("[%d, %d, %d, %d]", r.R, r.G, r.B, r.A)

	return s
}
func averageColor(colors []RGB) RGB {
	n := len(colors)

	var r, g, b, a int64
	for _, c := range colors {

		r += (c.R * c.R)
		g += (c.G * c.G)
		b += (c.B * c.B)
		a += (c.A * c.A)

	}

	return RGB{
		R: int64(math.Sqrt(float64(r / int64(n)))),
		G: int64(math.Sqrt(float64(g / int64(n)))),
		B: int64(math.Sqrt(float64(b / int64(n)))),
		A: int64(math.Sqrt(float64(a / int64(n)))),
	}

}

type Pixel struct {
	Point    image.Point
	Color    RGB
	OwnColor RGB
}

func (p Pixel) String() string {
	s := fmt.Sprintf("Point : [%d, %d], RGB: [%d, %d, %d], OC : [%d, %d, %d]",
		p.Point.X, p.Point.Y,
		p.Color.R, p.Color.G, p.Color.B,
		p.OwnColor.R, p.OwnColor.G, p.OwnColor.B,
	)

	return s
}

type ImageSet interface {
	Set(x, y int, c color.Color)
}

// Get tiles array of group to make tiles and get average color
func getTilesArray(tilesize, borderWidth int, img image.Image) [][]Pixel {
	pointsArray := make([][]Pixel, 0)
	countsize := (borderWidth * borderWidth) / (tilesize * tilesize)

	var offx, offy int

	for count := 0; count < countsize; count++ {
		if offy+tilesize == borderWidth {
			offx += tilesize
			offy = 0

		} else {

			if count != 0 {
				offy += tilesize
			}
		}

		temparray := make([]Pixel, 0)
		//	colorsArray := make([]RGB, 0)
		point := Pixel{}
		for x := 0; x < tilesize; x++ {
			for y := 0; y < tilesize; y++ {
				point.Point = image.Point{x + offx, y + offy}

				r, g, b, a := img.At(x+offx, y+offy).RGBA()

				rgba := RGB{
					R: int64(r >> 8),
					G: int64(g >> 8),
					B: int64(b >> 8),
					A: int64(a >> 8),
				}

				point.OwnColor = rgba
				//				colorsArray = append(colorsArray, rgba)
				temparray = append(temparray, point)
			}
		}

		// 		avgColor := averageColor(colorsArray)
		// 		setColor(temparray, avgColor)
		pointsArray = append(pointsArray, temparray)

	}

	return pointsArray
}

func setColor(colors []*Pixel, color RGB) {
	for _, p := range colors {
		p.Color = color
	}

}
func main() {

	// 600x600 image path
	twoCopy2 := "./assets/two_copy2.jpg"

	reader, readErr := os.Open(twoCopy2)

	if readErr != nil {
		log.Fatal("Error in opening image")
	}
	defer reader.Close()

	// // Decode image

	img, _, imgErr := image.Decode(reader)

	if imgErr != nil {
		log.Fatal("Error in decoding image")
	}

	//bounds := img.Bounds()

	//fmt.Println(bounds.Max.X, bounds.Max.Y)

	tilesArray := getTilesArray(2, 600, img)

	fmt.Println("Tilesarray length", len(tilesArray))

	// tileArrayCopy := &tilesArray

	// for i, tiles := range tilesArray {
	// for j, tile := range tiles {

	// }
	// }

	r, g, b, _ := img.At(599, 599).RGBA()
	fmt.Printf("[%d, %d, %d]", int64(r>>8), int64(g>>8), int64(b>>8))
}
