//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Pets = newPetsTable("public", "pets", "")

type petsTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnString
	APIID     postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestampz
	UpdatedAt postgres.ColumnTimestampz
	DeletedAt postgres.ColumnTimestampz
	Name      postgres.ColumnString
	Tag       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PetsTable struct {
	petsTable

	EXCLUDED petsTable
}

// AS creates new PetsTable with assigned alias
func (a PetsTable) AS(alias string) *PetsTable {
	return newPetsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PetsTable with assigned schema name
func (a PetsTable) FromSchema(schemaName string) *PetsTable {
	return newPetsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PetsTable with assigned table prefix
func (a PetsTable) WithPrefix(prefix string) *PetsTable {
	return newPetsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PetsTable with assigned table suffix
func (a PetsTable) WithSuffix(suffix string) *PetsTable {
	return newPetsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPetsTable(schemaName, tableName, alias string) *PetsTable {
	return &PetsTable{
		petsTable: newPetsTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newPetsTableImpl("", "excluded", ""),
	}
}

func newPetsTableImpl(schemaName, tableName, alias string) petsTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		APIIDColumn     = postgres.IntegerColumn("api_id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn = postgres.TimestampzColumn("deleted_at")
		NameColumn      = postgres.StringColumn("name")
		TagColumn       = postgres.StringColumn("tag")
		allColumns      = postgres.ColumnList{IDColumn, APIIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, TagColumn}
		mutableColumns  = postgres.ColumnList{APIIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, TagColumn}
	)

	return petsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		APIID:     APIIDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		Name:      NameColumn,
		Tag:       TagColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
