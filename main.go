package main

import (
	"fmt"
	"github.com/larspensjo/quadtree"
)

type ball struct {
	quadtree.Handle
	// Add other attributes here
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExampleBalls() {
	upperLeft := quadtree.Twof{0, 0}
	lowerRight := quadtree.Twof{1, 1}
	tree := quadtree.MakeQuadtree(upperLeft, lowerRight)
	// Create 10 balls and add them to the quadtree
	for i := 0; i < 10; i++ {
		var b ball
		tree.Add(&b, quadtree.Twof{float64(i) / 10.0, 0})
	}
	list := tree.FindNearObjects(quadtree.Twof{0.5, 0.1}, 0.2)
	fmt.Println("Found", len(list))
}

func main() {
	getElevation(50.918961, 14.057732)
	getElevation(50.851495, 14.301564)
	getElevation(50.4163577778, 14.9198269444)
}
