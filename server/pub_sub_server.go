package pubsubserverrestapi

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	pubsubserver "github.com/wadda0714/Golang_PubSubServer/server/gen/pub_sub_server"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/input"
	"github.com/wadda0714/Golang_PubSubServer/server/usecase/output"
	"github.com/wadda0714/Golang_PubSubServer/util"
	"log"
)

// PubSubServer service example implementation.
// The example methods log the requests and return zero values.
type pubSubServersrvc struct {
	logger *log.Logger
}

// NewPubSubServer returns the PubSubServer service implementation.
func NewPubSubServer(logger *log.Logger) pubsubserver.Service {
	return &pubSubServersrvc{logger}
}

// Publish implements publish.
func (s *pubSubServersrvc) Publish(ctx context.Context, p *pubsubserver.PublishPayload) (res string, err error) {
	s.logger.Print("pubSubServer.publish")
	return "aaa", nil

}

// Subscribe implements subscribe.
func (s *pubSubServersrvc) Subscribe(ctx context.Context, p *pubsubserver.SubscribePayload) (res string, err error) {
	s.logger.Print("pubSubServer.subscribe")
	db, err := sql.Open("postgres", "user=root dbname=userdb password=root sslmode=disable port=5430 host=localhost ")
	defer db.Close()
	user := usecase.GeneratedNewUserDI(db)
	var op output.SubscribeOutput
	op, err = user.Subscribe(input.SubscribeInput{RoomID: 1})
	if err != nil {
		panic(err)
	}

	return op.Messages[0].Content, nil
}

// SendMessage implements sendMessage.
func (s *pubSubServersrvc) SendMessage(ctx context.Context, p *pubsubserver.SendMessagePayload) (res string, err error) {
	s.logger.Print("pubSubServer.sendMessage")

	f, err := util.CreateFile("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\n" + *p.Message)

	return "ok", nil

}
