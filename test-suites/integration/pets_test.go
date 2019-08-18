package integration

import (
	"context"
	"fmt"

	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestListPets() {
	t := suite.T()
	ctx := context.Background()
	params := pkg.ListPetsParams{}
	_, err := suite.apiClient.ListPetsWithResponse(ctx, &params)
	fmt.Println(err)
	assert.Equal(t, err, nil)
}

func (suite *TestSuite) TestCreatePet() {
	t := suite.T()
	ctx := context.Background()
	_, err := suite.apiClient.CreatePetsWithResponse(ctx)
	fmt.Println(err)
	assert.Equal(t, err, nil)
}

func (suite *TestSuite) TestShowPetById() {
	t := suite.T()
	ctx := context.Background()
	_, err := suite.apiClient.ShowPetByIdWithResponse(ctx, "2")
	fmt.Println(err)
	assert.Equal(t, err, nil)
}
