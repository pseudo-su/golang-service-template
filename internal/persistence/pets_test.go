package persistence

import "testing"

func TestBuildSQLGetPetByAPIID(t *testing.T) {
	mockTimeNow()
	defer resetTimeNow()

	stmt := buildSqlGetPetByAPIID(1)
	assertStatement(
		t,
		stmt,
		[]interface{}{int64(1)},
		`
SELECT pets.id AS "pets.id",
     pets.api_id AS "pets.api_id",
     pets.created_at AS "pets.created_at",
     pets.updated_at AS "pets.updated_at",
     pets.deleted_at AS "pets.deleted_at",
     pets.name AS "pets.name",
     pets.tag AS "pets.tag"
FROM public.pets
WHERE pets.api_id = 1;
`, `
SELECT pets.id AS "pets.id",
     pets.api_id AS "pets.api_id",
     pets.created_at AS "pets.created_at",
     pets.updated_at AS "pets.updated_at",
     pets.deleted_at AS "pets.deleted_at",
     pets.name AS "pets.name",
     pets.tag AS "pets.tag"
FROM public.pets
WHERE pets.api_id = $1;
`)
}

func TestBuildSQLCreatePetWithTag(t *testing.T) {
	mockTimeNow()
	defer resetTimeNow()

	stmt := buildSqlCreatePet(&PetValues{
		Name: "My pet",
		Tag:  strRef("dog"),
	})
	assertStatement(
		t,
		stmt,
		[]interface{}{"My pet", "dog"},
		`
INSERT INTO public.pets (name, tag)
VALUES ('My pet', 'dog')
RETURNING pets.id AS "pets.id",
          pets.api_id AS "pets.api_id",
          pets.created_at AS "pets.created_at",
          pets.updated_at AS "pets.updated_at",
          pets.deleted_at AS "pets.deleted_at",
          pets.name AS "pets.name",
          pets.tag AS "pets.tag";
`, `
INSERT INTO public.pets (name, tag)
VALUES ($1, $2)
RETURNING pets.id AS "pets.id",
          pets.api_id AS "pets.api_id",
          pets.created_at AS "pets.created_at",
          pets.updated_at AS "pets.updated_at",
          pets.deleted_at AS "pets.deleted_at",
          pets.name AS "pets.name",
          pets.tag AS "pets.tag";
`)
}

func TestBuildSQLCreatePetWithoutTag(t *testing.T) {
	mockTimeNow()
	defer resetTimeNow()

	stmt := buildSqlCreatePet(&PetValues{
		Name: "My pet",
		Tag:  nil,
	})
	assertStatement(
		t,
		stmt,
		[]interface{}{"My pet", nil},
		`
INSERT INTO public.pets (name, tag)
VALUES ('My pet', NULL)
RETURNING pets.id AS "pets.id",
          pets.api_id AS "pets.api_id",
          pets.created_at AS "pets.created_at",
          pets.updated_at AS "pets.updated_at",
          pets.deleted_at AS "pets.deleted_at",
          pets.name AS "pets.name",
          pets.tag AS "pets.tag";
`, `
INSERT INTO public.pets (name, tag)
VALUES ($1, $2)
RETURNING pets.id AS "pets.id",
          pets.api_id AS "pets.api_id",
          pets.created_at AS "pets.created_at",
          pets.updated_at AS "pets.updated_at",
          pets.deleted_at AS "pets.deleted_at",
          pets.name AS "pets.name",
          pets.tag AS "pets.tag";
`)
}

func TestBuildSQLListPets(t *testing.T) {
	mockTimeNow()
	defer resetTimeNow()

	stmt := buildSqlListPets(&PaginationValues{
		Offset: 0,
		Limit:  10,
	})
	assertStatement(
		t,
		stmt,
		[]interface{}{int64(10), int64(0)},
		`
SELECT pets.id AS "pets.id",
     pets.api_id AS "pets.api_id",
     pets.created_at AS "pets.created_at",
     pets.updated_at AS "pets.updated_at",
     pets.deleted_at AS "pets.deleted_at",
     pets.name AS "pets.name",
     pets.tag AS "pets.tag"
FROM public.pets
LIMIT 10
OFFSET 0;
`, `
SELECT pets.id AS "pets.id",
     pets.api_id AS "pets.api_id",
     pets.created_at AS "pets.created_at",
     pets.updated_at AS "pets.updated_at",
     pets.deleted_at AS "pets.deleted_at",
     pets.name AS "pets.name",
     pets.tag AS "pets.tag"
FROM public.pets
LIMIT $1
OFFSET $2;
`)
}
