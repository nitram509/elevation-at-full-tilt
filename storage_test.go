package main

import (
	"testing"
)

const BASE_PATH = "d:\\srtm-data\\http__e4ftl01.cr.usgs.gov__SRTM__SRTMGL3.003__2000.02.11"

func Test_loading_files_into_storage(t *testing.T) {

	loadHgtFilesIntoStorage(BASE_PATH)
//		actualElevation := getElevation(c.lat, c.lon)
	//		if actualElevation != c.expectedElevation {
	//			t.Errorf("Lat=%f, Lon=%f, expected elevation=%d --- but actual elevation=%d", c.lat, c.lon, c.expectedElevation, actualElevation)
	//		}

}
