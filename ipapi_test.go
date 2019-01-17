package ipapi

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	resp *http.Response
	err  error
}

func (m mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func TestFetchIPLocation(t *testing.T) {
	testCases := []struct {
		name       string
		httpClient HTTPClient
		ip         net.IP
		isPro      bool

		expLoc   IPLocation
		expError error
	}{
		{
			name: "All good",
			httpClient: mockHTTPClient{
				resp: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(`
						{
							"query": "24.48.0.1",
							"status": "success",
							"country": "Canada",
							"countryCode": "CA",
							"region": "QC",
							"regionName": "Quebec",
							"city": "Vaudreuil-Dorion",
							"zip": "H9X",
							"lat": 45.4,
							"lon": -73.9333,
							"timezone": "America/Toronto",
							"isp": "Le Groupe Videotron Ltee",
							"org": "Videotron Ltee",
							"as": "AS5769 Videotron Telecom Ltee"
						}
				`)),
				},
			},
			ip: net.IP{24, 48, 0, 1},
			expLoc: IPLocation{
				Query:       "24.48.0.1",
				Status:      "success",
				Country:     "Canada",
				CountryCode: "CA",
				Region:      "QC",
				RegionName:  "Quebec",
				City:        "Vaudreuil-Dorion",
				Zip:         "H9X",
				Lat:         45.4,
				Lon:         -73.9333,
				Timezone:    "America/Toronto",
				ISP:         "Le Groupe Videotron Ltee",
				Org:         "Videotron Ltee",
				AS:          "AS5769 Videotron Telecom Ltee",
			},
		},
		{
			name:     "With invalid IP",
			expError: ErrInvalidIP,
		},
		{
			name: "With Http Client error",
			httpClient: mockHTTPClient{
				err: errors.New("some error"),
			},
			ip:       net.IP{24, 48, 0, 1},
			expError: errors.New("some error"),
		},
		{
			name: "With reserved IP range error",
			httpClient: mockHTTPClient{
				resp: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(`
						{
							"status": "fail",
							"message": "reserved range"
						}
					`)),
				},
			},
			ip:       net.IP{24, 48, 0, 1},
			expError: ErrReservedIP,
		},
		{
			name: "With private IP range error",
			httpClient: mockHTTPClient{
				resp: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(`
					{
						"status": "fail",
						"message": "private range"
					}
				`)),
				},
			},
			ip:       net.IP{24, 48, 0, 1},
			expError: ErrPrivateRangeIP,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client := NewWithClient(tt.httpClient, "")

			loc, err := client.FetchIPLocation(tt.ip)
			if err != nil {
				if err.Error() != tt.expError.Error() {
					t.Errorf("\nExp error: %v\nGot error: %v", tt.expError, err)
				}
				return
			}

			if !locMatch(tt.expLoc, loc) {
				t.Errorf("\nExp loc: %v\nGot loc: %v", tt.expLoc, loc)
				return
			}
		})
	}
}

func locMatch(loc1, loc2 IPLocation) bool {
	return loc1.Query == loc2.Query &&
		loc1.Status == loc2.Status &&
		loc1.Message == loc2.Message &&
		loc1.Country == loc2.Country &&
		loc1.CountryCode == loc2.CountryCode &&
		loc1.Region == loc2.Region &&
		loc1.RegionName == loc2.RegionName &&
		loc1.City == loc2.City &&
		loc1.Zip == loc2.Zip &&
		loc1.Lat == loc2.Lat &&
		loc1.Lon == loc2.Lon &&
		loc1.Timezone == loc2.Timezone &&
		loc1.ISP == loc2.ISP &&
		loc1.Org == loc2.Org &&
		loc1.AS == loc2.AS
}
