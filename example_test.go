// +build example

package ipapi_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/ges-sh/ipapi"
)

// Errors omitted for brevity

func TestFetchIPLocation(t *testing.T) {
	client := ipapi.New("")

	ipLoc, _ := client.FetchIPLocation(net.ParseIP("24.48.0.1"))

	fmt.Printf("%+v\n", ipLoc) // {Query:24.48.0.1 Status:success Message: Country:Canada CountryCode:CA Region:QC RegionName:Quebec City:Vaudreuil-Dorion Zip:H9X Lat:45.4 Lon:-73.9333 Timezone:America/Toronto ISP:Le Groupe Videotron Ltee Org:Videotron Ltee AS:AS5769 Videotron Telecom Ltee}
}

func TestFetchIPLocationPro(t *testing.T) {
	client := ipapi.New("")
	client.Pro("your_api_key")

	ipLoc, _ := client.FetchIPLocation(net.ParseIP("24.48.0.1"))

	fmt.Printf("%+v\n", ipLoc) // {Query:24.48.0.1 Status:success Message: Country:Canada CountryCode:CA Region:QC RegionName:Quebec City:Vaudreuil-Dorion Zip:H9X Lat:45.4 Lon:-73.9333 Timezone:America/Toronto ISP:Le Groupe Videotron Ltee Org:Videotron Ltee AS:AS5769 Videotron Telecom Ltee}
}
