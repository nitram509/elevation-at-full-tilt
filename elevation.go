package main

import (
	"github.com/larspensjo/quadtree"
	"bufio"
	"os"
	"math"
)

const aFile string = "test-data/N50E014.hgt"
const NO_OF_PIXELS_PER_LINE = 1201

func getElevation(lat float64, lon float64) int16 {

	f, err := os.Open(aFile)
	check(err)

	upperLeft, lowerRight := calculateUpperLeftAndLowerRightLikeGdalDataSet()

	dLat := math.Abs(lat - upperLeft[1])
	dLon := math.Abs(lon - upperLeft[0])
	distLat := math.Abs(upperLeft[1] - lowerRight[1])
	distLon := math.Abs(upperLeft[0] - lowerRight[0])
	nearestLat := int(dLat * float64(NO_OF_PIXELS_PER_LINE) / distLat);
	nearestLon := int(dLon * float64(NO_OF_PIXELS_PER_LINE) / distLon)
	offset := int64(NO_OF_PIXELS_PER_LINE * nearestLat + nearestLon) << 1
	f.Seek(offset, 0)
	check(err)
	r4 := bufio.NewReader(f)
	heightBuf, err := r4.Peek(2)
	check(err)
	return int16(int(heightBuf[0]) << 8 + int(heightBuf[1]))
}

func calculateUpperLeftAndLowerRightLikeGdalDataSet() (quadtree.Twof, quadtree.Twof) {
	wbc := 13.99916667
	nbc := 51.00083333
	ebc := 15.00083333
	sbc := 49.99916667

	upperLeft := quadtree.Twof{wbc + 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1), nbc - 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1)}
	lowerRight := quadtree.Twof{ebc - 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1), sbc + 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1)}

	return upperLeft, lowerRight
}