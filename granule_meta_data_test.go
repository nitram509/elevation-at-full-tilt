package main

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

const TEST_XML_FILE string = "test-data/N50E014.hgt.xml"

func Test_read_bounding_coordinates_from_file(t *testing.T) {
	boundingRectangle, err := readBoundingRectangle(TEST_XML_FILE)
	then.AssertThat(t, err, is.Nil())

	then.AssertThat(t, boundingRectangle.EastBoundingCoordinate, is.EqualTo(15.00083333))
	then.AssertThat(t, boundingRectangle.WestBoundingCoordinate, is.EqualTo(13.99916667))
	then.AssertThat(t, boundingRectangle.NorthBoundingCoordinate, is.EqualTo(51.00083333))
	then.AssertThat(t, boundingRectangle.SouthBoundingCoordinate, is.EqualTo(49.99916667))
}
