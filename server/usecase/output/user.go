package output

import (
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
)

type SubscribeOutput struct {
	Messages []*model.Message
}

type UnsubscribeOutput struct {
}

type PublishOutput struct {
}

type SendMessageOutput struct {
}
