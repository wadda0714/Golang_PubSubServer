package repository

import (
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
)

type UserRepository interface {
	AddUserToDB(user *model.User) error
	DeleteUserFromDB(user *model.User) error
	AddMessageToDB(message *model.Message) error
	EnterRoom() error
}
