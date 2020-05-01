package speedtest

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Eldius/speedtest/geolocation"
)

const (
	// OoklaServerListURL is the server list URL
	OoklaServerListURL = "https://c.speedtest.net/speedtest-servers-static.php"
)

/*
ServerListClientConfigWrapper is a server list configuration
*/
type ServerListClientConfigWrapper struct {
	XMLName xml.Name    `xml:"settings"`
	Text    string      `xml:",chardata"`
	Servers ServersList `xml:"servers"`
}

/*
ServersList is a wrapper for the servers list
*/
type ServersList struct {
	Text       string       `xml:",chardata"`
	ServerList []TestServer `xml:"server"`
}

/*
TestServer is the test server definition
*/
type TestServer struct {
	Text    string  `xml:",chardata"`
	URL     string  `xml:"url,attr"`
	Lat     float64 `xml:"lat,attr"`
	Lon     float64 `xml:"lon,attr"`
	Name    string  `xml:"name,attr"`
	Country string  `xml:"country,attr"`
	Cc      string  `xml:"cc,attr"`
	Sponsor string  `xml:"sponsor,attr"`
	ID      string  `xml:"id,attr"`
	Host    string  `xml:"host,attr"`
}

/*
GetLocation returns the geolocation coordinates
of server (lat, lon)
*/
func (t *TestServer) GetLocation() *geolocation.LatLon {
	return geolocation.NewLatLon(t.Lat, t.Lon)
}

/*
FindServers finds some servers from Speedtest
*/
func (c *OoklaClient) FindServers() (servers []TestServer, err error) {
	//servers = make([]ServerSpec, 0)
	res, err := http.Get(OoklaServerListURL)
	if err != nil {
		return servers, err
	}
	defer res.Body.Close()
	return parseServerlistResponse(res.Body)
	//return servers, nil
}

func parseServerlistResponse(r io.ReadCloser) (servers []TestServer, err error) {
	var wrapper ServerListClientConfigWrapper
	configxml, err := ioutil.ReadAll(r)
	if err != nil {
		return servers, err
	}
	err = xml.Unmarshal(configxml, &wrapper)
	if err == nil {
		servers = wrapper.Servers.ServerList
	}
	return
}

/*
FindNearestServers find N nearest servers
*/
func (c *OoklaClient) FindNearestServers(servers []TestServer, p *geolocation.LatLon, q int) []TestServer {
	//c.FetchServers()
	var sortByDistance = func(s1, s2 *TestServer) bool {
		return s1.GetLocation().DistanceFrom(p) < s2.GetLocation().DistanceFrom(p)
	}
	SortServerBy(sortByDistance).SortServer(servers)
	return servers[:q]
}
