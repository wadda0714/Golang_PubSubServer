package usecase

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/wadda0714/Golang_PubSubServer/server/infra/query"
	"github.com/wadda0714/Golang_PubSubServer/server/infra/repository"
)

func NewUserDI(db *sql.DB) User {
	wire.Build(
		query.NewUser,
		repository.NewUser,
		NewUser,
	)
	return User{}
}
