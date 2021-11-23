package euler15

import "fmt"

type dimensions struct {
	width  int
	height int
}

func (d *dimensions) length() int {
	return d.height + d.width
}

type path []bool
type grid []path

func LatticePath(d dimensions) int {

	var p path
	var g grid

	// Create first solution
	for i := 0; i < d.width; i++ {
		p = append(p, true)
	}
	for i := 0; i < d.height; i++ {
		p = append(p, false)
	}

	// Add first solution to ps
	g = append(g, p)

	for true {
		passedDown := false

		if 
	}

	fmt.Printf("%+v\n", g)

	return 0
}
