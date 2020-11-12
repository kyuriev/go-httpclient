package go_httpclient

import "github.com/kyuriev/go-httpclient/gohttp"

func basicExample() {
	client := gohttp.New()
	client.Get()
}
