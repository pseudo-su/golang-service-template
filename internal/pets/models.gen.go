// Package pets provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package pets

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NotPet defines model for NotPet.
type NotPet struct {
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// NotPets defines model for NotPets.
type NotPets []NotPet

// Pet defines model for Pet.
type Pet struct {
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pets defines model for Pets.
type Pets []Pet

// ListPetsParams defines parameters for ListPets.
type ListPetsParams struct {

	// How many items to return at one time (max 100)
	Limit *int32 `json:"limit,omitempty"`
}