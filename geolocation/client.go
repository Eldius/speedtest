package geolocation

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"
)

const (
	IPGeolocationAPIEndpoint = "http://www.geoplugin.net/json.gp"
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
	GeopluginLatitude               string  `json:"geoplugin_latitude"`
	GeopluginLongitude              string  `json:"geoplugin_longitude"`
	GeopluginLocationAccuracyRadius string  `json:"geoplugin_locationAccuracyRadius"`
	GeopluginTimezone               string  `json:"geoplugin_timezone"`
	GeopluginCurrencyCode           string  `json:"geoplugin_currencyCode"`
	GeopluginCurrencySymbol         string  `json:"geoplugin_currencySymbol"`
	GeopluginCurrencySymbolUTF8     string  `json:"geoplugin_currencySymbol_UTF8"`
	GeopluginCurrencyConverter      float64 `json:"geoplugin_currencyConverter"`
}

func NewLatLon(lat, lon float64) (p *LatLon) {
	p = &LatLon{
		Lat: lat,
		Lon: lon,
	}

	return
}

func (l *LatLon) DistanceFrom(from *LatLon) (hypotenuse float64) {
	side1 := ((l.Lat - from.Lat) * (l.Lat - from.Lat))
	side2 := ((l.Lon - from.Lon) * (l.Lon - from.Lon))
	hypotenuse = math.Sqrt(side1 + side2)

	return
}

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

type GeolocationClient struct {
}

func (c *GeolocationClient) FetchIPLocation() (response IPGeolocationResponse, err error) {
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
