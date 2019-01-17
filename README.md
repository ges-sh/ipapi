### Go client for ip-api API 

#### What is ip-api? 
http://ip-api.com/

### How to install:
`go get github.com/ges-sh/ipapi`

[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/ges-sh/ipapi)


### How to use:
```go
package main

import (
	"fmt"
	"net"

	"github.com/ges-sh/ipapi"
)

func main() {
	client, _ := ipapi.New("en")

	ipLoc, _ := client.FetchIPLocation(net.ParseIP("24.48.0.1"))

	fmt.Printf("%+v\n", ipLoc) // {Query:24.48.0.1 Status:success Message: Country:Canada CountryCode:CA Region:QC RegionName:Quebec City:Vaudreuil-Dorion Zip:H9X Lat:45.4 Lon:-73.9333 Timezone:America/Toronto ISP:Le Groupe Videotron Ltee Org:Videotron Ltee AS:AS5769 Videotron Telecom Ltee}
}
```

#### Using string instead of net.IP
```go
...
	ipLoc, _ := client.FetchIPLocationStr("24.48.0.1")

	fmt.Printf("%+v\n", ipLoc) // {Query:24.48.0.1 Status:success Message: Country:Canada CountryCode:CA Region:QC RegionName:Quebec City:Vaudreuil-Dorion Zip:H9X Lat:45.4 Lon:-73.9333 Timezone:America/Toronto ISP:Le Groupe Videotron Ltee Org:Videotron Ltee AS:AS5769 Videotron Telecom Ltee}
}
```

 ### To use Pro version of api:
 ```go
  client.Pro("your_api_key")
 ```
