package pkg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Client ...
type Client struct {
	gen generatedClientInterface
}

var _ ClientInterface = &Client{}

// ClientInterface ...
type ClientInterface interface {
	// ListPets request
	ListPets(ctx context.Context, params *ListPetsParams) (*Pets, error)

	// CreatePets request
	CreatePets(ctx context.Context) error

	// ShowPetById request
	ShowPetById(ctx context.Context, petId string) (*Pet, error)
}

// NewClient ...
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

// ListPets ...
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
		return nil, errors.New("Unexpected response")
	}
}

// ListPets ...
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
		return errors.New("Unexpected response")
	}
}

// ListPets ...
func (c *Client) ShowPetById(ctx context.Context, petId string) (*Pet, error) {
	resp, err := c.gen.ShowPetByIdWithResponse(ctx, petId)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.JSON200 != nil:
		return resp.JSON200, nil
	case resp.JSONDefault != nil:
		return nil, resp.JSONDefault
	default:
		return nil, fmt.Errorf("Unexpected response: statusCode '%v', status '%s'", resp.StatusCode(), resp.Status())
	}
}
