package smoke

import (
	"context"
	"fmt"
	"testing"

	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/stretchr/testify/assert"
)

func TestStatements(t *testing.T) {
	_, apiClient := smokeSuiteSetup()

	t.Run("List Pets", func(t *testing.T) {
		ctx := context.Background()
		params := pkg.ListPetsParams{}
		_, err := apiClient.ListPetsWithResponse(ctx, &params)
		fmt.Println(err)
		assert.Equal(t, err, nil)
	})

	t.Run("Create Pet", func(t *testing.T) {
		ctx := context.Background()
		_, err := apiClient.CreatePetsWithResponse(ctx)
		fmt.Println(err)
		assert.Equal(t, err, nil)
	})

	t.Run("Show Pet By ID", func(t *testing.T) {
		ctx := context.Background()
		_, err := apiClient.ShowPetByIdWithResponse(ctx, "2")
		fmt.Println(err)
		assert.Equal(t, err, nil)
	})
}
