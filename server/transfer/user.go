package transfer

import (
	"github.com/wadda0714/Golang_PubSubServer/pkg/userdb"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
)

func ConvertMessages(msg *userdb.Message) *model.Message {
	return &model.Message{
		Content: msg.Content,
		//TimeStamp: msg.Timestamp,
	}

}
