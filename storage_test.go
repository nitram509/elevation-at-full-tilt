package main

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

const BASE_PATH = "/home/maki/srtm-data/lz4-hc"

func Test_a_tile_name_is_computed_from_lat_lon(t *testing.T) {
	cases := []struct {
		lat          float64
		lon          float64
		expectedName string
	}{
		{50.918961, 14.057732, "N50W014"},
		{50.851495, -14.301564, "N50E014"},
		{-50.4163577778, 0, "S50E000"},
		{0, 13.13, "N00W013"},
	}
	for _, c := range cases {
		name := computeTileName(c.lat, c.lon)
		then.AssertThat(t, name, is.EqualTo(c.expectedName))
	}
}

func Test_loading_files_into_storage(t *testing.T) {

	loadHgtFilesIntoStorage(BASE_PATH)
	//		actualElevation := getElevation(c.lat, c.lon)
	//		if actualElevation != c.expectedElevation {
	//			t.Errorf("Lat=%f, Lon=%f, expected elevation=%d --- but actual elevation=%d", c.lat, c.lon, c.expectedElevation, actualElevation)
	//		}

}
