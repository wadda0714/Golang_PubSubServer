// Code generated by goa v3.2.5, DO NOT EDIT.
//
// PubSubServer endpoints
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server

package pubsubserver

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "PubSubServer" service endpoints.
type Endpoints struct {
	Publish     goa.Endpoint
	Subscribe   goa.Endpoint
	SendMessage goa.Endpoint
}

// NewEndpoints wraps the methods of the "PubSubServer" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Publish:     NewPublishEndpoint(s),
		Subscribe:   NewSubscribeEndpoint(s),
		SendMessage: NewSendMessageEndpoint(s),
	}
}

// Use applies the given middleware to all the "PubSubServer" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Publish = m(e.Publish)
	e.Subscribe = m(e.Subscribe)
	e.SendMessage = m(e.SendMessage)
}

// NewPublishEndpoint returns an endpoint function that calls the method
// "publish" of service "PubSubServer".
func NewPublishEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PublishPayload)
		return s.Publish(ctx, p)
	}
}

// NewSubscribeEndpoint returns an endpoint function that calls the method
// "subscribe" of service "PubSubServer".
func NewSubscribeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubscribePayload)
		return s.Subscribe(ctx, p)
	}
}

// NewSendMessageEndpoint returns an endpoint function that calls the method
// "sendMessage" of service "PubSubServer".
func NewSendMessageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendMessagePayload)
		return s.SendMessage(ctx, p)
	}
}
