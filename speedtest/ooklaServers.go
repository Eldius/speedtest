package speedtest

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
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
	ID      string  `xml:"id,attr" gorm:"type:varchar(100);PRIMARY_KEY"`
	Host    string  `xml:"host,attr"`
}

/*
SelectedServer a workaround to save
the nearest servers (for now)
*/
type SelectedServer struct {
	ID       string
	Server   TestServer
	Distance float64
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
	w, err := c.fetchServersData()
	if err == nil {
		servers = w.Servers.ServerList
	}
	return
}

/*
FindServers finds some servers from Speedtest
*/
func (c *OoklaClient) fetchServersData() (w ServerListClientConfigWrapper, err error) {
	//servers = make([]ServerSpec, 0)
	res, err := http.Get(OoklaServerListURL)
	defer res.Body.Close()
	return parseServerlistResponse(res.Body)
}

func parseServerlistResponse(r io.ReadCloser) (wrapper ServerListClientConfigWrapper, err error) {
	configxml, err := ioutil.ReadAll(r)
	if err != nil {
		return wrapper, err
	}
	err = xml.Unmarshal(configxml, &wrapper)
	if err != nil {
		log.Panic(err.Error())
	}
	return
}

/*
FindNearestServers find N nearest servers
*/
func (c *OoklaClient) FindNearestServers(p *geolocation.LatLon, q int) []TestServer {
	//c.FetchServers()
	if w, err := c.fetchServersData(); err != nil {
		panic(err.Error())
	} else {
		return w.FindNearestServers(p, q)
	}
}

/*
FindNearestServers find N nearest servers
*/
func (c *ServerListClientConfigWrapper) FindNearestServers(p *geolocation.LatLon, q int) []TestServer {
	//c.FetchServers()
	var sortByDistance = func(s1, s2 *TestServer) bool {
		return s1.GetLocation().DistanceFrom(p) < s2.GetLocation().DistanceFrom(p)
	}
	servers := c.Servers.ServerList
	SortServerBy(sortByDistance).SortServer(c.Servers.ServerList)
	return servers[:q]
}

/*
FindFastestServersFromPing find N nearest servers
*/
func (c *OoklaClient) FindFastestServersFromPing(servers []TestServer, p *geolocation.LatLon, q int) TestServer {
	var sortByDistance = func(s1, s2 *TestServer) bool {
		return s1.GetLocation().DistanceFrom(p) < s2.GetLocation().DistanceFrom(p)
	}
	SortServerBy(sortByDistance).SortServer(servers)
	return servers[0]
}

/*
ToSelectedServer parse server to SelectedServer
*/
func (t *TestServer) ToSelectedServer(l geolocation.LatLon) SelectedServer {
	return SelectedServer{
		ID:       t.ID,
		Server:   *t,
		Distance: t.GetLocation().DistanceFrom(&l),
	}
}
