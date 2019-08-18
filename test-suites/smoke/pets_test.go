package smoke

import (
	"context"
	"fmt"

	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestListPets() {
	t := suite.T()
	ctx := context.Background()
	params := &pkg.FindPetsParams{}
	_, err := suite.apiClient.FindPets(ctx, params)
	fmt.Println(err)
	assert.Equal(t, err, nil)
}

func (suite *TestSuite) TestCreatePet() {
	t := suite.T()
	ctx := context.Background()
	body := pkg.AddPetJSONRequestBody{
		Name: "Bobby",
		Tag:  nil,
	}
	_, err := suite.apiClient.AddPetWithResponse(ctx, body)
	fmt.Println(err)
	assert.Equal(t, err, nil)
}

func (suite *TestSuite) TestShowPetById() {
	t := suite.T()
	ctx := context.Background()
	_, err := suite.apiClient.FindPetByIdWithResponse(ctx, int64(2))
	fmt.Println(err)
	assert.Equal(t, err, nil)
}

func (suite *TestSuite) TestDeletePetById() {
	t := suite.T()
	ctx := context.Background()
	_, err := suite.apiClient.DeletePetWithResponse(ctx, int64(1))
	fmt.Println(err)
	assert.Equal(t, err, nil)
}
