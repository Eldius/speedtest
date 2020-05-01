package speedtest

import (
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/Eldius/speedtest/geolocation"
)

const (
	distanceTestTolerance = 0.000000001
)

func getSampleFilePath(file string, t *testing.T) string {
	currPath, err := os.Getwd()
	if err != nil {
		t.Errorf("Error trying to find current dir")
	}
	log.Println("current path", currPath)
	sampleFilePath := filepath.Join(currPath, file)
	log.Println("config file", sampleFilePath)

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

func TestParseServerlistResponse(t *testing.T) {

	f := openSampleFile("samples/speedtest_servers_response.xml", t)
	w, err := parseServerlistResponse(f)
	serverList := w.Servers.ServerList
	f.Close()
	if err != nil {
		t.Errorf("Failed to read struct from sample file: \n%s", err.Error())
	}

	if len(serverList) != 10 {
		t.Errorf("We must have 2 servers, but we have %d", len(serverList))
	}

	s0 := serverList[0]
	s1 := serverList[1]

	if (s0.ID != "22085") && s1.ID != "22085" {
		t.Errorf("We must have at leas one server with ID 22085, but we have %s and %s", s0.ID, s1.ID)
	}
	if (s0.ID != "22448") && s1.ID != "22448" {
		t.Errorf("We must have at leas one server with ID 22448, but we have %s and %s", s0.ID, s1.ID)
	}
}

func TestDistanceFrom(t *testing.T) {
	f := openSampleFile("samples/speedtest_servers_response.xml", t)
	w, err := parseServerlistResponse(f)
	serverList := w.Servers.ServerList
	f.Close()
	if err != nil {
		t.Errorf("Failed to read struct from sample file: \n%s", err.Error())
	}

	testLocation := geolocation.NewLatLon(-22.9201, -43.3307)
	for _, s := range serverList {
		log.Println(s.ID, ") distance:", testLocation.DistanceFrom(s.GetLocation()))
		switch s.ID {
		case "22085":
			expectedValue := 13.4429193589199
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "22448":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "27842":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "9316":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "8998":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "30423":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "30610":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "30525":
			expectedValue := 15.0740816039992
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "26318":
			expectedValue := 15.2819887959137
			validateDistanceCalc(s, testLocation, expectedValue, t)
		case "32447":
			expectedValue := 15.2819887959137
			validateDistanceCalc(s, testLocation, expectedValue, t)
		}
	}
}

func TestFindNearestServers(t *testing.T) {
	f := openSampleFile("samples/speedtest_servers_response.xml", t)
	w, err := parseServerlistResponse(f)
	serverList := w.Servers.ServerList
	f.Close()
	if err != nil {
		t.Errorf("Failed to read struct from sample file: \n%s", err.Error())
	}

	testLocation := geolocation.NewLatLon(-22.9201, -43.3307)

	nearestServers := w.FindNearestServers(testLocation, 2)

	log.Println("---")
	log.Println("ID, lat, lon, test lat, test lon, distance")
	for _, s := range nearestServers {
		log.Println(s.ID, ",", s.GetLocation().Lat, ",", s.GetLocation().Lon, ",", testLocation.Lat, ",", testLocation.Lon, ",", s.GetLocation().DistanceFrom(testLocation))
	}
	log.Println("---")
	log.Println("---")
	log.Println("ID, lat, lon, test lat, test lon, distance")
	for _, s := range serverList {
		log.Println(s.ID, ",", s.GetLocation().Lat, ",", s.GetLocation().Lon, ",", testLocation.Lat, ",", testLocation.Lon, ",", s.GetLocation().DistanceFrom(testLocation))
	}
	log.Println("---")

	if len(nearestServers) != 2 {
		t.Errorf("Must have 2 servers, but has %d", len(nearestServers))
	}

	if nearestServers[0].ID != "22085" {
		t.Errorf("Nearest server must be ID 22085, but is %s", nearestServers[0].ID)
	}
	if nearestServers[1].ID != "22448" {
		t.Errorf("Nearest server must be ID 22448, but is %s", nearestServers[1].ID)
	}
}

func validateDistanceCalc(s TestServer, location *geolocation.LatLon, expectedDistance float64, t *testing.T) {
	distance := s.GetLocation().DistanceFrom(location)
	if math.Abs(expectedDistance-distance) > distanceTestTolerance {
		t.Errorf("Failed to validate result for ID: %s (%f != %f/%f)", s.ID, expectedDistance, distance, math.Abs(expectedDistance-distance))
	}
}
