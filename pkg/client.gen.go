// Package pkg provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/pseudo-su/oapi-codegen DO NOT EDIT.
package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pseudo-su/oapi-codegen/pkg/runtime"
	"io/ioutil"
	"net/http"
	"strings"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(req *http.Request, ctx context.Context) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListPets request
	ListPets(ctx context.Context, params *ListPetsParams) (*http.Response, error)

	// CreatePets request
	CreatePets(ctx context.Context) (*http.Response, error)

	// ShowPetById request
	ShowPetById(ctx context.Context, petId string) (*http.Response, error)
}

func (c *Client) ListPets(ctx context.Context, params *ListPetsParams) (*http.Response, error) {
	req, err := NewListPetsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CreatePets(ctx context.Context) (*http.Response, error) {
	req, err := NewCreatePetsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ShowPetById(ctx context.Context, petId string) (*http.Response, error) {
	req, err := NewShowPetByIdRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewListPetsRequest generates requests for ListPets
func NewListPetsRequest(server string, params *ListPetsParams) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/pets", server)

	var queryStrings []string

	var queryParam0 string
	if params.Limit != nil {

		queryParam0, err = runtime.StyleParam("form", true, "limit", *params.Limit)
		if err != nil {
			return nil, err
		}

		queryStrings = append(queryStrings, queryParam0)
	}

	if len(queryStrings) != 0 {
		queryUrl += "?" + strings.Join(queryStrings, "&")
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreatePetsRequest generates requests for CreatePets
func NewCreatePetsRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/pets", server)

	req, err := http.NewRequest("POST", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewShowPetByIdRequest generates requests for ShowPetById
func NewShowPetByIdRequest(server string, petId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "petId", petId)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/pets/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: http.Client{},
			Server: server,
		},
	}
}

// NewClientWithResponsesAndRequestEditorFunc takes in a RequestEditorFn callback function and returns a ClientWithResponses with a default Client:
func NewClientWithResponsesAndRequestEditorFunc(server string, reqEditorFn RequestEditorFn) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client:        http.Client{},
			Server:        server,
			RequestEditor: reqEditorFn,
		},
	}
}

type listPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pets
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r listPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r listPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type createPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r createPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r createPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type showPetByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r showPetByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r showPetByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListPetsWithResponse request returning *ListPetsResponse
func (c *ClientWithResponses) ListPetsWithResponse(ctx context.Context, params *ListPetsParams) (*listPetsResponse, error) {
	rsp, err := c.ListPets(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParselistPetsResponse(rsp)
}

// CreatePetsWithResponse request returning *CreatePetsResponse
func (c *ClientWithResponses) CreatePetsWithResponse(ctx context.Context) (*createPetsResponse, error) {
	rsp, err := c.CreatePets(ctx)
	if err != nil {
		return nil, err
	}
	return ParsecreatePetsResponse(rsp)
}

// ShowPetByIdWithResponse request returning *ShowPetByIdResponse
func (c *ClientWithResponses) ShowPetByIdWithResponse(ctx context.Context, petId string) (*showPetByIdResponse, error) {
	rsp, err := c.ShowPetById(ctx, petId)
	if err != nil {
		return nil, err
	}
	return ParseshowPetByIdResponse(rsp)
}

// ParselistPetsResponse parses an HTTP response from a ListPetsWithResponse call
func ParselistPetsResponse(rsp *http.Response) (*listPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &listPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &Pets{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}
	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}
	}

	return response, nil
}

// ParsecreatePetsResponse parses an HTTP response from a CreatePetsWithResponse call
func ParsecreatePetsResponse(rsp *http.Response) (*createPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &createPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.StatusCode == 201:
		break // No content-type
	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}
	}

	return response, nil
}

// ParseshowPetByIdResponse parses an HTTP response from a ShowPetByIdWithResponse call
func ParseshowPetByIdResponse(rsp *http.Response) (*showPetByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &showPetByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &Pet{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}
	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}
	}

	return response, nil
}
