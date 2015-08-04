package main

import (
	"fmt"
	"github.com/larspensjo/quadtree"
	"bufio"
	"os"
	"math"
)

const aFile string = "test-data/N50E014.hgt"
const NO_OF_LINES = int32(1201)

func getElevation(lat float64, lon float64) int16 {
	fmt.Printf("Lat/Lon: %f, %f\n", lat, lon)
	f, err := os.Open(aFile)
	check(err)

	upperLeft := quadtree.Twof{13.9995833, 51.0004167}
	lowerRight := quadtree.Twof{15.0004167, 49.9995833}

	dLat := math.Abs(lat - upperLeft[1])
	dLon := math.Abs(lon - upperLeft[0])
	distLat := math.Abs(upperLeft[1] - lowerRight[1])
	distLon := math.Abs(upperLeft[0] - lowerRight[0])
	nearestLat := int32(dLat * float64(NO_OF_LINES) / distLat);
	nearestLon := int32(dLon * float64(NO_OF_LINES) / distLon)
	offset := int64(NO_OF_LINES * nearestLat + nearestLon) << 1
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
