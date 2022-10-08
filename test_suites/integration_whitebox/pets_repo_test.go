package integration

import (
	"context"

	"github.com/pseudo-su/golang-service-template/internal/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestPetsRepo() {
	t := suite.T()
	ctx := context.Background()

	repo := persistence.NewPetsRepository(suite.sqlDB)

	pet1, err := repo.CreatePet(ctx, &persistence.PetValues{
		Name: "Pet 1",
		Tag:  nil,
	})
	require.NoError(t, err)
	assert.Equal(t, pet1.Name, "Pet 1")

	dog := "dog"
	pet2, err := repo.CreatePet(ctx, &persistence.PetValues{
		Name: "Pet 2",
		Tag:  &dog,
	})
	require.NoError(t, err)
	assert.Equal(t, pet2.Name, "Pet 2")

	fetchedPet, err := repo.GetPetByAPIID(ctx, pet2.APIID)
	require.NoError(t, err)
	assert.Equal(t, fetchedPet.Name, "Pet 2")

	petsResp, err := repo.ListPets(ctx, &persistence.PaginationValues{
		Offset: 0,
		Limit:  10,
	})
	require.NoError(t, err)
	require.Equal(t, len(petsResp), 2)
}
