package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// make 2d slices
	p := make([][]uint8, dy)
	for i := range p {
		// make dx slices
		p[i] = make([]uint8, dx)
	}

	// fill with numbers
	for y, row := range p {
		for x := range row {
			row[x] = uint8(x^y)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
