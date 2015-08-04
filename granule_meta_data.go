package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type GranuleMetaData struct {
	BoundingRectangle BoundingRectangle `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle"`
}

type BoundingRectangle struct {
	WestBoundingCoordinate  float64
	NorthBoundingCoordinate float64
	EastBoundingCoordinate  float64
	SouthBoundingCoordinate float64
}

func readBoundingRectangle(fName string) (GranuleMetaData, error) {
	xmlFile, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return GranuleMetaData{}, err
	}

	var granuleMetaData GranuleMetaData
	xml.Unmarshal(xmlFile, &granuleMetaData)

	fmt.Println(granuleMetaData.BoundingRectangle.EastBoundingCoordinate)

	return granuleMetaData, nil
}
