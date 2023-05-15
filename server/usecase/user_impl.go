package usecase

import (
	"github.com/wadda0714/Golang_PubSubServer/server/domain/query"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/repository"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/input"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/output"
)

type user struct {
	userQuery query.UserQuery
	userRepo  repository.UserRepository
}

func NewUser(
	userQuery query.UserQuery,
	userRepo repository.UserRepository,
) *user {
	return &user{
		userQuery,
		userRepo,
	}
}

func (u *user) Publish(p input.PublishInput) (output.PublishOutput, error) {
	return output.PublishOutput{}, nil
}

func (u *user) Subscribe(p input.SubscribeInput) (output.SubscribeOutput, error) {
	return output.SubscribeOutput{}, nil
}

func (u *user) Unsubscribe(p input.UnsubscribeInput) (output.UnsubscribeOutput, error) {
	return output.UnsubscribeOutput{}, nil
}

func (u *user) SendMessage(p input.SendMessageInput) (output.SendMessageOutput, error) {
	return output.SendMessageOutput{}, nil
}
