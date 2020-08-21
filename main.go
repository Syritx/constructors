package main

import (
	"fmt"
	"strconv"
)

type Tile struct {
	x int
	y int
}

// CONTRUCTOR
func (tile *Tile) Init(x int, y int) {
	tile.x = x
	tile.y = y
}

func main()  {
	tile := new(Tile)
	tile.Init(1,1)

	fmt.Printf(strconv.Itoa(tile.x) + "X\n")
	fmt.Printf(strconv.Itoa(tile.y) + "Y\n")
}
