package main

import (
	"github.com/larspensjo/quadtree"
	"bufio"
	"os"
	"math"
	"github.com/bkaradzic/go-lz4"
	"io/ioutil"
	"fmt"
)

const aFile string = "test-data/N50E014.hgt"
const aLz4File string = "test-data/N50E014.hgt.lz4"
const aXmlFile string = "test-data/N50E014.hgt.xml"
const NO_OF_PIXELS_PER_LINE = 1201

func getElevation(lat float64, lon float64) int16 {

	f, err := os.Open(aFile)
	check(err)
	defer f.Close()

	boundingRectangle, err := readBoundingRectangle(aXmlFile)
	check(err)
	upperLeft, lowerRight := calculateUpperLeftAndLowerRightLikeGdalDataSet(boundingRectangle)

	f.Seek(calculateOffset(upperLeft, lowerRight, lat, lon), 0)
	check(err)
	r4 := bufio.NewReader(f)
	heightBuf, err := r4.Peek(2)
	check(err)
	return int16(int(heightBuf[0]) << 8 + int(heightBuf[1]))
}

func getElevationLz4(lat float64, lon float64) int16 {


	boundingRectangle, err := readBoundingRectangle(aXmlFile)
	check(err)
	upperLeft, lowerRight := calculateUpperLeftAndLowerRightLikeGdalDataSet(boundingRectangle)

	offset := calculateOffset(upperLeft, lowerRight, lat, lon)

	compressedData, err := ioutil.ReadFile(aLz4File)
	check(err)

	bufSize := 2 * NO_OF_PIXELS_PER_LINE * NO_OF_PIXELS_PER_LINE
	var dst []byte
	dst = make([]byte, bufSize)
	xdst, err := lz4.Decode(dst, compressedData)
	check(err)

	fmt.Printf("%x\n", dst[0])
	fmt.Printf("%x\n", dst[1])
	fmt.Printf("%x\n", dst[2])
	fmt.Printf("%x\n", dst[3])
	fmt.Printf("%d\n", len(xdst))

	return int16(int(dst[offset]) << 8 + int(dst[offset+1]))
}

func calculateOffset(upperLeft quadtree.Twof, lowerRight quadtree.Twof, lat float64, lon float64) int64 {
	dLat := math.Abs(lat - upperLeft[1])
	dLon := math.Abs(lon - upperLeft[0])
	distLat := math.Abs(upperLeft[1] - lowerRight[1])
	distLon := math.Abs(upperLeft[0] - lowerRight[0])
	nearestLat := int(dLat * float64(NO_OF_PIXELS_PER_LINE) / distLat);
	nearestLon := int(dLon * float64(NO_OF_PIXELS_PER_LINE) / distLon)
	return int64(NO_OF_PIXELS_PER_LINE * nearestLat + nearestLon) << 1
}

func calculateUpperLeftAndLowerRightLikeGdalDataSet(boundingRectangle BoundingRectangle) (quadtree.Twof, quadtree.Twof) {
	upperLeft := quadtree.Twof{
		boundingRectangle.WestBoundingCoordinate + 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1),
		boundingRectangle.NorthBoundingCoordinate - 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1),
	}
	lowerRight := quadtree.Twof{
		boundingRectangle.EastBoundingCoordinate - 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1),
		boundingRectangle.SouthBoundingCoordinate + 0.5 / float64(NO_OF_PIXELS_PER_LINE - 1),
	}
	return upperLeft, lowerRight
}