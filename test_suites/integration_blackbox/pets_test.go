package integration

import (
	"context"
	"fmt"

	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestCreateAndListPets() {
	t := suite.T()
	ctx := context.Background()
	var err error
	// Create pets
	pet1, err := suite.apiClient.CreatePetWithResponse(ctx, pkg.CreatePetJSONRequestBody{
		Name: "Robert",
		Tag:  nil,
	})
	assert.NoError(t, err)
	_, err = suite.apiClient.CreatePetWithResponse(ctx, pkg.CreatePetJSONRequestBody{
		Name: "Peter",
		Tag:  nil,
	})
	assert.NoError(t, err)

	// List pets
	_, err = suite.apiClient.ListPetsWithResponse(ctx, &pkg.ListPetsParams{})
	assert.NoError(t, err)

	// Get a pet by ID
	_, err = suite.apiClient.GetPetByIdWithResponse(ctx, fmt.Sprint(pet1.JSON201.Id))
	assert.NoError(t, err)
}
