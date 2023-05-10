package pubsubserverrestapi

import (
	"context"
	"log"

	pubsubserver "github.com/wadda0714/Golang_PubSubServer/server/gen/pub_sub_server"
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
	return
}

// Subscribe implements subscribe.
func (s *pubSubServersrvc) Subscribe(ctx context.Context, p *pubsubserver.SubscribePayload) (res string, err error) {
	s.logger.Print("pubSubServer.subscribe")
	return
}

// SendMessage implements sendMessage.
func (s *pubSubServersrvc) SendMessage(ctx context.Context, p *pubsubserver.SendMessagePayload) (res string, err error) {
	s.logger.Print("pubSubServer.sendMessage")
	return
}
