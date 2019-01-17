package ipapi

import "net/url"

func mustParseURL(addr string) *url.URL {
	url, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	return url
}
