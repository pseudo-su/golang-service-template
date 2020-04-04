package pkg

import (
	"context"
	"fmt"
	"net/http"
)

// Client is the API client to make requests against golang-service-template
type Client struct {
	gen generatedClientInterface
}

var _ ClientInterface = &Client{}

// Client is the API client interface to make requests against golang-service-template
type ClientInterface interface {
	// ListPets request
	ListPets(ctx context.Context, params *ListPetsParams) (*Pets, error)

	// CreatePets request
	CreatePets(ctx context.Context) error

	// ShowPetByID request
	ShowPetByID(ctx context.Context, petID string) (*Pet, error)
}

// NewClient initializes and returns a Client configured using options
func NewClient(server string, opts ...clientOption) (*Client, error) {
	genClient, err := newGeneratedClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{gen: genClient}, nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: '%v', message: '%s'", e.Code, e.Message)
}

func unexpectedResponse(resp *http.Response) error {
	return fmt.Errorf(
		"unexpected response from url '%s': statusCode '%v', status '%s'",
		resp.Request.URL.String(),
		resp.StatusCode,
		resp.Status,
	)
}

// ListPets returns a list of pets
func (c *Client) ListPets(ctx context.Context, params *ListPetsParams) (*Pets, error) {
	resp, err := c.gen.ListPetsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.JSON200 != nil:
		return resp.JSON200, nil
	case resp.JSONDefault != nil:
		return nil, resp.JSONDefault
	default:
		return nil, unexpectedResponse(resp.HTTPResponse)
	}
}

// CreatePets creates a new Pet
func (c *Client) CreatePets(ctx context.Context) error {
	resp, err := c.gen.CreatePetsWithResponse(ctx)
	if err != nil {
		return err
	}

	switch {
	case resp.StatusCode() == http.StatusOK:
		return nil
	case resp.JSONDefault != nil:
		return resp.JSONDefault
	default:
		return unexpectedResponse(resp.HTTPResponse)
	}
}

// ShowPetById returns a Pet if one is found for a given ID
func (c *Client) ShowPetByID(ctx context.Context, petID string) (*Pet, error) {
	resp, err := c.gen.ShowPetByIdWithResponse(ctx, petID)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.JSON200 != nil:
		return resp.JSON200, nil
	case resp.JSONDefault != nil:
		return nil, resp.JSONDefault
	default:
		return nil, unexpectedResponse(resp.HTTPResponse)
	}
}
