// Package internal provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/pseudo-su/oapi-ui-codegen DO NOT EDIT.
package internal

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVQW/jNhP9K4P5vkMLKJazW/SgU9PtAjXQdgM0t20OE2kscSuRzHDk2DD03wtSUux1",
	"7KYLtMCebJKjxzdv3nD2WLrOO8tWAxZ7DGXDHaW/70WcxD9enGdRw2m7dBXH37WTjhQLNFbfvsEMded5",
	"XHLNgkOGHYdAdYqeDoOKsTUOQ4bCj70RrrD4OGIe4u+fwdzDJy41Yt2yvuRiqlMm3393loml7hyNDJXq",
	"1+mZCieIC8xGMspd+vN/4TUW+L/8IG0+6ZrHNIZnEBKhHQ7xPmPXbpTXKpUpV+7ItFjMWz/UriVbXwWW",
	"jSn5SrnzLSkvStfN/Aq8EIQZ9hLBaqNN/xC/yX3gvnJXoc8vfTRkWHEoxXg1zmKBd40JYAKQBd5S51uG",
	"D57tze0KgucS1EEfGCgAgfCahW3J8NSwhaAkamwNBF5cFA/6ENdfwugPixm2pmQbUj2npH9d3SVVjbZx",
	"+fsT1TULxMqok5j8hiWMKVwvlotljHaeLXmDBb5NWxl60iZVMPdTTevRdZ+L8IsJCtS2oA1DikxgQvF8",
	"VWGBrQl6Ox54EupYWQIWH0+RfnZP0JHdQfJOVE9Ye7FACs4yqOkYvuloC9fL5bcYXYIFPvYsu0PFW9MZ",
	"xWxq3X/UmsN9NHjwzoaxkd4sl7P52KaUyfvWlCml/FOIbPdHN7xi8TB6+vNkb8BTzRUk04Nbz9I1TFWS",
	"Z4/bK8vbM4rfQGvsn1GfKHmMSVgR5JDGsQKn/TyyWVPf6r+W5/g+nkm0t7z1XCpXwFNMhqHvOpLdsX2m",
	"/JXq6A1My/shQ+/CGQ3eCZNy7B7WF4Yr0+FkuZPKXr/E+q1v22fl8CvS5iTJE2mGbOzMfO9ZV9VwsUNX",
	"du1g7QQoPUtmbcqzsoXGPd2y/rhbVa+16l3DYKrouKnrp24VwxueOzM+IIfGTCTxeJSo9Px3Nv2P2/Jc",
	"Qd7P5ZgvjnkRbKg1ce+x56Bfk0Mul/bULPEzls1cznH8XRos+eYaT5m860XYKvzkOjIWh2wGaVR9KPJ8",
	"X6WT4UtAP2jDMkNmuCEx9NCO5R7hRjtPaiN5s5hG7TToT+pnlQXKPqjrYARI7yRJPRX8mHWRpwaKQ3ER",
	"xim5MC7yjN6bFNwfGTi8uPAmDT5j6zSvKgdPRpvxLRvuh78CAAD//5HtG8FQCgAA",
}

// GetOpenAPISpec returns the Swagger specification corresponding to the generated code
// in this file.
func GetOpenAPISpec() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
