package query

import (
	"database/sql"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/query"
)

type User struct {
	//DBとの接続情報
	DB *sql.DB
}

func NewUser(db *sql.DB) query.UserQuery {
	return &User{
		DB: db,
	}
}

func (u *User) GetUsers(roomID int) ([]*model.User, error) {
	return nil, nil
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	return nil, nil
}
