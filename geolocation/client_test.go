package geolocation

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func getSampleFilePath(file string, t *testing.T) string {
	currPath, err := os.Getwd()
	if err != nil {
		t.Errorf("Error trying to find current dir")
	}
	fmt.Println("current path", currPath)
	sampleFilePath := filepath.Join(currPath, file)
	fmt.Println("config file", sampleFilePath)

	if _, err := os.Stat(sampleFilePath); err != nil {
		t.Errorf("Config file doesn't exists:\n'%s'", sampleFilePath)
	}

	return sampleFilePath
}

func openSampleFile(file string, t *testing.T) io.ReadCloser {
	if f, err := os.Open(getSampleFilePath(file, t)); err != nil {
		t.Errorf("Filed to open sample file: \n%s", err.Error())
	} else {
		return f
	}
	return nil
}

func TestParseLocationResponse(t *testing.T) {
	f := openSampleFile("samples/geolocation_response.json", t)
	var expectedLat float64 = -22.9201
	var expectedLon float64 = -43.3307
	var expectedLatStr string = "-22.9201"
	var expectedLonStr string = "-43.3307"

	location, err := parseLocationResponse(f)
	if err != nil {
		t.Errorf("Error trying to load sample file: \n%s \n---", err.Error())
	}

	if location.GeopluginLatitude != expectedLatStr {
		t.Errorf("Latitude must be %s, but was %s", expectedLatStr, location.GeopluginLatitude)
	}

	if location.GeopluginLongitude != expectedLonStr {
		t.Errorf("Latitude must be %s, but was %s", expectedLonStr, location.GeopluginLongitude)
	}

	l := location.GetLocation()
	if l.Lat != expectedLat {
		t.Errorf("Latitude must be %f, but was %f", expectedLat, l.Lat)
	}

	if l.Lon != expectedLon {
		t.Errorf("Longitude must be %f, but was %f", expectedLon, l.Lon)
	}

	distance := location.GetLocation().DistanceFromFloat(expectedLat, expectedLon)
	if distance != 0 {
		t.Errorf("Distance must be 0, but was %f", distance)
	}
}
