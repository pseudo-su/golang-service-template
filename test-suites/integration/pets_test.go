package integration

import (
	"context"
	"fmt"

	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestListPets() {
	t := suite.T()
	ctx := context.Background()
	params := pkg.ListPetsParams{}
	pets, err := suite.apiClient.ListPets(ctx, &params)
	fmt.Println(err)
	require.Equal(t, nil, err)
	assert.NotEmpty(t, pets)
}

func (suite *TestSuite) TestCreatePet() {
	t := suite.T()
	ctx := context.Background()
	err := suite.apiClient.CreatePets(ctx)
	fmt.Println(err)
	require.Equal(t, nil, err)
}

func (suite *TestSuite) TestShowPetById() {
	t := suite.T()
	ctx := context.Background()
	pet, err := suite.apiClient.ShowPetById(ctx, "2")
	fmt.Println(err)
	require.Equal(t, nil, err)
	assert.NotEmpty(t, pet)
}
