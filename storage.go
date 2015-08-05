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

	if strings.HasSuffix(path, ".hgt") {
		metaDataFile := path + ".xml"
		if _, err := os.Stat(metaDataFile); err == nil {
			srtmTile := SrtmTile{}
			srtmTile.BoundingRectangle, err = readBoundingRectangle(metaDataFile)
			check(err)
			srtmTile.CompressedData, err = ioutil.ReadFile(path)
			check(err)
			upperLeft, lowerRight := calculateUpperLeftAndLowerRightLikeGdalDataSet(srtmTile.BoundingRectangle)
			srtmTileCollector.tree.Add(srtmTile, upperLeft)
			srtmTileCollector.tree.Add(srtmTile, lowerRight)
			srtmTileCollector.loadedDataSize += int64(len(srtmTile.CompressedData))
			fmt.Printf("Visited: %s\n", path)
			fmt.Printf(">> loadedDataSize: %d\n", srtmTileCollector.loadedDataSize)
	}
}
return nil
}


func (collector SrtmTileCollector) init() {
	collector.loadedDataSize = 0
	upperLeft := quadtree.Twof{+90, -180}
	lowerRight := quadtree.Twof{-90, +180}
	collector.tree = quadtree.MakeQuadtree(upperLeft, lowerRight)
}

func loadHgtFilesIntoStorage(basePath string) {
	srtmTileCollector.init()
	err := filepath.Walk(basePath, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}