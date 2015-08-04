package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type BoundingRectangle struct {
	WestBoundingCoordinate  float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>WestBoundingCoordinate"`
	NorthBoundingCoordinate float64  `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>NorthBoundingCoordinate"`
	EastBoundingCoordinate  float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>EastBoundingCoordinate"`
	SouthBoundingCoordinate float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>SouthBoundingCoordinate"`
}

func readBoundingRectangle(fName string) (BoundingRectangle, error) {
	xmlFile, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return BoundingRectangle{}, err
	}
	var x BoundingRectangle
	xml.Unmarshal(xmlFile, &x)
	return x, nil
}
