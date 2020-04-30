package speedtest

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

func TestParseServerlistResponse(t *testing.T) {

	f := openSampleFile("samples/speedtest_servers_response.xml", t)
	s, err := parseServerlistResponse(f)
	f.Close()
	if err != nil {
		t.Errorf("Failed to read struct from sample file: \n%s", err.Error())
	}

	if len(s.Servers.ServerList) != 10 {
		t.Errorf("We must have 2 servers, but we have %d", len(s.Servers.ServerList))
	}

	s0 := s.Servers.ServerList[0]
	s1 := s.Servers.ServerList[1]

	if (s0.ID != "22085") && s1.ID != "22085" {
		t.Errorf("We must have at leas one server with ID 22085, but we have %s and %s", s0.ID, s1.ID)
	}
	if (s0.ID != "22448") && s1.ID != "22448" {
		t.Errorf("We must have at leas one server with ID 22448, but we have %s and %s", s0.ID, s1.ID)
	}
}

func TestDistanceFrom(t *testing.T) {
	f := openSampleFile("samples/speedtest_servers_response.xml", t)
	s, err := parseServerlistResponse(f)
	f.Close()
	if err != nil {
		t.Errorf("Failed to read struct from sample file: \n%s", err.Error())
	}

	serverList := s.Servers.ServerList

	testLocation := NewLatLon(-22.9201, -43.3307)
	for _, s := range serverList {
		fmt.Println(s.ID, ") distance:", testLocation.DistanceFrom(s.GetLocation()))
		switch s.ID {
		case "22085":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.13102946233576745
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "22448":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "27842":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "9316":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "8998":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "30423":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "30610":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "30525":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.1356094760700771
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "26318":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.14154073618573917
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		case "32447":
			distance := testLocation.DistanceFrom(s.GetLocation())
			expectedValue := 0.14154073618573917
			if distance != expectedValue {
				t.Errorf("Failed to validate result for ID: %s (%f != %f)", s.ID, expectedValue, distance)
			}
		}
	}
}
