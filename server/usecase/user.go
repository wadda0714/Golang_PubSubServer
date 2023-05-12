package usecase

import (
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/input"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/output"
)

type UserUsecase interface {
	Publish(p input.PublishInput) (o output.PublishOutput, err error)
	Subscribe(p input.SubscribeInput) (o output.SubscribeOutput, err error)
	Unsubscribe(p input.UnsubscribeInput) (o output.UnsubscribeOutput, err error)
	SendMessage(p input.SendMessageInput) (o output.SendMessageOutput, err error)
}
