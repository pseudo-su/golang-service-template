package internal

import (
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate ../tools/mockgen -destination=./pets/mocks_gen_test.go -package=pets github.com/pseudo-su/golang-service-template/internal/persistence PetsRepositoryInterface
