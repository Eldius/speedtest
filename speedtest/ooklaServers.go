package speedtest

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"math"
	"net/http"
)

const (
	OoklaServerListURL = "https://c.speedtest.net/speedtest-servers-static.php"
)

type LatLon struct {
	Lat float64
	Lon float64
}

type IPGeolocationResponse struct {
	GeopluginRequest                string  `json:"geoplugin_request"`
	GeopluginStatus                 int     `json:"geoplugin_status"`
	GeopluginDelay                  string  `json:"geoplugin_delay"`
	GeopluginCredit                 string  `json:"geoplugin_credit"`
	GeopluginCity                   string  `json:"geoplugin_city"`
	GeopluginRegion                 string  `json:"geoplugin_region"`
	GeopluginRegionCode             string  `json:"geoplugin_regionCode"`
	GeopluginRegionName             string  `json:"geoplugin_regionName"`
	GeopluginAreaCode               string  `json:"geoplugin_areaCode"`
	GeopluginDmaCode                string  `json:"geoplugin_dmaCode"`
	GeopluginCountryCode            string  `json:"geoplugin_countryCode"`
	GeopluginCountryName            string  `json:"geoplugin_countryName"`
	GeopluginInEU                   int     `json:"geoplugin_inEU"`
	GeopluginEuVATrate              bool    `json:"geoplugin_euVATrate"`
	GeopluginContinentCode          string  `json:"geoplugin_continentCode"`
	GeopluginContinentName          string  `json:"geoplugin_continentName"`
	GeopluginLatitude               float64 `json:"geoplugin_latitude"`
	GeopluginLongitude              float64 `json:"geoplugin_longitude"`
	GeopluginLocationAccuracyRadius string  `json:"geoplugin_locationAccuracyRadius"`
	GeopluginTimezone               string  `json:"geoplugin_timezone"`
	GeopluginCurrencyCode           string  `json:"geoplugin_currencyCode"`
	GeopluginCurrencySymbol         string  `json:"geoplugin_currencySymbol"`
	GeopluginCurrencySymbolUTF8     string  `json:"geoplugin_currencySymbol_UTF8"`
	GeopluginCurrencyConverter      float64 `json:"geoplugin_currencyConverter"`
}

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

func (t *TestServer) GetLocation() LatLon {
	return NewLatLon(t.Lat, t.Lon)
}

func NewLatLon(lat, lon float64) (p LatLon) {
	p.Lat = lat
	p.Lon = lon

	return
}

func (l *LatLon) DistanceFrom(from LatLon) (hypotenuse float64) {
	side1 := ((l.Lat - from.Lat) * (l.Lat - from.Lat))
	side2 := ((l.Lon - from.Lon) * (l.Lon - from.Lon))
	hypotenuse = math.Sqrt(side1 + side2)

	return
}

func (l *LatLon) DistanceFromFloat(lat, lon float64) float64 {
	return l.DistanceFrom(NewLatLon(lat, lon))
}

func (r *IPGeolocationResponse) GetLocation() LatLon {
	return NewLatLon(r.GeopluginLatitude, r.GeopluginLongitude)
}

/*
FindServers finds some servers
*/
func FindServers() (servers ServerListClientConfigWrapper, err error) {
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

func (c *OoklaClient) FindNearestServers() (servers []TestServer, err error) {
	//c.FetchServers()
	return nil, nil
}
