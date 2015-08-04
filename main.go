package main

import (
	"fmt"
	"github.com/larspensjo/quadtree"
	"gopkg.in/alecthomas/kingpin.v1"
	"strconv"
)

const (
	VERSION = "0.0.1"
	DEFAULT_MAX_MBYTES = 1000
	DEFAULT_NO_PATH = "<<no-path>>"
)

var (
	flagMaxMegaBytes = kingpin.Flag("maxMegaBytes", "Max MegaBytes").Short('m').Default(strconv.Itoa(DEFAULT_MAX_MBYTES)).Int64()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type SrtmTile struct {
	quadtree.Handle
	// Add other attributes here
}

func printSrtmTiles() {
	upperLeft := quadtree.Twof{0, 0}
	lowerRight := quadtree.Twof{1, 1}
	tree := quadtree.MakeQuadtree(upperLeft, lowerRight)
	// Create 10 balls and add them to the quadtree
	for i := 0; i < 10; i++ {
		var b SrtmTile
		tree.Add(&b, quadtree.Twof{float64(i) / 10.0, 0})
	}
	list := tree.FindNearObjects(quadtree.Twof{0.5, 0.1}, 0.2)
	fmt.Println("Found", len(list))
}

func main() {

	kingpin.Version(VERSION)
	kingpin.Parse()

	getElevation(50.918961, 14.057732)
	getElevation(50.851495, 14.301564)
	getElevation(50.4163577778, 14.9198269444)
}
