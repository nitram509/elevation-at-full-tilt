package main

import (
	"testing"
)

const TEST_XML_FILE string = "test-data/N50E014.hgt.xml"

func Test_read_bounding_coordinates_from_file(t *testing.T) {
	boundingRectangle, err := readBoundingRectangle(TEST_XML_FILE)

	if err != nil {
		t.Errorf("the test data should be readable")
	}

	assertEqualsCoordinate(t, "EastBoundingCoordinate", boundingRectangle.EastBoundingCoordinate, 15.00083333)
	assertEqualsCoordinate(t, "WestBoundingCoordinate", boundingRectangle.WestBoundingCoordinate, 13.99916667)
	assertEqualsCoordinate(t, "NorthBoundingCoordinate", boundingRectangle.NorthBoundingCoordinate, 51.00083333)
	assertEqualsCoordinate(t, "SouthBoundingCoordinate", boundingRectangle.SouthBoundingCoordinate, 49.99916667)
}

func assertEqualsCoordinate(t *testing.T, name string, coordinate float64, expected float64) {
	if coordinate != expected {
		t.Errorf("%s: expected = %f --- but was = %f", name, expected, coordinate)
	}
}
