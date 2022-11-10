package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type BoundingRectangle struct {
	WestBoundingCoordinate  float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>WestBoundingCoordinate"`
	NorthBoundingCoordinate float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>NorthBoundingCoordinate"`
	EastBoundingCoordinate  float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>EastBoundingCoordinate"`
	SouthBoundingCoordinate float64 `xml:"GranuleURMetaData>SpatialDomainContainer>HorizontalSpatialDomainContainer>BoundingRectangle>SouthBoundingCoordinate"`
}

func readBoundingRectangle(fName string) (BoundingRectangle, error) {
	xmlFile, err := os.ReadFile(fName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return BoundingRectangle{}, err
	}
	var x BoundingRectangle
	xml.Unmarshal(xmlFile, &x)
	return x, nil
}
