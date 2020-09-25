// Implementation of Slices exercise from Golang's official tutorial 
// https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8,dy)
	b := make([]uint8, dx)
	
	for i := range a {
		for j := range b {
			b[j] = uint8((i^j)*(i^j))
		}
		a[i] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}