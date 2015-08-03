package main

import (
	"fmt"
	"github.com/larspensjo/quadtree"
	"bufio"
	"os"
	"math"
)

const aFile string = "test-data/N50E014.hgt"
const NO_OF_LINES = float64(1201)

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
	// Output: Found 3
}

func getElevation(lat float64, lon float64) int16 {
	fmt.Printf("Lat/Lon: %f, %f\n", lat,lon)
	f, err := os.Open(aFile)
	check(err)

	/*
			<BoundingRectangle>
				<WestBoundingCoordinate>13.99916667</WestBoundingCoordinate>
				<NorthBoundingCoordinate>51.00083333</NorthBoundingCoordinate>
				<EastBoundingCoordinate>15.00083333</EastBoundingCoordinate>
				<SouthBoundingCoordinate>49.99916667</SouthBoundingCoordinate>
			</BoundingRectangle>
*/
	upperLeft := quadtree.Twof{13.99916667, 51.00083333}
	lowerRight := quadtree.Twof{15.00083333, 49.99916667}
	dLat := math.Abs(lat - upperLeft[1])
	dLon := math.Abs(lon - upperLeft[0])
	distLat := math.Abs(upperLeft[1] - lowerRight[1])
	distLon := math.Abs(upperLeft[0] - lowerRight[0])
	offset := int64((dLat * float64(NO_OF_LINES) / distLat) + NO_OF_LINES * dLon * float64(NO_OF_LINES) / distLon) << 1
	f.Seek(offset, 0)
	check(err)
	r4 := bufio.NewReader(f)
	heightBuf, err := r4.Peek(2)
	check(err)
	fmt.Printf("Offset: %d\n", offset)
	fmt.Printf("Height-dez: %d\n", int(int(heightBuf[0]) << 8 + int(heightBuf[1])))
	fmt.Printf("Height-hex: %.2x%.2x\n", heightBuf[0], heightBuf[1])

	return int16(int(heightBuf[0]) << 8 + int(heightBuf[1]))
}

func main() {
	getElevation(50.918961, 14.057732)
	getElevation(50.851495, 14.301564)
	getElevation(50.4163577778, 14.9198269444)
}
