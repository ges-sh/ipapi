// Package ipapi provides easy way to fetch IP location data from ip-api API. For now package supports JSON responses only.
package ipapi

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	baseURL *url.URL
	apiKey  string
	client  HTTPClient
	lang    string
}

var freeIPAPIEndpoint = mustParseURL("http://ip-api.com")
var proIPAPIEndpoint = mustParseURL("http://pro.ip-api.com")

// New returns new Client. If no lang is specified, default (en) will be used.
func New(lang string) Client {
	return NewWithClient(&http.Client{}, lang)
}

const defaultLang = "en"

// NewWithClient is the same as New, but allows to pass custom HTTPClient to the Client. If no lang is specified, default (en) will be used.
func NewWithClient(httpClient HTTPClient, lang string) Client {
	if lang == "" {
		lang = defaultLang
	}

	client := Client{
		baseURL: freeIPAPIEndpoint,
		client:  httpClient,
		lang:    lang,
	}

	return client
}

// IPLocation contains location data about ip address
type IPLocation struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
}

// ErrInvalidIP is returned when provided ip address is invalid
var ErrInvalidIP = errors.New("invalid ip address")

// Errors that may be returned by ip api
var (
	ErrPrivateRangeIP = errors.New("ip is from private ip range")
	ErrReservedIP     = errors.New("ip is from reserved ip range")
	ErrInvalidQuery   = errors.New("invalid query")
)

// ipAPIErrors assigns ip api error messages to predefined errors
var ipAPIErrors = map[string]error{
	"reserved range": ErrReservedIP,
	"private range":  ErrPrivateRangeIP,
	"invalid query":  ErrInvalidQuery,
}

// FetchIPLocation fetches location info about provided ip address. If ip is invalid, ErrInvalidIP will be returned
func (c Client) FetchIPLocation(ip net.IP) (IPLocation, error) {
	if ip == nil {
		return IPLocation{}, ErrInvalidIP
	}

	locEndpoint := mustParseURL("/json/" + ip.String())

	targetURL := c.baseURL.ResolveReference(locEndpoint)

	if c.apiKey != "" {
		query := url.Values{}
		query.Set("key", c.apiKey)
		targetURL.RawQuery = query.Encode()
	}

	request, err := http.NewRequest(http.MethodGet, targetURL.String(), nil)
	if err != nil {
		return IPLocation{}, err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return IPLocation{}, err
	}
	defer resp.Body.Close()

	var locationData IPLocation
	err = json.NewDecoder(resp.Body).Decode(&locationData)
	if err != nil {
		return IPLocation{}, err
	}

	return locationData, ipAPIErrors[locationData.Message]
}

// Pro enables usage of pro version of ip-api API.
func (c *Client) Pro(apiKey string) {
	c.baseURL = proIPAPIEndpoint
	c.apiKey = apiKey
}
