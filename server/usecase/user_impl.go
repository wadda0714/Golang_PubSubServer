package usecase

import (
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/input"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/output"
)

type UserImpl struct {
	ID string
}

//func NewUserImpl(UserInfo *model.User) *UserImpl {
//	return &UserImpl{
//		UserInfo: UserInfo,
//	}
//}

func (u *UserImpl) Publish(p input.PublishInput) (output.PublishOutput, error) {
	return output.PublishOutput{}, nil
}

func (u *UserImpl) Subscribe(p input.SubscribeInput) (output.SubscribeOutput, error) {
	return output.SubscribeOutput{}, nil
}

func (u *UserImpl) Unsubscribe(p input.UnsubscribeInput) (output.UnsubscribeOutput, error) {
	return output.UnsubscribeOutput{}, nil
}

func (u *UserImpl) SendMessage(p input.SendMessageInput) (output.SendMessageOutput, error) {
	return output.SendMessageOutput{}, nil
}
