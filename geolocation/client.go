package geolocation

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"
)

const (
	// IPGeolocationAPIEndpoint is the API endpoint
	IPGeolocationAPIEndpoint = "http://www.geoplugin.net/json.gp"
)

/*
LatLon represents a point in earth globe
*/
type LatLon struct {
	Lat float64
	Lon float64
}

/*
IPGeolocationResponse represents the API response
*/
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
	GeopluginLatitude               string  `json:"geoplugin_latitude"`
	GeopluginLongitude              string  `json:"geoplugin_longitude"`
	GeopluginLocationAccuracyRadius string  `json:"geoplugin_locationAccuracyRadius"`
	GeopluginTimezone               string  `json:"geoplugin_timezone"`
	GeopluginCurrencyCode           string  `json:"geoplugin_currencyCode"`
	GeopluginCurrencySymbol         string  `json:"geoplugin_currencySymbol"`
	GeopluginCurrencySymbolUTF8     string  `json:"geoplugin_currencySymbol_UTF8"`
	GeopluginCurrencyConverter      float64 `json:"geoplugin_currencyConverter"`
}

/*
NewLatLon creates a new LatLon object
*/
func NewLatLon(lat, lon float64) (p *LatLon) {
	p = &LatLon{
		Lat: lat,
		Lon: lon,
	}

	return
}

/*
DistanceFrom calculates the distance in km
based on the Haversine formula
https://en.wikipedia.org/wiki/Haversine_formula
*/
func (l *LatLon) DistanceFrom(from *LatLon) float64 {
	earthRadius := 6378.137
	dLat := l.Lat*math.Pi/180 - from.Lat*math.Pi/180
	dLon := l.Lon*math.Pi/180 - from.Lon*math.Pi/180
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(l.Lat*math.Pi/180)*math.Cos(from.Lat*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

/*
DistanceFromFloat calculates the distance between points
*/
func (l *LatLon) DistanceFromFloat(lat float64, lon float64) float64 {
	return l.DistanceFrom(NewLatLon(lat, lon))
}

/*
GetLocation gets the location from response
*/
func (r *IPGeolocationResponse) GetLocation() *LatLon {
	lat, err := strconv.ParseFloat(r.GeopluginLatitude, 64)
	if err != nil {
		panic(err.Error())
	}
	lon, err := strconv.ParseFloat(r.GeopluginLongitude, 64)
	if err != nil {
		panic(err.Error())
	}
	l := NewLatLon(lat, lon)
	if err != nil {
		panic(err.Error())
	}
	return l
}

/*
Client is the client to fetch location from API
*/
type Client struct {
}

/*
FetchIPLocation query APi for location
*/
func (c *Client) FetchIPLocation() (response IPGeolocationResponse, err error) {
	res, err := http.Get(IPGeolocationAPIEndpoint)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	response, err = parseLocationResponse(res.Body)

	return
}

func parseLocationResponse(r io.ReadCloser) (response IPGeolocationResponse, err error) {
	err = json.NewDecoder(r).Decode(&response)

	return
}
