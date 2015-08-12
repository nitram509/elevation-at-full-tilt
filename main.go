package main

import (
	"gopkg.in/alecthomas/kingpin.v1"
)

const (
	VERSION = "0.0.1"
	DEFAULT_MAX_MBYTES = 100
)

var (
	flagMaxMegaBytes int64= int64(DEFAULT_MAX_MBYTES) //kingpin.Flag("maxMegaBytes", "Max MegaBytes").Short('m').Default(strconv.Itoa(DEFAULT_MAX_MBYTES)).Int64()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	kingpin.Version(VERSION)
	kingpin.Parse()

	getElevation(50.918961, 14.057732)
	getElevation(50.851495, 14.301564)
	getElevation(50.4163577778, 14.9198269444)
}
