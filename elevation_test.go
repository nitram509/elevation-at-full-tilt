package main

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func Test_resolve_elevation_for_well_known_points(t *testing.T) {
	cases := []struct {
		lat               float64
		lon               float64
		expectedElevation int16
	}{
		{50.918961, 14.057732, 355},
		{50.851495, 14.301564, 305},
		{50.4163577778, 14.9198269444, 216},
	}
	for _, c := range cases {
		elevation := getElevation(c.lat, c.lon)
		then.AssertThat(t, elevation, is.EqualTo(c.expectedElevation))
	}
}

func Test_resolve_elevation_for_well_known_points_with_LZ4(t *testing.T) {
	cases := []struct {
		lat               float64
		lon               float64
		expectedElevation int16
	}{
		{50.918961, 14.057732, 355},
		{50.851495, 14.301564, 305},
		{50.4163577778, 14.9198269444, 216},
	}
	for _, c := range cases {
		elevation := getElevationLz4(c.lat, c.lon)
		then.AssertThat(t, elevation, is.EqualTo(c.expectedElevation))
	}
}
