// Code generated by goa v3.11.3, DO NOT EDIT.
//
// PubSubServer service
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server

package pubsubserver

import (
	"context"
)

// The PubSubServer service performs operations on messages.
type Service interface {
	// Publish implements publish.
	Publish(context.Context, *PublishPayload) (res string, err error)
	// Subscribe implements subscribe.
	Subscribe(context.Context, *SubscribePayload) (res string, err error)
	// SendMessage implements sendMessage.
	SendMessage(context.Context, *SendMessagePayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "PubSubServer"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"publish", "subscribe", "sendMessage"}

// PublishPayload is the payload type of the PubSubServer service publish
// method.
type PublishPayload struct {
	// roomName to publish
	RoomName *string
}

// SendMessagePayload is the payload type of the PubSubServer service
// sendMessage method.
type SendMessagePayload struct {
	// roomName to publish
	RoomName *string
	// message to publish
	Message *string
}

// SubscribePayload is the payload type of the PubSubServer service subscribe
// method.
type SubscribePayload struct {
	// roomName to subscribe
	RoomName *string
}
