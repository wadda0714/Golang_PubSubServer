package usecase

import (
	"github.com/wadda0714/Golang_PubSubServer/server/domain/query"
	//"github.com/wadda0714/Golang_PubSubServer/server/domain/repository"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/input"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/output"
)

type User struct {
	userQuery query.UserQuery
	//userRepo  repository.UserRepository
}

func NewUser(
	userQuery query.UserQuery,
	//userRepo repository.UserRepository,
) User {
	return User{
		userQuery,
		//userRepo,
	}
}

func (u *User) Publish(p input.PublishInput) (output.PublishOutput, error) {
	return output.PublishOutput{}, nil
}

func (u *User) Subscribe(p input.SubscribeInput) (output.SubscribeOutput, error) {
	Messages, err := u.userQuery.GetMessagesByRoomID(p.RoomID)
	if err != nil {
		return output.SubscribeOutput{}, err
	}
	return output.SubscribeOutput{
		Messages: Messages,
	}, nil
}

func (u *User) Unsubscribe(p input.UnsubscribeInput) (output.UnsubscribeOutput, error) {
	return output.UnsubscribeOutput{}, nil
}

func (u *User) SendMessage(p input.SendMessageInput) (output.SendMessageOutput, error) {
	return output.SendMessageOutput{}, nil
}
