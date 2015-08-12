package main

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"github.com/larspensjo/quadtree"
	"io/ioutil"
)

type SrtmTile struct {
	quadtree.Handle
	BoundingRectangle BoundingRectangle
	CompressedData    []byte
}

type SrtmTileCollector struct {
	loadedDataSize int64
	tree           *quadtree.Quadtree
}
var srtmTileCollector = SrtmTileCollector{}


func visit(path string, f os.FileInfo, err error) error {
	if (srtmTileCollector.loadedDataSize / 1000 / 1000) > flagMaxMegaBytes {
		return nil
	}

	if strings.HasSuffix(path, ".hgt.lz4") {
		metaDataFile := path[0:len(path)-4] + ".xml"
		if _, err := os.Stat(metaDataFile); err == nil {
			srtmTile := SrtmTile{}
			srtmTile.BoundingRectangle, err = readBoundingRectangle(metaDataFile)
			check(err)
			srtmTile.CompressedData, err = ioutil.ReadFile(path)
			check(err)
			upperLeft, lowerRight := calculateUpperLeftAndLowerRightLikeGdalDataSet(srtmTile.BoundingRectangle)
			srtmTileCollector.tree.Add(&srtmTile, upperLeft)
			srtmTileCollector.tree.Add(&srtmTile, lowerRight)
			srtmTileCollector.loadedDataSize += int64(len(srtmTile.CompressedData))
			fmt.Printf("Visited: %s\n", path)
			fmt.Printf(">> loadedDataSize: %d\n", srtmTileCollector.loadedDataSize)
		}
	}
	return nil
}


func initSrtmTileCollector() {
	srtmTileCollector.loadedDataSize = 0
	upperLeft := quadtree.Twof{+90, -180}
	lowerRight := quadtree.Twof{-90, +180}
	srtmTileCollector.tree = quadtree.MakeQuadtree(upperLeft, lowerRight)
}

func loadHgtFilesIntoStorage(basePath string) {
	initSrtmTileCollector()
	err := filepath.Walk(basePath, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	fmt.Printf(">> final loadedDataSize: %d\n", srtmTileCollector.loadedDataSize)
}