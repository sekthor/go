package main

import (
	"fmt"

	"github.com/paulmach/orb"
)

func main() {
	p1 := orb.Point{48.0000000, 6.0000000}
	p2 := orb.Point{46.0000000, 9.0000000}

	bound := orb.MultiPoint{p1, p2}.Bound()

	bern := orb.Point{46.9480, 7.4474}
	berlin := orb.Point{52.518611111111, 13.40833333333334}

	fmt.Println(bound.Contains(bern))
	fmt.Println(bound.Contains(berlin))
}
