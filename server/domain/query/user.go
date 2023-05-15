package query

import (
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
)

type UserQuery interface {
	GetUsers(roomID int) ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
}
