package speedtest

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Eldius/speedtest/geolocation"
)

const (
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

type ServersList struct {
	Text       string       `xml:",chardata"`
	ServerList []TestServer `xml:"server"`
}

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

func (t *TestServer) GetLocation() *geolocation.LatLon {
	return geolocation.NewLatLon(t.Lat, t.Lon)
}

/*
FindServers finds some servers from Speedtest
*/
func (c *OoklaClient) FindServers() (servers ServerListClientConfigWrapper, err error) {
	//servers = make([]ServerSpec, 0)
	res, err := http.Get(OoklaServerListURL)
	if err != nil {
		return servers, err
	}
	defer res.Body.Close()
	return parseServerlistResponse(res.Body)
	//return servers, nil
}

func parseServerlistResponse(r io.ReadCloser) (servers ServerListClientConfigWrapper, err error) {
	configxml, err := ioutil.ReadAll(r)
	if err != nil {
		return servers, err
	}
	err = xml.Unmarshal(configxml, &servers)
	return
}

/*
FindNearestServers find N nearest servers
*/
func (c *OoklaClient) FindNearestServers(q int) (servers []TestServer, err error) {
	//c.FetchServers()
	return nil, nil
}
