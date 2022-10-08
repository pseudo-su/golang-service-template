package persistence

import (
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/pseudo-su/golang-service-template/internal/persistence/gojet/model"
	"github.com/pseudo-su/golang-service-template/internal/persistence/gojet/table"
)

type PetsRepositoryInterface interface {
	GetPetByAPIID(ctx context.Context, id int64) (*Pet, error)
	CreatePet(ctx context.Context, in *PetValues) (*Pet, error)
	ListPets(ctx context.Context, in *PaginationValues) ([]Pet, error)
}

type PetsRepository struct {
	sqlDB *sql.DB
}

var _ PetsRepositoryInterface = &PetsRepository{}

// NewPetsRepository create new repository
func NewPetsRepository(db *sql.DB) *PetsRepository {
	return &PetsRepository{
		sqlDB: db,
	}
}

type Pet model.Pets

func (r *PetsRepository) GetPetByAPIID(ctx context.Context, id int64) (*Pet, error) {
	stmt := buildSqlGetPetByAPIID(id)

	var dest model.Pets
	err := stmt.QueryContext(ctx, r.sqlDB, &dest)

	if err != nil {
		return nil, err
	}
	resp := Pet(dest)
	return &resp, nil
}

func buildSqlGetPetByAPIID(id int64) postgres.Statement {
	stmt := postgres.SELECT(
		table.Pets.AllColumns,
	).FROM(
		table.Pets,
	).WHERE(
		table.Pets.APIID.EQ(postgres.Int(id)),
	)

	return stmt
}

type PetValues struct {
	Name string
	Tag  *string
}

func (r *PetsRepository) CreatePet(ctx context.Context, in *PetValues) (*Pet, error) {
	stmt := buildSqlCreatePet(in)

	var dest model.Pets
	err := stmt.QueryContext(ctx, r.sqlDB, &dest)

	if err != nil {
		return nil, err
	}

	resp := Pet(dest)
	return &resp, nil
}

func buildSqlCreatePet(in *PetValues) postgres.Statement {
	stmt := table.Pets.INSERT(
		table.Pets.Name,
		table.Pets.Tag,
	).MODEL(
		&model.Pets{
			Name: in.Name,
			Tag:  in.Tag,
		},
	).RETURNING(
		table.Pets.AllColumns,
	)

	return stmt
}

type PaginationValues struct {
	Offset int64
	Limit  int64
}

func (r *PetsRepository) ListPets(ctx context.Context, in *PaginationValues) ([]Pet, error) {
	stmt := buildSqlListPets(in)

	var dest []model.Pets
	err := stmt.QueryContext(ctx, r.sqlDB, &dest)

	if err != nil {
		return nil, err
	}
	resp := []Pet{}
	for _, dbItem := range dest {
		resp = append(resp, Pet(dbItem))
	}
	return resp, nil
}

func buildSqlListPets(in *PaginationValues) postgres.Statement {
	stmt := postgres.SELECT(
		table.Pets.AllColumns,
	).FROM(
		table.Pets,
	).LIMIT(
		in.Limit,
	).OFFSET(
		in.Offset,
	)

	return stmt
}
