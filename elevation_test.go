package main

import (
	"testing"
)

const TEST_FILE string = "test-data/N50E014.hgt"

func Test_resolve_elevation_for_well_known_points(t *testing.T) {
	cases := []struct {
		lat               float64;
		lon               float64;
		expectedElevation int16;
	}{
		{50.918961, 14.057732, 355},
		{50.851495, 14.301564, 305},
		{50.4163577778, 14.9198269444, 216},
	}
	for _, c := range cases {
		actualElevation := getElevation(c.lat, c.lon)
		if actualElevation != c.expectedElevation {
			t.Errorf("Lat=%f, Lon=%f, expected elevation=%d --- but actual elevation=%d", c.lat, c.lon, c.expectedElevation, actualElevation)
		}
	}
}

func Test_resolve_elevation_for_well_known_points_with_LZ4(t *testing.T) {
	cases := []struct {
		lat               float64;
		lon               float64;
		expectedElevation int16;
	}{
		{50.918961, 14.057732, 355},
		{50.851495, 14.301564, 305},
		{50.4163577778, 14.9198269444, 216},
	}
	for _, c := range cases {
		actualElevation := getElevationLz4(c.lat, c.lon)
		if actualElevation != c.expectedElevation {
			t.Errorf("Lat=%f, Lon=%f, expected elevation=%d --- but actual elevation=%d", c.lat, c.lon, c.expectedElevation, actualElevation)
		}
	}
}