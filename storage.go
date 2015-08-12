package main

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"math"
	"path"
)

type SrtmTile struct {
	name              string
	BoundingRectangle BoundingRectangle
	CompressedData    []byte
}

type SrtmTileCollector struct {
	loadedDataSize int64
	tileIndex      map[string]*SrtmTile
}
var srtmTileCollector = SrtmTileCollector{}


func initSrtmTileCollector() {
	srtmTileCollector.loadedDataSize = 0
	srtmTileCollector.tileIndex = make(map[string]*SrtmTile)
}

func computeTileName(lat float64, lon float64) string {
	var ns string
	var ew string
	if (lat < 0) {
		ns = "S"
	} else {
		ns = "N"
	}
	if (lon <= 0) {
		ew = "E"
	} else {
		ew = "W"
	}
	return fmt.Sprintf("%s%.2d%s%.3d", ns, int(math.Abs(lat)), ew, int(math.Abs(lon)))
}

const LEN_DOT_HGT_DOT_LZ4 = 8
const LEN_DOT_LZ4 = 4

func readSrtmTile(heightDataFile string, metaDataFile string) SrtmTile {
	srtmTile := SrtmTile{}
	var err error
	srtmTile.BoundingRectangle, err = readBoundingRectangle(metaDataFile)
	check(err)
	srtmTile.CompressedData, err = ioutil.ReadFile(heightDataFile)
	check(err)
	srtmTile.name = path.Base(heightDataFile)
	srtmTile.name = srtmTile.name[0:len(srtmTile.name)-LEN_DOT_HGT_DOT_LZ4]
	return srtmTile
}

func visit(pathName string, f os.FileInfo, err error) error {
	if (srtmTileCollector.loadedDataSize / 1000 / 1000) > flagMaxMegaBytes {
		return nil
	}

	if strings.HasSuffix(pathName, ".hgt.lz4") {
		metaDataFile := pathName[0:len(pathName)-LEN_DOT_LZ4] + ".xml"
		if _, err := os.Stat(metaDataFile); err == nil {
			srtmTile := readSrtmTile(pathName, metaDataFile)
			srtmTileCollector.tileIndex[srtmTile.name] = &srtmTile
			srtmTileCollector.loadedDataSize += int64(len(srtmTile.CompressedData))
			fmt.Printf("Visited: %s\n", pathName)
			fmt.Printf(">> loadedDataSize: %d\n", srtmTileCollector.loadedDataSize)
		}
	}
	return nil
}

func loadHgtFilesIntoStorage(basePath string) {
	initSrtmTileCollector()
	err := filepath.Walk(basePath, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	fmt.Printf(">> final loadedDataSize: %d\n", srtmTileCollector.loadedDataSize)
}

func getTile(lat float64, lon float64) {
	name := computeTileName(lat, lon)
	tile := srtmTileCollector.tileIndex[name]
	return tile
}