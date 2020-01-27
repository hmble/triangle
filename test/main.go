package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	s := fmt.Sprintf("[%d, %d]", p.X, p.Y)

	return s
}
func main() {
	// [0,0] [0,1]
	// [1,0] [1,1]
	// [0,2] [0,3]
	// [1,2] [1,3]
	// [2,0] [2,1]
	// [3,0] [3,1]
	// [2,2] [2,3]
	// [3,2] [3,3]
	var offx, offy int

	pointsArray := make([][]Point, 0)
	boardWidth := 4

	tilesize := 2
	countsize := (boardWidth * boardWidth) / (tilesize * tilesize)
	for count := 0; count < countsize; count++ {

		if offy+tilesize == boardWidth {
			offx += tilesize
			offy = 0

		} else {
			//			offx = 0

			if count != 0 {
				offy += tilesize
			}
		}

		temparray := make([]Point, 0)
		for x := 0; x < tilesize; x++ {
			for y := 0; y < tilesize; y++ {
				//	fmt.Printf("[%d, %d] => %d \t", x+offx, y+offy, twodArray[x+offx][y+offy])
				temparray = append(temparray, Point{x + offx, y + offy})
			}
			//fmt.Println()
		}
		pointsArray = append(pointsArray, temparray)
		//fmt.Println()

	}

	for _, i := range pointsArray {

		fmt.Printf("Length of parent %d\n", len(i))
		for _, j := range i {
			fmt.Println(j.String())
		}
	}

}

func printPoint(x, y int) {
	fmt.Printf("[%d, %d]\t", x, y)
}
